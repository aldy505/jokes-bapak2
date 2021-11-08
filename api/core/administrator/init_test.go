package administrator_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

var administratorsData = []interface{}{
	1, "very secure", "not the real one", time.Now().Format(time.RFC3339),
}

func TestMain(m *testing.M) {
	defer Teardown()
	Setup()

	os.Exit(m.Run())
}

func Setup() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Minute))
	defer cancel()

	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	db, err = pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		panic(err)
	}

	conn, err := db.Acquire(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(
		ctx,
		`CREATE TABLE IF NOT EXISTS administrators (
			id SERIAL PRIMARY KEY,
			key VARCHAR(255) NOT NULL UNIQUE,
			token TEXT,
			last_used VARCHAR(255)
		)`,
	)
	if err != nil {
		panic(err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		panic(err)
	}
}

func Teardown() (err error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Minute))
	defer cancel()

	defer db.Close()
	
	c, err := db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer c.Release()

	tx, err := c.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "TRUNCATE TABLE submission RESTART IDENTITY CASCADE")
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, "TRUNCATE TABLE jokesbapak2 RESTART IDENTITY CASCADE")
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, "TRUNCATE TABLE administrators RESTART IDENTITY CASCADE")
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return
}

func Flush() error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	conn, err := db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "TRUNCATE TABLE administrators RESTART IDENTITY CASCADE")
	if err != nil {
		return err
	}

	return nil
}
