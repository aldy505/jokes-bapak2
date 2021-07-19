package core

import (
	"context"
	"jokes-bapak2-api/app/v1/models"
	"math/rand"

	"github.com/allegro/bigcache/v3"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pquerna/ffjson/ffjson"
)

// GetAllJSONJokes fetch the database for all the jokes then output it as a JSON []byte.
// Keep in mind, you will need to store it to memory yourself.
func GetAllJSONJokes(db *pgxpool.Pool) ([]byte, error) {
	var jokes []models.Joke
	results, err := db.Query(context.Background(), "SELECT \"id\",\"link\" FROM \"jokesbapak2\" ORDER BY \"id\"")
	if err != nil {
		return nil, err
	}

	err = pgxscan.ScanAll(&jokes, results)
	if err != nil {
		return nil, err
	}

	data, err := ffjson.Marshal(jokes)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetRandomJokeFromCache returns a link string of a random joke from cache.
func GetRandomJokeFromCache(memory *bigcache.BigCache) (string, error) {
	jokes, err := memory.Get("jokes")
	if err != nil {
		if err.Error() == "Entry not found" {
			return "", models.ErrNotFound
		}
		return "", err
	}

	var data []models.Joke
	err = ffjson.Unmarshal(jokes, &data)
	if err != nil {
		return "", nil
	}

	// Return an error if the database is empty
	dataLength := len(data)
	if dataLength == 0 {
		return "", models.ErrEmpty
	}

	random := rand.Intn(dataLength)
	joke := data[random].Link

	return joke, nil
}

// CheckJokesCache checks if there is some value inside jokes cache.
func CheckJokesCache(memory *bigcache.BigCache) (bool, error) {
	_, err := memory.Get("jokes")
	if err != nil {
		if err.Error() == "Entry not found" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// CheckTotalJokesCache literally does what the name is for
func CheckTotalJokesCache(memory *bigcache.BigCache) (bool, error) {
	_, err := memory.Get("total")
	if err != nil {
		if err.Error() == "Entry not found" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// GetCachedJokeByID returns a link string of a certain ID from cache.
func GetCachedJokeByID(memory *bigcache.BigCache, id int) (string, error) {
	jokes, err := memory.Get("jokes")
	if err != nil {
		if err.Error() == "Entry not found" {
			return "", models.ErrNotFound
		}
		return "", err
	}

	var data []models.Joke
	err = ffjson.Unmarshal(jokes, &data)
	if err != nil {
		return "", nil
	}

	// This is a simple solution, might convert it to goroutines and channels sometime soon.
	for _, v := range data {
		if v.ID == id {
			return v.Link, nil
		}
	}

	return "", nil
}

// GetCachedTotalJokes
func GetCachedTotalJokes(memory *bigcache.BigCache) (int, error) {
	total, err := memory.Get("total")
	if err != nil {
		if err.Error() == "Entry not found" {
			return 0, models.ErrNotFound
		}
		return 0, err
	}

	return int(total[0]), nil
}
