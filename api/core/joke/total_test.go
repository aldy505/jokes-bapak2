package joke_test

import (
	"context"
	"jokes-bapak2-api/core/joke"
	"testing"
	"time"
)

func TestGetTotalJoke(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	total, err := joke.GetTotalJoke(ctx, bucket, cache, memory)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if total != 5 {
		t.Errorf("expecting total to be 5 instead got %d", total)
	}
}
