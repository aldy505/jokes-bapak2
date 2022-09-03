package joke

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
)

func ListJokesFromBucket(ctx context.Context, bucket *minio.Client, cache *redis.Client) ([]Joke, error) {
	cached, err := cache.Get(ctx, "jokes:list").Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return []Joke{}, fmt.Errorf("acquiring joke list from cache: %w", err)
	}

	if err == nil {
		var jokes []Joke
		err := json.Unmarshal([]byte(cached), &jokes)
		if err != nil {
			return []Joke{}, fmt.Errorf("unmarshalling json: %w", err)
		}

		return jokes, nil
	}

	objects := bucket.ListObjects(ctx, JokesBapak2Bucket, minio.ListObjectsOptions{Recursive: true})

	var jokes []Joke
	for object := range objects {
		if object.Err != nil {
			return []Joke{}, fmt.Errorf("enumerating objects: %w", object.Err)
		}

		if !object.IsDeleteMarker {
			jokes = append(jokes, Joke{ModifiedAt: object.Restore.ExpiryTime, FileName: object.Key, ContentType: object.ContentType})
		}
	}

	sort.SliceStable(jokes, func(i, j int) bool {
		return jokes[i].ModifiedAt.Before(jokes[i].ModifiedAt)
	})

	defer func(jokes []Joke) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		marshalled, err := json.Marshal(jokes)
		if err != nil {
			log.Printf("marshalling json: %s", err.Error())
			return
		}

		err = cache.Set(ctx, "jokes:list", string(marshalled), time.Hour*6).Err()
		if err != nil {
			log.Printf("setting jokes:list cache: %s", err.Error())
		}
	}(jokes)

	return jokes, nil
}
