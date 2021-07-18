package handler_test

import (
	"context"
	"io/ioutil"
	v1 "jokes-bapak2-api/app/v1"
	"jokes-bapak2-api/app/v1/platform/database"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	err := database.Setup()
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Query(context.Background(), "INSERT INTO \"administrators\" (id, key, token, last_used) VALUES ($1, $2, $3, $4);", 1, "very secure", "not the real one", time.Now().Format(time.RFC3339))
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Query(context.Background(), "INSERT INTO \"jokesbapak2\" (id, link, creator) VALUES ($1, $2, $3), ($4, $5, $6), ($7, $8, $9);", jokesData...)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(cleanup)

	app := v1.New()

	t.Run("Health - should return 200", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/health", nil)
		res, err := app.Test(req, -1)

		assert.Equalf(t, false, err != nil, "health")
		assert.Equalf(t, 200, res.StatusCode, "health")
		assert.NotEqualf(t, 0, res.ContentLength, "health")
		_, err = ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, "health")
	})
}
