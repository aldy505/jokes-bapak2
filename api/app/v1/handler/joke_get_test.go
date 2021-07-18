package handler_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	v1 "jokes-bapak2-api/app/v1"
	"jokes-bapak2-api/app/v1/platform/database"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

var db = database.New()
var jokesData = []interface{}{1, "https://picsum.photos/id/1/200/300", 1, 2, "https://picsum.photos/id/2/200/300", 1, 3, "https://picsum.photos/id/3/200/300", 1}

func cleanup() {
	_, err := db.Query(context.Background(), "DROP TABLE \"jokesbapak2\"")
	if err != nil {
		panic(err)
	}
	_, err = db.Query(context.Background(), "DROP TABLE \"administrators\"")
	if err != nil {
		panic(err)
	}
}

/// Need to find some workaround for this test
func TestJokeGet(t *testing.T) {
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

	t.Run("TodayJoke - should return 200", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/today", nil)
		res, err := app.Test(req, -1)

		assert.Equalf(t, false, err != nil, "today joke")
		assert.Equalf(t, 200, res.StatusCode, "today joke")
		assert.NotEqualf(t, 0, res.ContentLength, "today joke")
		_, err = ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, "today joke")
	})

	t.Run("SingleJoke - should return 200", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		res, err := app.Test(req, -1)

		assert.Equalf(t, false, err != nil, "single joke")
		assert.Equalf(t, 200, res.StatusCode, "single joke")
		assert.NotEqualf(t, 0, res.ContentLength, "single joke")
		_, err = ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, "single joke")
	})

	t.Run("JokeByID - should return 200", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/id/1", nil)
		res, err := app.Test(req, -1)

		assert.Equalf(t, false, err != nil, "joke by id")
		assert.Equalf(t, 200, res.StatusCode, "joke by id")
		assert.NotEqualf(t, 0, res.ContentLength, "joke by id")
		_, err = ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, "joke by id")
	})

	t.Run("JokeByID - should return 404", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/id/300", nil)
		res, err := app.Test(req, -1)

		assert.Equalf(t, false, err != nil, "joke by id")
		assert.Equalf(t, 404, res.StatusCode, "joke by id")
		assert.NotEqualf(t, 0, res.ContentLength, "joke by id")
		body, err := ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, "joke by id")
		assert.Equalf(t, "Requested ID was not found.", string(body), "joke by id")
	})
}
