package utils_test

import (
	"testing"
	"time"

	"jokes-bapak2-api/app/v1/utils"
)

func TestIsToday(t *testing.T) {
	t.Run("should be able to tell if it's today", func(t *testing.T) {
		today, err := utils.IsToday(time.Now().UTC().Format(time.RFC3339))
		if err != nil {
			t.Error(err.Error())
		}
		if today == false {
			t.Error("today should be true:", today)
		}
	})

	t.Run("should be able to tell if it's not today", func(t *testing.T) {
		today, err := utils.IsToday("2021-01-01T11:48:24Z")
		if err != nil {
			t.Error(err.Error())
		}
		if today == true {
			t.Error("today should be false:", today)
		}
	})
}
