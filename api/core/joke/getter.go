package joke

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
)

// GetRandomJoke will acquire a random joke from the bucket.
func GetRandomJoke(ctx context.Context, bucket *minio.Client, cache *redis.Client, memory *bigcache.BigCache) (image []byte, contentType string, err error) {
	totalJokes, err := GetTotalJoke(ctx, bucket, cache, memory)
	if err != nil {
		return []byte{}, "", fmt.Errorf("getting total joke: %w", err)
	}

	randomIndex := rand.Intn(totalJokes - 1)

	joke, contentType, err := GetJokeByID(ctx, bucket, cache, memory, randomIndex)
	if err != nil {
		return []byte{}, "", fmt.Errorf("getting joke by id: %w", err)
	}

	return joke, contentType, nil
}

// GetJokeByID wil acquire a joke by its' ID.
//
// An ID is defined as the index on the joke list that is sorted
// by it's creation (or modification) time.
func GetJokeByID(ctx context.Context, bucket *minio.Client, cache *redis.Client, memory *bigcache.BigCache, id int) (image []byte, contentType string, err error) {
	jokeFromMemory, err := memory.Get("id:" + strconv.Itoa(id))
	if err != nil && !errors.Is(err, bigcache.ErrEntryNotFound) {
		return []byte{}, "", fmt.Errorf("acquiring joke from memory: %w", err)
	}

	if err == nil {
		contentTypeFromMemory, err := memory.Get("id:" + strconv.Itoa(id) + ":content-type")
		if err != nil && !errors.Is(err, bigcache.ErrEntryNotFound) {
			return []byte{}, "", fmt.Errorf("acquiring joke content type from memory: %w", err)
		}

		return jokeFromMemory, string(contentTypeFromMemory), nil
	}

	jokeFromCache, err := cache.Get(ctx, "jokes:id:"+strconv.Itoa(id)).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return []byte{}, "", fmt.Errorf("acquiring joke from cache: %w", err)
	}

	if err == nil {
		// Get content type
		contentTypeFromCache, err := cache.Get(ctx, "jokes:id:"+strconv.Itoa(id)+":content-type").Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return []byte{}, "", fmt.Errorf("acquiring content type from cache: %w", err)
		}

		// Decode hex string to bytes
		imageBytes, err := hex.DecodeString(jokeFromCache)
		if err != nil {
			return []byte{}, "", fmt.Errorf("decoding hex string: %w", err)
		}

		defer func(id int, imageBytes []byte) {
			err := memory.Set("id:"+strconv.Itoa(id), imageBytes)
			if err != nil {
				log.Printf("setting memory cache: %s", err.Error())
			}

			err = memory.Set("id:"+strconv.Itoa(id)+":content-type", []byte(contentTypeFromCache))
			if err != nil {
				log.Printf("setting memory cache: %s", err.Error())
			}
		}(id, imageBytes)

		return imageBytes, contentTypeFromCache, nil
	}

	jokes, err := ListJokesFromBucket(ctx, bucket, cache)
	if err != nil {
		return []byte{}, "", fmt.Errorf("listing jokes: %w", err)
	}

	object, err := bucket.GetObject(ctx, JokesBapak2Bucket, jokes[id].FileName, minio.GetObjectOptions{})
	if err != nil {
		return []byte{}, "", fmt.Errorf("getting object: %w", err)
	}
	defer func() {
		err := object.Close()
		if err != nil {
			log.Printf("closing image reader: %s", err.Error())
		}
	}()

	image, err = io.ReadAll(object)
	if err != nil {
		return []byte{}, "", fmt.Errorf("reading object: %w", err)
	}

	defer func(id int, image []byte) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		imageString := hex.EncodeToString(image)

		err := cache.Set(ctx, "jokes:id:"+strconv.Itoa(id), imageString, time.Hour*1).Err()
		if err != nil {
			log.Printf("setting cache: %s", err.Error())
		}

		err = cache.Set(ctx, "jokes:id:"+strconv.Itoa(id)+":content-type", jokes[id].ContentType, time.Hour*1).Err()
		if err != nil {
			log.Printf("setting cache: %s", err.Error())
		}
	}(id, image)

	return image, jokes[id].ContentType, nil
}
