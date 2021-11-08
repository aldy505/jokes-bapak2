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
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}

	conn, err := db.Acquire(context.Background())
	if err != nil {
		panic(err)
	}
	defer conn.Release()

	tx, err := conn.Begin(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "DROP TABLE IF EXISTS submission")
	if err != nil {
		panic(err)
	}
	_, err = tx.Exec(context.Background(), "DROP TABLE IF EXISTS jokesbapak2")
	if err != nil {
		panic(err)
	}
	_, err = tx.Exec(context.Background(), "DROP TABLE IF EXISTS administrators")
	if err != nil {
		panic(err)
	}

	_, err = tx.Exec(
		context.Background(),
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

	err = tx.Commit(context.Background())
	if err != nil {
		panic(err)
	}
}

func Teardown() (err error) {
	tx, err := db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "DROP TABLE IF EXISTS submission")
	if err != nil {
		return err
	}
	_, err = tx.Exec(context.Background(), "DROP TABLE IF EXISTS jokesbapak2")
	if err != nil {
		return err
	}
	_, err = tx.Exec(context.Background(), "DROP TABLE IF EXISTS administrators")
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	db.Close()
	return
}

func Flush() error {
	conn, err := db.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), "TRUNCATE TABLE administrators")
	if err != nil {
		return err
	}

	return nil
}
