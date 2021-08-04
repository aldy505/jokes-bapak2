package joke_test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateJoke_200(t *testing.T) {
	t.SkipNow()
	err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	reqBody := strings.NewReader("{\"link\":\"https://picsum.photos/id/9/200/300\",\"key\":\"very secure\",\"token\":\"password\"}")
	req, _ := http.NewRequest("PATCH", "/id/1", reqBody)
	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "joke update")
	assert.Equalf(t, 200, res.StatusCode, "joke update")
	assert.NotEqualf(t, 0, res.ContentLength, "joke update")
	body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "joke update")
	assert.Equalf(t, "{\"message\":\"specified joke id has been deleted\"}", string(body), "joke update")
}

func TestUpdateJoke_NotExists(t *testing.T) {
	t.SkipNow()
	err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	reqBody := strings.NewReader("{\"link\":\"https://picsum.photos/id/9/200/300\",\"key\":\"very secure\",\"token\":\"password\"}")
	req, _ := http.NewRequest("PATCH", "/id/100", reqBody)
	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "joke update")
	assert.Equalf(t, 406, res.StatusCode, "joke update")
	assert.NotEqualf(t, 0, res.ContentLength, "joke update")
	body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "joke update")
	assert.Equalf(t, "{\"message\":\"specified joke id does not exists\"}", string(body), "joke update")
}
