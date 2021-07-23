package utils_test

import (
	"jokes-bapak2-api/app/v1/utils"
	"testing"
)

func TestIsIn(t *testing.T) {
	arr := []string{"John", "Matthew", "Thomas", "Adam"}
	t.Run("should return true", func(t *testing.T) {
		check := utils.IsIn(arr, "Thomas")
		if !check {
			t.Error("check should be true: ", check)
		}
	})

	t.Run("should return false", func(t *testing.T) {
		check := utils.IsIn(arr, "James")
		if check {
			t.Error("check should be false: ", check)
		}
	})
}