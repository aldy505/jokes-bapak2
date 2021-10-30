package joke

import (
	"context"
	"errors"
	"jokes-bapak2-api/core/schema"
	"math/rand"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/allegro/bigcache/v3"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pquerna/ffjson/ffjson"
)

// GetAllJSONJokes fetch the database for all the jokes then output it as a JSON []byte.
// Keep in mind, you will need to store it to memory yourself.
func GetAllJSONJokes(db *pgxpool.Pool, ctx context.Context) ([]byte, error) {
	conn, err := db.Acquire(ctx)
	if err != nil {
		return []byte{}, err
	}
	defer conn.Release()

	var jokes []schema.Joke
	results, err := conn.Query(context.Background(), "SELECT \"id\",\"link\" FROM \"jokesbapak2\" ORDER BY \"id\"")
	if err != nil {
		return nil, err
	}
	defer results.Close()

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

// Only return a link
func GetRandomJokeFromDB(db *pgxpool.Pool, ctx context.Context) (string, error) {
	conn, err := db.Acquire(ctx)
	if err != nil {
		return "", err
	}

	var link string
	err = conn.QueryRow(context.Background(), "SELECT link FROM jokesbapak2 ORDER BY random() LIMIT 1").Scan(&link)
	if err != nil {
		return "", err
	}

	return link, nil
}

// GetRandomJokeFromCache returns a link string of a random joke from cache.
func GetRandomJokeFromCache(memory *bigcache.BigCache) (string, error) {
	jokes, err := memory.Get("jokes")
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return "", schema.ErrNotFound
		}
		return "", err
	}

	var data []schema.Joke
	err = ffjson.Unmarshal(jokes, &data)
	if err != nil {
		return "", nil
	}

	// Return an error if the database is empty
	dataLength := len(data)
	if dataLength == 0 {
		return "", schema.ErrEmpty
	}

	random := rand.Intn(dataLength)
	joke := data[random].Link

	return joke, nil
}

// CheckJokesCache checks if there is some value inside jokes cache.
func CheckJokesCache(memory *bigcache.BigCache) (bool, error) {
	_, err := memory.Get("jokes")
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
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
		if errors.Is(err, bigcache.ErrEntryNotFound) {
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
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return "", schema.ErrNotFound
		}
		return "", err
	}

	var data []schema.Joke
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
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return 0, schema.ErrNotFound
		}
		return 0, err
	}

	return int(total[0]), nil
}

func CheckJokeExists(db *pgxpool.Pool, ctx context.Context, id string) (bool, error) {
	conn, err := db.Acquire(ctx)
	if err != nil {
		return false, err
	}
	defer conn.Release()

	var query = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	sql, args, err := query.
		Select("id").
		From("jokesbapak2").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return false, err
	}

	var jokeID int
	err = conn.QueryRow(context.Background(), sql, args...).Scan(&jokeID)
	if err != nil && err != pgx.ErrNoRows {
		return false, err
	}

	return strconv.Itoa(jokeID) == id, nil
}
