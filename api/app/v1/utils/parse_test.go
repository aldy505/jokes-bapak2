package utils_test

import (
	"testing"

	"jokes-bapak2-api/app/v1/utils"
)

func TestParseToJSONBody(t *testing.T) {
	t.Run("should be able to parse a json string", func(t *testing.T) {
		body := map[string]interface{}{
			"name": "Scott",
			"age":  32,
			"fat":  true,
		}
		parsed, err := utils.ParseToJSONBody(body)
		if err != nil {
			t.Error(err.Error())
		}
		result := "{\"age\":32,\"fat\":true,\"name\":\"Scott\"}"
		if string(parsed) != result {
			t.Error("parsed string is not the same as result:", string(parsed))
		}
	})
}

func TestParseToFormBody(t *testing.T) {
	t.Run("should be able to parse a form body", func(t *testing.T) {
		body := map[string]interface{}{
			"name": "Scott",
			"age":  32,
			"fat":  true,
		}
		parsed, err := utils.ParseToFormBody(body)
		if err != nil {
			t.Error(err.Error())
		}
		result := "name=Scott&age=32&fat=true&"
		if string(parsed) != result {
			t.Error("parsed string is not the same as result:", string(parsed))
		}
	})
}
