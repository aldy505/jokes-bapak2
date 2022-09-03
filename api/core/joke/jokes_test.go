package joke_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var bucket *minio.Client
var cache *redis.Client
var memory *bigcache.BigCache

func TestMain(m *testing.M) {
	redisUrl, ok := os.LookupEnv("REDIS_URL")
	if !ok {
		redisUrl = "redis://@localhost:6379"
	}

	minioHost, ok := os.LookupEnv("MINIO_HOST")
	if !ok {
		minioHost = "localhost:9000"
	}

	minioID, ok := os.LookupEnv("MINIO_ACCESS_ID")
	if !ok {
		minioID = "minio"
	}

	minioSecret, ok := os.LookupEnv("MINIO_SECRET_KEY")
	if !ok {
		minioSecret = "password"
	}

	minioToken, ok := os.LookupEnv("MINIO_TOKEN")
	if !ok {
		minioToken = ""
	}

	parsedRedisUrl, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Fatalf("parsing redis url: %s", err.Error())
		return
	}

	redisClient := redis.NewClient(parsedRedisUrl)

	minioClient, err := minio.New(minioHost, &minio.Options{
		Creds: credentials.NewStaticV4(minioID, minioSecret, minioToken),
	})
	if err != nil {
		log.Fatalf("creating minio client: %s", err.Error())
	}

	memoryInstance, err := bigcache.NewBigCache(bigcache.DefaultConfig(time.Second * 30))
	if err != nil {
		log.Fatalf("creating bigcache client: %s", err.Error())
	}

	bucket = minioClient
	cache = redisClient
	memory = memoryInstance

	exitCode := m.Run()

	cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), time.Minute)
	defer cleanupCancel()

	err = redisClient.FlushAll(cleanupCtx).Err()
	if err != nil {
		log.Printf("flushing redis: %s", err.Error())
	}

	err = cache.Close()
	if err != nil {
		log.Printf("closing cache client: %s", err.Error())
	}

	err = redisClient.Close()
	if err != nil {
		log.Printf("closing redis client: %s", err.Error())
	}

	os.Exit(exitCode)
}
