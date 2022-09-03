package joke_test

import (
	"context"
	"fmt"
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
	redisURL, ok := os.LookupEnv("REDIS_URL")
	if !ok {
		redisURL = "redis://@localhost:6379"
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

	parsedRedisURL, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("parsing redis url: %s", err.Error())
		return
	}

	redisClient := redis.NewClient(parsedRedisURL)

	minioClient, err := minio.New(minioHost, &minio.Options{
		Creds: credentials.NewStaticV4(minioID, minioSecret, minioToken),
	})
	if err != nil {
		log.Fatalf("creating minio client: %s", err.Error())
	}

	memoryInstance, err := bigcache.NewBigCache(bigcache.DefaultConfig(time.Second * 30))
	if err != nil {
		log.Fatalf("creating bigcache client: %s", err.Error())
		return
	}

	bucket = minioClient
	cache = redisClient
	memory = memoryInstance

	setupCtx, setupCancel := context.WithTimeout(context.Background(), time.Minute)
	defer setupCancel()

	err = setupBucketStorage(setupCtx, minioClient)
	if err != nil {
		log.Fatalf("set up bucket storage: %v", err)
		return
	}

	exitCode := m.Run()

	cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), time.Minute)
	defer cleanupCancel()

	err = redisClient.FlushAll(cleanupCtx).Err()
	if err != nil {
		log.Printf("flushing redis: %s", err.Error())
	}

	err = minioClient.RemoveBucketWithOptions(cleanupCtx, "jokesbapak2", minio.RemoveBucketOptions{ForceDelete: true})
	if err != nil {
		log.Printf("removing bucket: %s", err.Error())
	}

	err = memoryInstance.Close()
	if err != nil {
		log.Printf("closing cache client: %s", err.Error())
	}

	err = redisClient.Close()
	if err != nil {
		log.Printf("closing redis client: %s", err.Error())
	}

	os.Exit(exitCode)
}

func setupBucketStorage(ctx context.Context, minioClient *minio.Client) error {
	bucketFound, err := minioClient.BucketExists(ctx, "jokesbapak2")
	if err != nil {
		return fmt.Errorf("checking MinIO bucket: %w", err)
	}

	if !bucketFound {
		err = minioClient.MakeBucket(ctx, "jokesbapak2", minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("creating MinIO bucket: %w", err)
		}

		policy := `{
			"Version":"2012-10-17",
			"Statement":[
			  {
				"Sid": "AddPerm",
				"Effect": "Allow",
				"Principal": "*",
				"Action":["s3:GetObject"],
				"Resource":["arn:aws:s3:::jokesbapak2/*"]
			  }
			]
		  }`

		err = minioClient.SetBucketPolicy(ctx, "jokesbapak2", policy)
		if err != nil {
			return fmt.Errorf("setting bucket policy: %w", err)
		}
	}

	sampleFiles := []string{
		"../../samples/sample1.jpg",
		"../../samples/sample2.jpg",
		"../../samples/sample3.jpg",
		"../../samples/sample4.jpg",
		"../../samples/sample5.jpg",
	}

	for i, file := range sampleFiles {
		_, err := minioClient.FPutObject(ctx, "jokesbapak2", fmt.Sprintf("sample%d.jpg", i), file, minio.PutObjectOptions{ContentType: "image/jpeg"})
		if err != nil {
			return fmt.Errorf("putting object: %w", err)
		}
	}

	return nil
}
