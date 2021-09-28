package main_test

import (
	"context"
	"errors"
	"flag"
	"io/ioutil"
	v1 "jokes-bapak2-api/app"
	"jokes-bapak2-api/app/platform/database"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

var jokesData = []interface{}{1, "https://via.placeholder.com/300/06f/fff.png", 1, 2, "https://via.placeholder.com/300/07f/fff.png", 1, 3, "https://via.placeholder.com/300/08f/fff.png", 1}
var submissionData = []interface{}{1, "https://via.placeholder.com/300/01f/fff.png", "2021-08-03T18:20:38Z", "Test <test@example.com>", 0, 2, "https://via.placeholder.com/300/02f/fff.png", "2021-08-04T18:20:38Z", "Test <test@example.com>", 1}
var administratorsData = []interface{}{1, "very secure", "not the real one", time.Now().Format(time.RFC3339), 2, "test", "$argon2id$v=19$m=65536,t=16,p=4$3a08c79fbf2222467a623df9a9ebf75802c65a4f9be36eb1df2f5d2052d53cb7$ce434bd38f7ba1fc1f2eb773afb8a1f7f2dad49140803ac6cb9d7256ce9826fb3b4afa1e2488da511c852fc6c33a76d5657eba6298a8e49d617b9972645b7106", ""}
var ctx context.Context = context.Background()

func TestMain(m *testing.M) {
	flag.Parse()

	log.Println("---- Preparing for integration test")
	time.Sleep(time.Second * 5)
	err := setup()
	if err != nil {
		log.Panicln(err)
	}
	time.Sleep(time.Second * 5)
	log.Println("---- Preparation complete")
	log.Print("\n")

	os.Exit(m.Run())
}

func setup() error {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return errors.New("Unable to create pool config: " + err.Error())
	}

	db, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return errors.New("Unable to create connection: " + err.Error())
	}

	dj, err := db.Query(ctx, "DROP TABLE IF EXISTS \"jokesbapak2\"")
	if err != nil {
		return err
	}
	dj.Close()

	ds, err := db.Query(ctx, "DROP TABLE IF EXISTS \"submission\"")
	if err != nil {
		return err
	}
	ds.Close()

	da, err := db.Query(ctx, "DROP TABLE IF EXISTS \"administrators\"")
	if err != nil {
		return err
	}
	da.Close()

	err = database.Setup(db, &ctx)
	if err != nil {
		return err
	}

	ia, err := db.Query(ctx, "INSERT INTO \"administrators\" (id, key, token, last_used) VALUES ($1, $2, $3, $4), ($5, $6, $7, $8);", administratorsData...)
	if err != nil {
		return err
	}
	ia.Close()

	ij, err := db.Query(ctx, "INSERT INTO \"jokesbapak2\" (id, link, creator) VALUES ($1, $2, $3), ($4, $5, $6), ($7, $8, $9);", jokesData...)
	if err != nil {
		return err
	}
	ij.Close()

	is, err := db.Query(ctx, "INSERT INTO \"submission\" (id, link, created_at, author, status) VALUES ($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10);", submissionData...)
	if err != nil {
		return err
	}
	is.Close()

	db.Close()

	return nil
}

var app *fiber.App = v1.New()

func TestHealth(t *testing.T) {
	req, _ := http.NewRequest("GET", "/health", nil)
	res, err := app.Test(req, int(time.Minute*2))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equalf(t, false, err != nil, "health")
	assert.Equalf(t, 200, res.StatusCode, "health")
	assert.NotEqualf(t, 0, res.ContentLength, "health")
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Nilf(t, err, "health")
}

func TestTodayJoke(t *testing.T) {
	req, _ := http.NewRequest("GET", "/today", nil)
	res, err := app.Test(req, int(time.Minute*2))

	assert.Equalf(t, false, err != nil, "today joke")
	assert.Equalf(t, 200, res.StatusCode, "today joke")
	assert.NotEqualf(t, 0, res.ContentLength, "today joke")
	_, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	assert.Nilf(t, err, "today joke")
}

func TestSingleJoke(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res, err := app.Test(req, int(time.Minute*2))

	assert.Equalf(t, false, err != nil, "single joke")
	assert.Equalf(t, 200, res.StatusCode, "single joke")
	assert.NotEqualf(t, 0, res.ContentLength, "single joke")
	_, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	assert.Nilf(t, err, "single joke")
}

func TestJokeByID_200(t *testing.T) {
	req, _ := http.NewRequest("GET", "/id/1", nil)
	res, err := app.Test(req, int(time.Minute*2))

	assert.Equalf(t, false, err != nil, "joke by id")
	assert.Equalf(t, 200, res.StatusCode, "joke by id")
	assert.NotEqualf(t, 0, res.ContentLength, "joke by id")
	_, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	assert.Nilf(t, err, "joke by id")
}

func TestJokeByID_404(t *testing.T) {
	req, _ := http.NewRequest("GET", "/id/300", nil)
	res, err := app.Test(req, int(time.Minute*2))

	assert.Equalf(t, false, err != nil, "joke by id")
	assert.Equalf(t, 404, res.StatusCode, "joke by id")
	assert.NotEqualf(t, 0, res.ContentLength, "joke by id")
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	assert.Nilf(t, err, "joke by id")
	assert.Equalf(t, "Requested ID was not found.", string(body), "joke by id")
}

func TestTotalJokes(t *testing.T) {
	req, _ := http.NewRequest("GET", "/total", nil)
	res, err := app.Test(req, int(time.Minute*2))

	assert.Equalf(t, false, err != nil, "joke total")
	assert.Equalf(t, 200, res.StatusCode, "joke total")
	assert.NotEqualf(t, 0, res.ContentLength, "joke total")
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	assert.Nilf(t, err, "joke total")
	assert.Equalf(t, "{\"message\":\"3\"}", string(body), "joke total")
}

func TestAddNewJoke_201(t *testing.T) {
	// TODO: Remove this line below, make this test works
	t.SkipNow()

	reqBody := strings.NewReader("{\"link\":\"https://via.placeholder.com/300/04f/ff0000.png\",\"key\":\"test\",\"token\":\"password\"}")
	req, _ := http.NewRequest("PUT", "/", reqBody)
	req.Header.Set("content-type", "application/json")
	req.Header.Add("accept", "application/json")
	res, err := app.Test(req, int(time.Minute*2))

	assert.Equalf(t, false, err != nil, "joke add")
	assert.Equalf(t, 201, res.StatusCode, "joke add")
	assert.NotEqualf(t, 0, res.ContentLength, "joke add")
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	assert.Nilf(t, err, "joke add")
	assert.Equalf(t, "{\"link\":\"https://via.placeholder.com/300/04f/ff0000.png\"}", string(body), "joke add")
}

func TestAddNewJoke_NotValidImage(t *testing.T) {
	// TODO: Remove this line below, make this test works
	t.SkipNow()

	reqBody := strings.NewReader("{\"link\":\"https://google.com/\",\"key\":\"test\",\"token\":\"password\"}")
	req, _ := http.NewRequest("PUT", "/", reqBody)
	req.Header.Set("content-type", "application/json")
	req.Header.Add("accept", "application/json")
	res, err := app.Test(req, int(time.Minute*2))

	assert.Equalf(t, false, err != nil, "joke add")
	assert.Equalf(t, 400, res.StatusCode, "joke add")
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	assert.Nilf(t, err, "joke add")
	assert.Equalf(t, "{\"error\":\"URL provided is not a valid image\"}", string(body), "joke add")
}

func TestGetSubmission_200(t *testing.T) {
	req, _ := http.NewRequest("GET", "/submit", nil)
	res, err := app.Test(req, int(time.Minute*2))

	assert.Equalf(t, false, err != nil, "get submission")
	assert.Equalf(t, 200, res.StatusCode, "get submission")
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	assert.Nilf(t, err, "get submission")
	assert.Equalf(t, "{\"count\":2,\"jokes\":[{\"id\":1,\"link\":\"https://via.placeholder.com/300/01f/fff.png\",\"created_at\":\"2021-08-03T18:20:38Z\",\"author\":\"Test \\u003ctest@example.com\\u003e\",\"status\":0},{\"id\":2,\"link\":\"https://via.placeholder.com/300/02f/fff.png\",\"created_at\":\"2021-08-04T18:20:38Z\",\"author\":\"Test \\u003ctest@example.com\\u003e\",\"status\":1}]}", string(body), "get submission")
}

func TestGetSubmission_Params(t *testing.T) {
	req, _ := http.NewRequest("GET", "/submit?page=1&limit=5&approved=true", nil)
	res, err := app.Test(req, int(time.Minute*2))

	assert.Equalf(t, false, err != nil, "get submission")
	assert.Equalf(t, 200, res.StatusCode, "get submission")
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	assert.Nilf(t, err, "get submission")
	assert.Equalf(t, "{\"count\":1,\"jokes\":[{\"id\":2,\"link\":\"https://via.placeholder.com/300/02f/fff.png\",\"created_at\":\"2021-08-04T18:20:38Z\",\"author\":\"Test \\u003ctest@example.com\\u003e\",\"status\":1}]}", string(body), "get submission")
}

func TestAddSubmission_200(t *testing.T) {
	// TODO: Remove this line below, make this test works
	t.Skip()

	reqBody := strings.NewReader(`{"link":"https://via.placeholder.com/400/02f/fff.png","author":"Test <test@mail.com>"}`)
	req, _ := http.NewRequest("POST", "/submit", reqBody)
	req.Header.Set("content-type", "application/json")
	req.Header.Add("accept", "application/json")
	res, err := app.Test(req, int(time.Minute*2))

	assert.Equalf(t, false, err != nil, "post submission")
	assert.Equalf(t, 201, res.StatusCode, "post submission")
	assert.NotEqualf(t, 0, res.ContentLength, "post submission")
	_, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	assert.Nilf(t, err, "post submission")
}
