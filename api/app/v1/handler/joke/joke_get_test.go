package joke_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	v1 "jokes-bapak2-api/app/v1"
	"jokes-bapak2-api/app/v1/platform/database"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

var db *pgxpool.Pool = database.New()
var jokesData = []interface{}{1, "https://via.placeholder.com/300/06f/fff.png", 1, 2, "https://via.placeholder.com/300/07f/fff.png", 1, 3, "https://via.placeholder.com/300/08f/fff.png", 1}
var app *fiber.App = v1.New()

func cleanup() {
	j, err := db.Query(context.Background(), "DROP TABLE \"jokesbapak2\"")
	if err != nil {
		panic(err)
	}
	a, err := db.Query(context.Background(), "DROP TABLE \"administrators\"")
	if err != nil {
		panic(err)
	}

	defer j.Close()
	defer a.Close()
}

func setup() error {
	err := database.Setup()
	if err != nil {
		return err
	}

	hashedToken := "$argon2id$v=19$m=65536,t=16,p=4$48beb241490caa57fbca8e63df1e1b5fba8934baf78205ee775f96a85f45b889$e6dfca3f69adbe7653dbb353f366d741a3640313c45e33eabaca0c217c16417de80d70ac67f217c9ca46634b0abaad5f4ea2b064caa44ce218fb110b4cba9d36"
	var args []interface{} = []interface{}{1, "very secure", hashedToken, time.Now().Format(time.RFC3339)}
	a, err := db.Query(context.Background(), "INSERT INTO \"administrators\" (id, key, token, last_used) VALUES ($1, $2, $3, $4);", args...)
	if err != nil {
		return err
	}

	defer a.Close()

	j, err := db.Query(context.Background(), "INSERT INTO \"jokesbapak2\" (id, link, creator) VALUES ($1, $2, $3), ($4, $5, $6), ($7, $8, $9);", jokesData...)
	if err != nil {
		return err
	}

	defer j.Close()

	return nil
}

/// Need to find some workaround for this test
func TestTodayJoke(t *testing.T) {
	err := setup()
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup()

	req, _ := http.NewRequest("GET", "/today", nil)
	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "today joke")
	assert.Equalf(t, 200, res.StatusCode, "today joke")
	assert.NotEqualf(t, 0, res.ContentLength, "today joke")
	_, err = ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "today joke")
}

func TestSingleJoke(t *testing.T) {
	err := setup()
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup()

	req, _ := http.NewRequest("GET", "/", nil)
	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "single joke")
	assert.Equalf(t, 200, res.StatusCode, "single joke")
	assert.NotEqualf(t, 0, res.ContentLength, "single joke")
	_, err = ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "single joke")
}

func TestJokeByID_200(t *testing.T) {
	err := setup()
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup()

	req, _ := http.NewRequest("GET", "/id/1", nil)
	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "joke by id")
	assert.Equalf(t, 200, res.StatusCode, "joke by id")
	assert.NotEqualf(t, 0, res.ContentLength, "joke by id")
	_, err = ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "joke by id")
}

func TestJokeByID_404(t *testing.T) {
	err := setup()
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup()

	req, _ := http.NewRequest("GET", "/id/300", nil)
	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "joke by id")
	assert.Equalf(t, 404, res.StatusCode, "joke by id")
	assert.NotEqualf(t, 0, res.ContentLength, "joke by id")
	body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "joke by id")
	assert.Equalf(t, "Requested ID was not found.", string(body), "joke by id")
}
