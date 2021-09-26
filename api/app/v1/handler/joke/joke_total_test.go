package joke_test

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTotalJokes(t *testing.T) {
	err := setup()
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup()

	req, _ := http.NewRequest("GET", "/total", nil)
	res, err := app.Test(req, int(time.Minute * 2))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equalf(t, false, err != nil, "joke total")
	assert.Equalf(t, 200, res.StatusCode, "joke total")
	assert.NotEqualf(t, 0, res.ContentLength, "joke total")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	
	assert.Nilf(t, err, "joke total")
	// FIXME: This should be "message": "3", not one. I don't know what's wrong as it's 1 AM.
	assert.Equalf(t, "{\"message\":\"3\"}", string(body), "joke total")

}
