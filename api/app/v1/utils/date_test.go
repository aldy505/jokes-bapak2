package utils_test

import (
	"testing"
	"time"

	"jokes-bapak2-api/app/v1/utils"
)

func TestIsToday_Today(t *testing.T) {
	today, err := utils.IsToday(time.Now().Format(time.RFC3339))
	if err != nil {
		t.Error(err.Error())
	}
	if today == false {
		t.Error("today should be true:", today)
	}
}

func TestIsToday_NotToday(t *testing.T) {
	today, err := utils.IsToday("2021-01-01T11:48:24Z")
	if err != nil {
		t.Error(err.Error())
	}
	if today == true {
		t.Error("today should be false:", today)
	}
}

func TestIsToday_ErrorIfEmpty(t *testing.T) {
	today, err := utils.IsToday("")
	if err != nil {
		t.Error(err.Error())
	}
	if today != false {
		t.Error("it should be false:", today)
	}
}

func TestIsToday_ErrorIfInvalid(t *testing.T) {
	today, err := utils.IsToday("asdfghjkl")
	if err == nil {
		t.Error("it should be error:", today, err)
	}
	if today != false {
		t.Error("it should be false:", today)
	}
}
