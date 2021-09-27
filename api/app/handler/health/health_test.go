package health_test

import (
	"io/ioutil"
	v1 "jokes-bapak2-api/app"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

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
