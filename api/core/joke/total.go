package joke

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
)

// GetTotalJoke returns the total jokes that exists on the bucket.
func GetTotalJoke(ctx context.Context, bucket *minio.Client, cache *redis.Client, memory *bigcache.BigCache) (int, error) {
	totalJokesFromMemory, err := memory.Get("total")
	if err != nil && !errors.Is(err, bigcache.ErrEntryNotFound) {
		return 0, fmt.Errorf("acquiring total joke from memory: %w", err)
	}

	if err == nil {
		total, err := strconv.Atoi(string(totalJokesFromMemory))
		if err != nil {
			return 0, fmt.Errorf("parsing string to int: %w", err)
		}

		return total, nil
	}

	totalJokesFromCache, err := cache.Get(ctx, "jokes:total").Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return 0, fmt.Errorf("acquiring total joke from redis: %w", err)
	}

	if err == nil {
		total, err := strconv.Atoi(string(totalJokesFromCache))
		if err != nil {
			return 0, fmt.Errorf("parsing string to int: %w", err)
		}

		return total, nil
	}

	jokes, err := ListJokesFromBucket(ctx, bucket, cache)
	if err != nil {
		return 0, fmt.Errorf("listing jokes: %w", err)
	}

	defer func(total int) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		err := cache.Set(ctx, "jokes:total", strconv.Itoa(total), time.Hour*3).Err()
		if err != nil {
			log.Printf("setting total jokes to redis: %s", err.Error())
		}

		err = memory.Set("total", []byte(strconv.Itoa(total)))
		if err != nil {
			log.Printf("setting total jokes to memory: %s", err.Error())
		}
	}(len(jokes))

	return len(jokes), nil
}
