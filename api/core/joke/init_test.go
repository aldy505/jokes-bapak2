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

	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}

	cache = redis.NewClient(opt)

	memory, err = bigcache.NewBigCache(bigcache.DefaultConfig(6 * time.Hour))
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

	_, err = tx.Exec(
		ctx,
		`CREATE TABLE IF NOT EXISTS jokesbapak2 (
			id SERIAL PRIMARY KEY,
			link TEXT UNIQUE,
			creator INT NOT NULL REFERENCES "administrators" ("id")
		)`,
	)
	if err != nil {
		panic(err)
	}
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

	err = cache.Close()
	if err != nil {
		return
	}
	err = memory.Close()
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
	_, err = conn.Exec(ctx, "TRUNCATE TABLE jokesbapak2 RESTART IDENTITY CASCADE")
	if err != nil {
		return err
	}
	_, err = conn.Exec(ctx, "TRUNCATE TABLE administrators RESTART IDENTITY CASCADE")
	if err != nil {
		return err
	}

	err = cache.FlushAll(ctx).Err()
	if err != nil {
		return err
	}

	err = memory.Reset()
	if err != nil {
		return err
	}

	return nil
}
