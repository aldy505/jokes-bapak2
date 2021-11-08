package administrator_test

import (
	"context"
	"jokes-bapak2-api/core/administrator"
	"testing"
	"time"
)

func TestCheckKeyExists_Success(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	defer Flush()

	c, err := db.Acquire(ctx)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer c.Release()

	_, err = c.Exec(
		ctx,
		"INSERT INTO administrators (id, key, token, last_used) VALUES ($1, $2, $3, $4)",
		administratorsData...,
	)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	key, err := administrator.CheckKeyExists(db, ctx, "very secure")
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if key != "not the real one" {
		t.Error("key isn't not the real one, got:", key)
	}
}

func TestCheckKeyExists_Failing(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	defer Flush()

	c, err := db.Acquire(ctx)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer c.Release()

	_, err = c.Exec(
		ctx,
		"INSERT INTO administrators (id, key, token, last_used) VALUES ($1, $2, $3, $4)",
		administratorsData...,
	)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	key, err := administrator.CheckKeyExists(db, ctx, "others")
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if key != "" {
		t.Error("key is not empty, got:", key)
	}
}
