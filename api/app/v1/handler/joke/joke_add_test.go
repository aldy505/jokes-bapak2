package joke_test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNewJoke_201(t *testing.T) {
	t.SkipNow()
	err := setup()
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup()

	reqBody := strings.NewReader("{\"link\":\"https://via.placeholder.com/300/07f/ff0000.png\",\"key\":\"very secure\",\"token\":\"password\"}")
	req, _ := http.NewRequest("PUT", "/", reqBody)
	req.Header.Set("content-type", "application/json")
	req.Header.Add("accept", "application/json")
	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "joke add")
	assert.Equalf(t, 201, res.StatusCode, "joke add")
	assert.NotEqualf(t, 0, res.ContentLength, "joke add")
	body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "joke add")
	assert.Equalf(t, "{\"link\":\"https://via.placeholder.com/300/07f/ff0000.png\"}", string(body), "joke add")
}

func TestAddNewJoke_NotValidImage(t *testing.T) {
	t.SkipNow()
	err := setup()
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup()

	reqBody := strings.NewReader("{\"link\":\"https://google.com/\",\"key\":\"very secure\",\"token\":\"password\"}")
	req, _ := http.NewRequest("PUT", "/", reqBody)
	req.Header.Set("content-type", "application/json")
	req.Header.Add("accept", "application/json")
	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "joke add")
	assert.Equalf(t, 400, res.StatusCode, "joke add")
	body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "joke add")
	assert.Equalf(t, "{\"error\":\"URL provided is not a valid image\"}", string(body), "joke add")
}
