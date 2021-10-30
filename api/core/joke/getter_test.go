package joke_test

import (
	"context"
	"jokes-bapak2-api/core/joke"
	"testing"

	"github.com/jackc/pgx/v4"
)

func TestGetAllJSONJokes(t *testing.T) {
	defer Teardown()
	conn, err := db.Acquire(context.Background())
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	err = conn.BeginFunc(context.Background(), func(t pgx.Tx) error {
		_, err := t.Exec(context.Background(), "INSERT INTO \"administrators\" (id, key, token, last_used) VALUES ($1, $2, $3, $4), ($5, $6, $7, $8);", administratorsData...)
		if err != nil {
			return err
		}
		_, err = t.Exec(context.Background(), "INSERT INTO \"jokesbapak2\" (id, link, creator) VALUES ($1, $2, $3), ($4, $5, $6), ($7, $8, $9);", jokesData...)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	j, err := joke.GetAllJSONJokes(db, context.Background())
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if string(j) == "" {
		t.Error("j should not be empty")
	}

}

func TestGetRandomJokeFromCache(t *testing.T) {
	defer Teardown()
	//
}

func TestCheckJokesCache(t *testing.T) {
	defer Teardown()
	//
}

func TestCheckTotalJokesCache(t *testing.T) {
	defer Teardown()
	//
}

func TestGetCachedJokeByID(t *testing.T) {
	defer Teardown()
	//
}

func TestGetCachedTotalJokes(t *testing.T) {
	defer Teardown()
	//
}

func TestCheckJokeExists(t *testing.T) {
	defer Teardown()
	//
}
