package joke

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

// Uploader uploads a reader stream (io.Reader) to bucket.
func Uploader(ctx context.Context, bucket *minio.Client, key string, payload io.Reader, fileSize int64, contentType string) (string, error) {
	info, err := bucket.PutObject(
		ctx,
		JokesBapak2Bucket, // bucketName
		key,               // object name,
		payload,           // reader
		fileSize,          // obuject size,
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
	if err != nil {
		return "", fmt.Errorf("uploading object: %w", err)
	}

	return info.Key, nil
}
