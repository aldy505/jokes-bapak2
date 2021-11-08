package submit_test

import (
	"context"
	"os"
	"testing"

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
		`CREATE TABLE IF NOT EXISTS submission (
			id SERIAL PRIMARY KEY,
			link VARCHAR(255) UNIQUE NOT NULL,
			created_at VARCHAR(255),
			author VARCHAR(255) NOT NULL,
			status SMALLINT DEFAULT 0
		);`,
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

	_, err = conn.Exec(context.Background(), "TRUNCATE TABLE submission")
	if err != nil {
		return err
	}

	return nil
}
