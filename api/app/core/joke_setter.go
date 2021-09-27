package core

import (
	"context"

	"github.com/allegro/bigcache/v3"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pquerna/ffjson/ffjson"
)

// SetAllJSONJoke fetches jokes data from GetAllJSONJokes then set it to memory cache.
func SetAllJSONJoke(db *pgxpool.Pool, memory *bigcache.BigCache, ctx *context.Context) error {
	jokes, err := GetAllJSONJokes(db, ctx)
	if err != nil {
		return err
	}
	err = memory.Set("jokes", jokes)
	if err != nil {
		return err
	}
	return nil
}

func SetTotalJoke(db *pgxpool.Pool, memory *bigcache.BigCache, ctx *context.Context) error {
	check, err := CheckJokesCache(memory)
	if err != nil {
		return err
	}

	if !check {
		err = SetAllJSONJoke(db, memory, ctx)
		if err != nil {
			return err
		}
	}

	jokes, err := memory.Get("jokes")
	if err != nil {
		return err
	}

	var data []Joke
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
