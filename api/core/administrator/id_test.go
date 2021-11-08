package administrator_test

import (
	"context"
	"jokes-bapak2-api/core/administrator"
	"testing"
)

func TestGetUserID_Success(t *testing.T) {
	t.Cleanup(func() { Flush() })

	c, err := db.Acquire(context.Background())
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer c.Release()

	_, err = c.Exec(
		context.Background(),
		`INSERT INTO administrators (id, key, token, last_used) VALUES ($1, $2, $3, $4)`,
		administratorsData...,
	)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	id, err := administrator.GetUserID(db, context.Background(), "very secure")
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if id != 1 {
		t.Error("id is not correct, want: 1, got:", id)
	}
}

func TestGetUserID_Failed(t *testing.T) {
	t.Cleanup(func() { Flush() })

	c, err := db.Acquire(context.Background())
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer c.Release()

	id, err := administrator.GetUserID(db, context.Background(), "very secure")
	if err == nil {
		t.Error("an error was expected, got:", id)
	}
}
