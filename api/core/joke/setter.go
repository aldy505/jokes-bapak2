package joke

import (
	"context"
	"jokes-bapak2-api/core/schema"

	"github.com/Masterminds/squirrel"
	"github.com/allegro/bigcache/v3"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pquerna/ffjson/ffjson"
)

// SetAllJSONJoke fetches jokes data from GetAllJSONJokes then set it to memory cache.
func SetAllJSONJoke(db *pgxpool.Pool, ctx context.Context, memory *bigcache.BigCache) error {
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

func SetTotalJoke(db *pgxpool.Pool, ctx context.Context, memory *bigcache.BigCache) error {
	check, err := CheckJokesCache(memory)
	if err != nil {
		return err
	}

	if !check {
		err = SetAllJSONJoke(db, ctx, memory)
		if err != nil {
			return err
		}
	}

	jokes, err := memory.Get("jokes")
	if err != nil {
		return err
	}

	var data []schema.Joke
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

func InsertJokeIntoDB(db *pgxpool.Pool, ctx context.Context, joke schema.Joke) error {
	conn, err := db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	var query = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	sql, args, err := query.
		Insert("jokesbapak2").
		Columns("link", "creator").
		Values(joke.Link, joke.Creator).
		ToSql()
	if err != nil {
		return err
	}

	r, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	defer r.Close()
	return nil
}

func DeleteSingleJoke(db *pgxpool.Pool, ctx context.Context, id int) error {
	conn, err := db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	var query = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	sql, args, err := query.
		Delete("jokesbapak2").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	r, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	defer r.Close()

	return nil
}

func UpdateJoke(db *pgxpool.Pool, ctx context.Context, newJoke schema.Joke) error {
	conn, err := db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	var query = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	sql, args, err := query.
		Update("jokesbapak2").
		Set("link", newJoke.Link).
		Set("creator", newJoke.Creator).
		Where(squirrel.Eq{"id": newJoke.ID}).
		ToSql()
	if err != nil {
		return err
	}

	r, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	defer r.Close()

	return nil
}
