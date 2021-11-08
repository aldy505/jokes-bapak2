package validator_test

import (
	"jokes-bapak2-api/core/validator"
	"testing"

	"github.com/gojek/heimdall/v7/httpclient"
)

func TestCheckImageValidity_Error(t *testing.T) {
	client := httpclient.NewClient()
	b, err := validator.CheckImageValidity(client, "http://lorem-ipsum")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	if b {
		t.Error("Expected false, got true")
	}

	if err.Error() != "URL must use HTTPS protocol" {
		t.Error("Expected error to be URL must use HTTPS protocol, got:", err)
	}
}

func TestCheckImageValidity_False(t *testing.T) {
	client := httpclient.NewClient()

	b, err := validator.CheckImageValidity(client, "https://www.youtube.com/watch?v=yTJV6T37Reo")
	if err != nil {
		t.Error("Expected nil, got error")
	}

	if b {
		t.Error("Expected false, got true")
	}
}

func TestCheckImageValidity_True(t *testing.T) {
	client := httpclient.NewClient()

	b, err := validator.CheckImageValidity(client, "https://i.ytimg.com/vi/yTJV6T37Reo/maxresdefault.jpg")
	if err != nil {
		t.Error("Expected nil, got error")
	}

	if !b {
		t.Error("Expected true, got false")
	}
}
