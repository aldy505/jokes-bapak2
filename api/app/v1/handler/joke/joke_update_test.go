package joke_test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUpdateJoke_200(t *testing.T) {
	t.SkipNow()
	err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	reqBody := strings.NewReader("{\"link\":\"https://picsum.photos/id/9/200/300\",\"key\":\"test\",\"token\":\"password\"}")
	req, _ := http.NewRequest("PATCH", "/id/1", reqBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := app.Test(req, int(time.Minute * 2))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equalf(t, false, err != nil, "joke update")
	assert.Equalf(t, 200, res.StatusCode, "joke update")
	assert.NotEqualf(t, 0, res.ContentLength, "joke update")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Nilf(t, err, "joke update")
	assert.Equalf(t, "{\"message\":\"specified joke id has been deleted\"}", string(body), "joke update")
}

func TestUpdateJoke_NotExists(t *testing.T) {
	// TODO: Remove this line below, make this test works
	t.SkipNow()
	err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	reqBody := strings.NewReader("{\"link\":\"https://picsum.photos/id/9/200/300\",\"key\":\"test\",\"token\":\"password\"}")
	req, _ := http.NewRequest("PATCH", "/id/100", reqBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := app.Test(req, int(time.Minute * 2))
if err != nil {
		t.Fatal(err)
	}
	
	assert.Equalf(t, false, err != nil, "joke update")
	assert.Equalf(t, 406, res.StatusCode, "joke update")
	assert.NotEqualf(t, 0, res.ContentLength, "joke update")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Nilf(t, err, "joke update")
	assert.Equalf(t, "{\"message\":\"specified joke id does not exists\"}", string(body), "joke update")
}
