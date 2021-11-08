package utils_test

import (
	"jokes-bapak2-api/utils"
	"strings"
	"testing"
)

func TestParseToJSONBody(t *testing.T) {
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
}

func TestParseToFormBody(t *testing.T) {
	body := map[string]interface{}{
		"age":  32,
		"fat":  true,
		"name": "Scott",
	}
	parsed, err := utils.ParseToFormBody(body)
	if err != nil {
		t.Error(err.Error())
	}
	result := [3]string{"age=32&", "fat=true&", "name=Scott&"}
	if !strings.Contains(string(parsed), result[0]) && !strings.Contains(string(parsed), result[1]) && !strings.Contains(string(parsed), result[2]) {
		t.Error("parsed string is not the same as result:", string(parsed))
	}
}
