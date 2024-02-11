package joke_test

import (
	"context"
	"testing"
	"time"

	"jokes-bapak2-api/core/joke"
)

func TestGetRandomJoke(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	image, contentType, err := joke.GetRandomJoke(ctx, bucket, cache, memory)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if contentType != "image/jpeg" {
		t.Errorf("expecting contentType of 'image/jpeg', instead got %s", contentType)
	}

	if len(image) == 0 {
		t.Error("empty image")
	}
}

func TestGetJokeById(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	image, contentType, err := joke.GetJokeByID(ctx, bucket, cache, memory, 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if contentType != "image/jpeg" {
		t.Errorf("expecting contentType of 'image/jpeg', instead got %s", contentType)
	}

	if len(image) == 0 {
		t.Error("empty image")
	}

	cachedImage, cachedContentType, err := joke.GetJokeByID(ctx, bucket, cache, memory, 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if cachedContentType != contentType {
		t.Errorf("difference in contentType: original %s vs cached %s", contentType, cachedContentType)
	}

	if string(cachedImage) != string(image) {
		t.Errorf("difference in image bytes")
	}

	cachedImage2, cachedContentType2, err := joke.GetJokeByID(ctx, bucket, cache, memory, 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if cachedContentType2 != contentType {
		t.Errorf("difference in contentType: original %s vs cached %s", contentType, cachedContentType2)
	}

	if string(cachedImage2) != string(image) {
		t.Errorf("difference in image bytes")
	}

}
