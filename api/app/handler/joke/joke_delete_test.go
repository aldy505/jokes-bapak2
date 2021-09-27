package joke_test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDeleteJoke_200(t *testing.T) {
	// TODO: Remove this line below, make this test works
	t.SkipNow()

	reqBody := strings.NewReader("{\"key\":\"very secure\",\"token\":\"password\"}")
	req, _ := http.NewRequest("DELETE", "/id/100", reqBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := app.Test(req, int(time.Minute*2))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equalf(t, false, err != nil, "joke delete")
	assert.Equalf(t, 200, res.StatusCode, "joke delete")
	assert.NotEqualf(t, 0, res.ContentLength, "joke delete")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Nilf(t, err, "joke delete")
	assert.Equalf(t, "{\"message\":\"specified joke id has been deleted\"}", string(body), "joke delete")
}
func TestDeleteJoke_NotExists(t *testing.T) {
	// TODO: Remove this line below, make this test works
	t.SkipNow()

	reqBody := strings.NewReader("{\"key\":\"very secure\",\"token\":\"password\"}")
	req, _ := http.NewRequest("DELETE", "/id/100", reqBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := app.Test(req, int(time.Minute*2))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equalf(t, false, err != nil, "joke delete")
	assert.Equalf(t, 406, res.StatusCode, "joke delete")
	assert.NotEqualf(t, 0, res.ContentLength, "joke delete")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Nilf(t, err, "joke delete")
	assert.Equalf(t, "{\"message\":\"specified joke id does not exists\"}", string(body), "joke delete")
}
