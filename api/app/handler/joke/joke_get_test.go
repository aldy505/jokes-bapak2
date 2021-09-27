package joke_test

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

/// Need to find some workaround for this test
func TestTodayJoke(t *testing.T) {
	req, _ := http.NewRequest("GET", "/today", nil)
	res, err := app.Test(req, int(time.Minute*2))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equalf(t, false, err != nil, "today joke")
	assert.Equalf(t, 200, res.StatusCode, "today joke")
	assert.NotEqualf(t, 0, res.ContentLength, "today joke")
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Nilf(t, err, "today joke")
}

func TestSingleJoke(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res, err := app.Test(req, int(time.Minute*2))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equalf(t, false, err != nil, "single joke")
	assert.Equalf(t, 200, res.StatusCode, "single joke")
	assert.NotEqualf(t, 0, res.ContentLength, "single joke")
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Nilf(t, err, "single joke")
}

func TestJokeByID_200(t *testing.T) {
	req, _ := http.NewRequest("GET", "/id/1", nil)
	res, err := app.Test(req, int(time.Minute*2))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equalf(t, false, err != nil, "joke by id")
	assert.Equalf(t, 200, res.StatusCode, "joke by id")
	assert.NotEqualf(t, 0, res.ContentLength, "joke by id")
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Nilf(t, err, "joke by id")
}

func TestJokeByID_404(t *testing.T) {
	req, _ := http.NewRequest("GET", "/id/300", nil)
	res, err := app.Test(req, int(time.Minute*2))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equalf(t, false, err != nil, "joke by id")
	assert.Equalf(t, 404, res.StatusCode, "joke by id")
	assert.NotEqualf(t, 0, res.ContentLength, "joke by id")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Nilf(t, err, "joke by id")
	assert.Equalf(t, "Requested ID was not found.", string(body), "joke by id")
}
