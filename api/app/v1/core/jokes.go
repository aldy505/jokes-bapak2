package core

import (
	"context"
	"encoding/json"
	"jokes-bapak2-api/app/v1/models"
	"math/rand"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/patrickmn/go-cache"
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

	data, err := json.Marshal(jokes)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetRandomJokeFromCache returns a link string of a random joke from cache.
func GetRandomJokeFromCache(memory *cache.Cache) (string, error) {
	jokes, found := memory.Get("jokes")
	if !found {
		return "", models.ErrNotFound
	}

	var data []models.Joke
	err := json.Unmarshal(jokes.([]byte), &data)
	if err != nil {
		return "", nil
	}

	dataLength := len(data)
	if dataLength == 0 {
		return "", models.ErrEmpty
	}

	random := rand.Intn(dataLength)
	joke := data[random].Link

	return joke, nil
}

// CheckJokesCache checks if there is some value inside jokes cache.
func CheckJokesCache(memory *cache.Cache) bool {
	_, found := memory.Get("jokes")
	return found
}

// GetCachedJokeByID returns a link string of a certain ID from cache.
func GetCachedJokeByID(memory *cache.Cache, id int) (string, error) {
	jokes, found := memory.Get("jokes")
	if !found {
		return "", models.ErrNotFound
	}

	var data []models.Joke
	err := json.Unmarshal(jokes.([]byte), &data)
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
