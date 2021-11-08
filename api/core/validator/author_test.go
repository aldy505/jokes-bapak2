package validator_test

import (
	"jokes-bapak2-api/core/validator"
	"testing"
)

func TestValidateAuthor_False(t *testing.T) {
	v := validator.ValidateAuthor("Test Author")
	if v != false {
		t.Error("Expected false, got true")
	}

	v = validator.ValidateAuthor("Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec.")
	if v != false {
		t.Error("Expected false, got true")
	}

	v = validator.ValidateAuthor("")
	if v != false {
		t.Error("Expected false, got true")
	}

	v = validator.ValidateAuthor("Test <Author")
	if v != false {
		t.Error("Expected false, got true")
	}

	v = validator.ValidateAuthor("Test <Author>")
	if v != false {
		t.Error("Expected false, got true")
	}
}

func TestValidateAuthor_True(t *testing.T) {
	v := validator.ValidateAuthor("Test Author <author@mail.com>")
	if v != true {
		t.Error("Expected true, got false")
	}
}
