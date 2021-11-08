package administrator_test

import (
	"context"
	"jokes-bapak2-api/core/administrator"
	"testing"
	"time"
)

func TestGetUserID_Success(t *testing.T) {
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
		`INSERT INTO administrators (id, key, token, last_used) VALUES ($1, $2, $3, $4)`,
		administratorsData...,
	)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	id, err := administrator.GetUserID(db, ctx, "very secure")
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if id != 1 {
		t.Error("id is not correct, want: 1, got:", id)
	}
}

func TestGetUserID_Failed(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	defer Flush()

	c, err := db.Acquire(ctx)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer c.Release()

	id, err := administrator.GetUserID(db, ctx, "very secure")
	if err == nil {
		t.Error("an error was expected, got:", id)
	}
}
