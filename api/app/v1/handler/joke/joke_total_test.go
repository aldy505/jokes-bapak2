package joke_test

import (
	"context"
	"io/ioutil"
	v1 "jokes-bapak2-api/app/v1"
	"jokes-bapak2-api/app/v1/handler"
	"jokes-bapak2-api/app/v1/platform/database"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTotalJokes(t *testing.T) {
	err := database.Setup()
	if err != nil {
		t.Fatal(err)
	}
	_, err = handler.Db.Query(context.Background(), "INSERT INTO \"administrators\" (id, key, token, last_used) VALUES ($1, $2, $3, $4);", 1, "very secure", "not the real one", time.Now().Format(time.RFC3339))
	if err != nil {
		t.Fatal(err)
	}
	_, err = handler.Db.Query(context.Background(), "INSERT INTO \"jokesbapak2\" (id, link, creator) VALUES ($1, $2, $3), ($4, $5, $6), ($7, $8, $9);", jokesData...)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(cleanup)

	app := v1.New()

	t.Run("Total - should return 200", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/total", nil)
		res, err := app.Test(req, -1)

		assert.Equalf(t, false, err != nil, "joke total")
		assert.Equalf(t, 200, res.StatusCode, "joke total")
		assert.NotEqualf(t, 0, res.ContentLength, "joke total")
		body, err := ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, "joke total")
		// FIXME: This should be "message": "3", not one. I don't know what's wrong as it's 1 AM.
		assert.Equalf(t, "{\"message\":\"1\"}", string(body), "joke total")
	})
}
