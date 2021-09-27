package main_test

import (
	"context"
	"errors"
	"flag"
	"jokes-bapak2-api/app/platform/database"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var jokesData = []interface{}{1, "https://via.placeholder.com/300/06f/fff.png", 1, 2, "https://via.placeholder.com/300/07f/fff.png", 1, 3, "https://via.placeholder.com/300/08f/fff.png", 1}
var submissionData = []interface{}{1, "https://via.placeholder.com/300/01f/fff.png", "2021-08-03T18:20:38Z", "Test <test@example.com>", 0, 2, "https://via.placeholder.com/300/02f/fff.png", "2021-08-04T18:20:38Z", "Test <test@example.com>", 1}
var administratorsData = []interface{}{1, "very secure", "not the real one", time.Now().Format(time.RFC3339), 2, "test", "$argon2id$v=19$m=65536,t=16,p=4$3a08c79fbf2222467a623df9a9ebf75802c65a4f9be36eb1df2f5d2052d53cb7$ce434bd38f7ba1fc1f2eb773afb8a1f7f2dad49140803ac6cb9d7256ce9826fb3b4afa1e2488da511c852fc6c33a76d5657eba6298a8e49d617b9972645b7106", ""}

var db *pgxpool.Pool
var ctx context.Context = context.Background()

func TestMain(m *testing.M) {
	flag.Parse()
	err := setup()
	if err != nil {
		log.Fatalln(err)
	}

	code := m.Run()
	os.Exit(code)
}

func cleanup() {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Unable to create pool config", err)
	}
	poolConfig.MaxConns = 3
	poolConfig.MinConns = 1

	db, err = pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalln("Unable to create connection", err)
	}
	defer db.Close()

	d, err := db.Query(ctx, "BEGIN; DROP TABLE \"jokesbapak2\"; DROP TABLE \"administrators\"; DROP TABLE \"submission\"; COMMIT;")
	if err != nil {
		panic(err)
	}
	defer d.Close()
}

func setup() error {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return errors.New("Unable to create pool config: " + err.Error())
	}
	poolConfig.MaxConns = 3
	poolConfig.MinConns = 1

	db, err = pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return errors.New("Unable to create connection: " + err.Error())
	}
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

	_, err = tx.Exec(ctx, "DROP TABLE IF EXISTS \"jokesbapak2\"")
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, "DROP TABLE IF EXISTS \"administrators\"")
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, "DROP TABLE IF EXISTS \"submission\"")
	if err != nil {
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	err = database.Setup(db, &ctx)
	if err != nil {
		return err
	}

	tx, err = conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "INSERT INTO \"administrators\" (id, key, token, last_used) VALUES ($1, $2, $3, $4), ($5, $6, $7, $8);", administratorsData...)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "INSERT INTO \"jokesbapak2\" (id, link, creator) VALUES ($1, $2, $3), ($4, $5, $6), ($7, $8, $9);", jokesData...)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "INSERT INTO \"submission\" (id, link, created_at, author, status) VALUES ($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10);", submissionData...)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
