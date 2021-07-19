package handler_test

import (
	"context"
	"io/ioutil"
	v1 "jokes-bapak2-api/app/v1"
	"jokes-bapak2-api/app/v1/platform/database"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDeleteJoke(t *testing.T) {
	err := database.Setup()
	if err != nil {
		t.Fatal(err)
	}
	hashedToken := "$argon2id$v=19$m=65536,t=16,p=4$48beb241490caa57fbca8e63df1e1b5fba8934baf78205ee775f96a85f45b889$e6dfca3f69adbe7653dbb353f366d741a3640313c45e33eabaca0c217c16417de80d70ac67f217c9ca46634b0abaad5f4ea2b064caa44ce218fb110b4cba9d36"
	_, err = db.Query(context.Background(), "INSERT INTO \"administrators\" (id, key, token, last_used) VALUES ($1, $2, $3, $4);", 1, "very secure", hashedToken, time.Now().Format(time.RFC3339))
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Query(context.Background(), "INSERT INTO \"jokesbapak2\" (id, link, creator) VALUES ($1, $2, $3), ($4, $5, $6), ($7, $8, $9);", jokesData...)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(cleanup)

	app := v1.New()

	t.Run("Delete - should return 200", func(t *testing.T) {
		t.SkipNow()
		reqBody := strings.NewReader("{\"key\":\"very secure\",\"token\":\"password\"}")
		req, _ := http.NewRequest("DELETE", "/id/1", reqBody)
		res, err := app.Test(req, -1)

		assert.Equalf(t, false, err != nil, "joke delete")
		assert.Equalf(t, 200, res.StatusCode, "joke delete")
		assert.NotEqualf(t, 0, res.ContentLength, "joke delete")
		body, err := ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, "joke delete")
		assert.Equalf(t, "{\"message\":\"specified joke id has been deleted\"}", string(body), "joke delete")
	})

	t.Run("Delete - id doesn't exists", func(t *testing.T) {
		t.SkipNow()
		reqBody := strings.NewReader("{\"key\":\"very secure\",\"token\":\"password\"}")
		req, _ := http.NewRequest("DELETE", "/id/100", reqBody)
		res, err := app.Test(req, -1)

		assert.Equalf(t, false, err != nil, "joke delete")
		assert.Equalf(t, 406, res.StatusCode, "joke delete")
		assert.NotEqualf(t, 0, res.ContentLength, "joke delete")
		body, err := ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, "joke delete")
		assert.Equalf(t, "{\"message\":\"specified joke id does not exists\"}", string(body), "joke delete")
	})
}
