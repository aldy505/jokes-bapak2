package utils_test

import (
	"net/http"
	"testing"

	"github.com/aldy505/jokes-bapak2-api/api/app/v1/utils"
)

func TestRequest(t *testing.T) {
	t.Run("should be able to do a get request", func(t *testing.T) {
		res, err := utils.Request(utils.RequestConfig{
			URL:    "https://jsonplaceholder.typicode.com/todos/1",
			Method: http.MethodGet,
			Headers: map[string]interface{}{
				"User-Agent": "Jokesbapak2 Test API",
				"Accept":     "application/json",
			},
		})
		if err != nil {
			t.Error(err.Error())
		}
		if res.StatusCode != 200 {
			t.Error("response does not have 200 status", res.Status)
		}
	})
}
