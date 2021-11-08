package administrator_test

import (
	"context"
	"jokes-bapak2-api/core/administrator"
	"testing"
)

func TestCheckKeyExists_Success(t *testing.T) {
	t.Cleanup(func() {
		Flush()
	})

	c, err := db.Acquire(context.Background())
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer c.Release()

	_, err = c.Exec(
		context.Background(),
		"INSERT INTO administrators (id, key, token, last_used) VALUES ($1, $2, $3, $4)",
		administratorsData...,
	)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	key, err := administrator.CheckKeyExists(db, context.Background(), "very secure")
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if key != "not the real one" {
		t.Error("key isn't not the real one, got:", key)
	}
}

func TestCheckKeyExists_Failing(t *testing.T) {
	t.Cleanup(func() {
		Flush()
	})

	c, err := db.Acquire(context.Background())
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer c.Release()

	_, err = c.Exec(
		context.Background(),
		"INSERT INTO administrators (id, key, token, last_used) VALUES ($1, $2, $3, $4)",
		administratorsData...,
	)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	key, err := administrator.CheckKeyExists(db, context.Background(), "others")
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if key != "" {
		t.Error("key is not empty, got:", key)
	}
}
