package submit_test

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/stretchr/testify/assert"
)

func TestGetSubmission_200(t *testing.T) {
	err := setup()
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup()

	req, _ := http.NewRequest("GET", "/submit", nil)
	res, err := app.Test(req, int(time.Minute * 2))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equalf(t, false, err != nil, "get submission")
	assert.Equalf(t, 200, res.StatusCode, "get submission")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

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
	res, err := app.Test(req, int(time.Minute * 2))
	if err != nil {
		t.Fatal(err)
	}
	
	assert.Equalf(t, false, err != nil, "get submission")
	assert.Equalf(t, 200, res.StatusCode, "get submission")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Nilf(t, err, "get submission")
	assert.Equalf(t, "{\"count\":1,\"jokes\":[{\"id\":2,\"link\":\"https://via.placeholder.com/300/02f/fff.png\",\"created_at\":\"2021-08-04T18:20:38Z\",\"author\":\"Test \\u003ctest@example.com\\u003e\",\"status\":1}]}", string(body), "get submission")
}
