package utils_test

import (
	"testing"

	"jokes-bapak2-api/utils"
)

func TestIsIn_True(t *testing.T) {
	arr := []string{"John", "Matthew", "Thomas", "Adam"}
	check := utils.IsIn(arr, "Thomas")
	if !check {
		t.Error("check should be true: ", check)
	}
}

func TestIsIn_False(t *testing.T) {
	arr := []string{"John", "Matthew", "Thomas", "Adam"}
	check := utils.IsIn(arr, "James")
	if check {
		t.Error("check should be false: ", check)
	}
}
