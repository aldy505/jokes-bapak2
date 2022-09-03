package joke

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
)

func GetTodaysJoke(ctx context.Context, bucket *minio.Client, cache *redis.Client, memory *bigcache.BigCache) (image []byte, contentType string, err error) {
	// Today's date:
	today := time.Now().Format("2006-01-02")

	jokeFromMemory, err := memory.Get("today:file:" + today)
	if err != nil && !errors.Is(err, bigcache.ErrEntryNotFound) {
		return []byte{}, "", fmt.Errorf("acquiring joke from memory: %w", err)
	}

	if err == nil {
		contentTypeFromMemory, err := memory.Get("today:content-type:" + today)
		if err != nil && !errors.Is(err, bigcache.ErrEntryNotFound) {
			return []byte{}, "", fmt.Errorf("acquiring joke content type from memory: %w", err)
		}

		return jokeFromMemory, string(contentTypeFromMemory), nil
	}

	jokeFromCache, err := cache.Get(ctx, "jokes:today:"+today).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return []byte{}, "", fmt.Errorf("acquiring joke from cache: %w", err)
	}

	if err == nil {
		// Get content type
		contentTypeFromCache, err := cache.Get(ctx, "jokes:today:"+today+":content-type").Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return []byte{}, "", fmt.Errorf("acquiring content type from cache: %w", err)
		}

		// Decode hex string to bytes
		imageBytes, err := hex.DecodeString(jokeFromCache)
		if err != nil {
			return []byte{}, "", fmt.Errorf("decoding hex string: %w", err)
		}

		defer func(today string, imageBytes []byte) {
			err := memory.Set("today:"+today, imageBytes)
			if err != nil {
				log.Printf("setting memory cache: %s", err.Error())
			}

			err = memory.Set("today:"+today+":content-type", []byte(contentTypeFromCache))
			if err != nil {
				log.Printf("setting memory cache: %s", err.Error())
			}
		}(today, imageBytes)

		return imageBytes, contentTypeFromCache, nil
	}

	// If everything not exists, we get a new random joke
	randomJoke, contentType, err := GetRandomJoke(ctx, bucket, cache, memory)
	if err != nil {
		return []byte{}, "", fmt.Errorf("acquiring new random joke: %w", err)
	}

	defer func(today string, joke []byte) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		// Encode to hex string
		encodedImage := hex.EncodeToString(joke)

		err := cache.Set(ctx, "jokes:today:"+today, encodedImage, time.Hour*24).Err()
		if err != nil {
			log.Printf("setting today cache to redis: %s", err.Error())
		}

		err = cache.Set(ctx, "jokes:today:"+today+":content-type", contentType, time.Hour*24).Err()
		if err != nil {
			log.Printf("setting today cache to redis: %s", err.Error())
		}
	}(today, randomJoke)

	return randomJoke, contentType, nil
}
