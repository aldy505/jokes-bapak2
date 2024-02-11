package utils_test

import (
	"testing"

	"jokes-bapak2-api/utils"
)

func TestRandomString_Valid(t *testing.T) {
	random, err := utils.RandomString(10)
	if err != nil {
		t.Error(err)
	}
	if len(random) != 20 {
		t.Error("result is not within the length of 10")
	}
}

func TestRandomString_Invalid(t *testing.T) {
	random, err := utils.RandomString(-10)
	if err != nil {
		t.Error(err)
	}
	if len(random) != 20 {
		t.Error("result is not within the length of 10")
	}
}
