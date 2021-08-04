package joke_test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteJoke_200(t *testing.T) {
	// TODO: Remove this line below, make this test works
	t.SkipNow()
	err := setup()
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup()

	reqBody := strings.NewReader("{\"key\":\"very secure\",\"token\":\"password\"}")
	req, _ := http.NewRequest("DELETE", "/id/1", reqBody)
	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "joke delete")
	assert.Equalf(t, 200, res.StatusCode, "joke delete")
	assert.NotEqualf(t, 0, res.ContentLength, "joke delete")
	body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "joke delete")
	assert.Equalf(t, "{\"message\":\"specified joke id has been deleted\"}", string(body), "joke delete")
}
func TestDeleteJoke_NotExists(t *testing.T) {
	t.SkipNow()
	err := setup()
	if err != nil {
		t.Fatal(err)
	}

	defer cleanup()

	reqBody := strings.NewReader("{\"key\":\"very secure\",\"token\":\"password\"}")
	req, _ := http.NewRequest("DELETE", "/id/100", reqBody)
	res, err := app.Test(req, -1)

	assert.Equalf(t, false, err != nil, "joke delete")
	assert.Equalf(t, 406, res.StatusCode, "joke delete")
	assert.NotEqualf(t, 0, res.ContentLength, "joke delete")
	body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "joke delete")
	assert.Equalf(t, "{\"message\":\"specified joke id does not exists\"}", string(body), "joke delete")
}
