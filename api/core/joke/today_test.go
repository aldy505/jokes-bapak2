package joke_test

import (
	"context"
	"jokes-bapak2-api/core/joke"
	"testing"
	"time"
)

func TestGetTodaysJoke(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	image, contentType, err := joke.GetTodaysJoke(ctx, bucket, cache, memory)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if contentType != "image/jpeg" {
		t.Errorf("expecting contentType of 'image/jpeg', instead got %s", contentType)
	}

	if len(image) == 0 {
		t.Errorf("empty image")
	}

	cachedImage, cachedContentType, err := joke.GetTodaysJoke(ctx, bucket, cache, memory)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if contentType != cachedContentType {
		t.Errorf("difference on contentType: original %s vs cached %s", contentType, cachedContentType)
	}

	if string(image) != string(cachedImage) {
		t.Errorf("difference in image")
	}
}
