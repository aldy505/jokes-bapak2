package utils_test

import (
	"jokes-bapak2-api/app/v1/utils"
	"testing"
)

func TestRandomString(t *testing.T) {
	t.Run("should create a random string with param", func(t *testing.T) {
		random, err := utils.RandomString(10)
		if err != nil {
			t.Error(err)
		}
		if len(random) != 20 {
			t.Error("result is not within the length of 10")
		}
	})
	t.Run("should create a random string with invalid params", func(t *testing.T) {
		random, err := utils.RandomString(10)
		if err != nil {
			t.Error(err)
		}
		if len(random) != 20 {
			t.Error("result is not within the length of 10")
		}
	})
}
