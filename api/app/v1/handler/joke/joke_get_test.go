package joke_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

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

	a, err := db.Query(context.Background(), "INSERT INTO \"administrators\" (\"id\", \"key\", \"token\", \"last_used\") VALUES	(1, 'test', '$argon2id$v=19$m=65536,t=16,p=4$3a08c79fbf2222467a623df9a9ebf75802c65a4f9be36eb1df2f5d2052d53cb7$ce434bd38f7ba1fc1f2eb773afb8a1f7f2dad49140803ac6cb9d7256ce9826fb3b4afa1e2488da511c852fc6c33a76d5657eba6298a8e49d617b9972645b7106', '');")
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
