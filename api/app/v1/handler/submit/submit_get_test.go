package submit_test

import (
	"context"
	"io/ioutil"
	v1 "jokes-bapak2-api/app/v1"
	"jokes-bapak2-api/app/v1/platform/database"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/joho/godotenv/autoload"

	"github.com/stretchr/testify/assert"
)

var db *pgxpool.Pool = database.New()
var submissionData = []interface{}{1, "https://via.placeholder.com/300/01f/fff.png", "2021-08-03T18:20:38Z", "Test <test@example.com>", 0, 2, "https://via.placeholder.com/300/02f/fff.png", "2021-08-04T18:20:38Z", "Test <test@example.com>", 1}
var app *fiber.App = v1.New()

func cleanup() {
	s, err := db.Query(context.Background(), "DROP TABLE \"submission\"")
	if err != nil {
		panic(err)
	}
	defer s.Close()
}

func setup() error {
	err := database.Setup()
	if err != nil {
		return err
	}

	s, err := db.Query(context.Background(), "INSERT INTO \"submission\" (id, link, created_at, author, status) VALUES ($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10);", submissionData...)
	if err != nil {
		return err
	}

	defer s.Close()

	return nil
}
func TestGetSubmission_200(t *testing.T) {
	err := setup()
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup()

	req, _ := http.NewRequest("GET", "/submit", nil)
	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "get submission")
	assert.Equalf(t, 200, res.StatusCode, "get submission")
	body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "get submission")
	assert.Equalf(t, "{\"count\":2,\"jokes\":[{\"id\":1,\"link\":\"https://via.placeholder.com/300/01f/fff.png\",\"created_at\":\"2021-08-03T18:20:38Z\",\"author\":\"Test \\u003ctest@example.com\\u003e\",\"status\":0},{\"id\":2,\"link\":\"https://via.placeholder.com/300/02f/fff.png\",\"created_at\":\"2021-08-04T18:20:38Z\",\"author\":\"Test \\u003ctest@example.com\\u003e\",\"status\":1}]}", string(body), "get submission")
}

func TestGetSubmission_Params(t *testing.T) {
	err := setup()
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup()

	req, _ := http.NewRequest("GET", "/submit?page=1&limit=5&approved=true", nil)
	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "get submission")
	assert.Equalf(t, 200, res.StatusCode, "get submission")
	body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "get submission")
	assert.Equalf(t, "{\"count\":1,\"jokes\":[{\"id\":2,\"link\":\"https://via.placeholder.com/300/02f/fff.png\",\"created_at\":\"2021-08-04T18:20:38Z\",\"author\":\"Test \\u003ctest@example.com\\u003e\",\"status\":1}]}", string(body), "get submission")
}
