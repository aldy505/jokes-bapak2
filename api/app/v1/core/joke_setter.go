package core

import (
	"jokes-bapak2-api/app/v1/models"

	"github.com/allegro/bigcache/v3"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pquerna/ffjson/ffjson"
)

// SetAllJSONJoke fetches jokes data from GetAllJSONJokes then set it to memory cache.
func SetAllJSONJoke(db *pgxpool.Pool, memory *bigcache.BigCache) error {
	jokes, err := GetAllJSONJokes(db)
	if err != nil {
		return err
	}
	err = memory.Set("jokes", jokes)
	if err != nil {
		return err
	}
	return nil
}

func SetTotalJoke(db *pgxpool.Pool, memory *bigcache.BigCache) error {
	check, err := CheckJokesCache(memory)
	if err != nil {
		return err
	}

	if !check {
		return models.ErrEmpty
	}

	err = SetAllJSONJoke(db, memory)
	if err != nil {
		return err
	}

	jokes, err := memory.Get("jokes")
	if err != nil {
		return err
	}

	var data []models.Joke
	err = ffjson.Unmarshal(jokes, &data)
	if err != nil {
		return err
	}

	var total = []byte{byte(len(data))}
	err = memory.Set("total", total)
	if err != nil {
		return err
	}

	return nil
}
