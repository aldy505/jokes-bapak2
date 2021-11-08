package submit_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

var submissionData = []interface{}{
	1, "https://via.placeholder.com/300/01f/fff.png", "2021-08-03T18:20:38Z", "Test <test@example.com>", 0,
	2, "https://via.placeholder.com/300/02f/fff.png", "2021-08-04T18:20:38Z", "Test <test@example.com>", 1,
}

func TestMain(m *testing.M) {
	defer Teardown()
	Setup()
	time.Sleep(3 * time.Second)

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
		`CREATE TABLE IF NOT EXISTS submission (
			id SERIAL PRIMARY KEY,
			link VARCHAR(255) UNIQUE NOT NULL,
			created_at VARCHAR(255),
			author VARCHAR(255) NOT NULL,
			status SMALLINT DEFAULT 0
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
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	defer db.Close()

	conn, err := db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "DROP TABLE IF EXISTS submission CASCADE")
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, "DROP TABLE IF EXISTS jokesbapak2 CASCADE")
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, "DROP TABLE IF EXISTS administrators CASCADE")
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

	_, err = conn.Exec(ctx, "TRUNCATE TABLE submission RESTART IDENTITY CASCADE")
	if err != nil {
		return err
	}

	return nil
}
