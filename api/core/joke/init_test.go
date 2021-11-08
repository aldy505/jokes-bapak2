package joke_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool
var cache *redis.Client
var memory *bigcache.BigCache

var jokesData = []interface{}{
	1, "https://via.placeholder.com/300/06f/fff.png", 1,
	2, "https://via.placeholder.com/300/07f/fff.png", 1,
	3, "https://via.placeholder.com/300/08f/fff.png", 1,
}
var administratorsData = []interface{}{
	1, "very secure", "not the real one", time.Now().Format(time.RFC3339), 2, "test", "$argon2id$v=19$m=65536,t=16,p=4$3a08c79fbf2222467a623df9a9ebf75802c65a4f9be36eb1df2f5d2052d53cb7$ce434bd38f7ba1fc1f2eb773afb8a1f7f2dad49140803ac6cb9d7256ce9826fb3b4afa1e2488da511c852fc6c33a76d5657eba6298a8e49d617b9972645b7106", "",
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

	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}

	cache = redis.NewClient(opt)

	memory, err = bigcache.NewBigCache(bigcache.DefaultConfig(6 * time.Hour))
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

	// Dropping all table first
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
		);`,
	)
	if err != nil {
		panic(err)
	}
	_, err = tx.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS jokesbapak2 (
			id SERIAL PRIMARY		 KEY,
			link TEXT UNIQUE,
			creator INT NOT NULL REFERENCES "administrators" ("id")
		);`,
	)
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

	err = cache.Close()
	if err != nil {
		return
	}
	err = memory.Close()
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
	_, err = conn.Exec(context.Background(), "TRUNCATE TABLE jokesbapak2")
	if err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(), "TRUNCATE TABLE administrators")
	if err != nil {
		return err
	}

	err = cache.FlushAll(context.Background()).Err()
	if err != nil {
		return err
	}

	err = memory.Reset()
	if err != nil {
		return err
	}

	return nil
}
