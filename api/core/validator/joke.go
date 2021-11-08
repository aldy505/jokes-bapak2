package validator

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Validate if link already exists
func JokeLinkExists(db *pgxpool.Pool, ctx context.Context, link string) (bool, error) {
	conn, err := db.Acquire(ctx)
	if err != nil {
		return false, err
	}
	defer conn.Release()

	var query = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	sql, args, err := query.
		Select("link").
		From("jokesbapak2").
		Where(squirrel.Eq{"link": link}).
		ToSql()
	if err != nil {
		return false, err
	}

	var validateLink string
	err = conn.QueryRow(ctx, sql, args...).Scan(&validateLink)
	if err != nil && err != pgx.ErrNoRows {
		return false, err
	}

	return validateLink != "", nil
}

// Check if the joke exists
func JokeIDExists(db *pgxpool.Pool, ctx context.Context, id int) (bool, error) {
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
	err = conn.QueryRow(ctx, sql, args...).Scan(&jokeID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return false, err
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}

	return true, nil
}
