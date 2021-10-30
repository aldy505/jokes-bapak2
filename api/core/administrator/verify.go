package administrator

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CheckKeyExists(db *pgxpool.Pool, ctx context.Context, key string) (string, error) {
	conn, err := db.Acquire(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Release()

	var query = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// Check if key exists
	sql, args, err := query.
		Select("token").
		From("administrators").
		Where(squirrel.Eq{"key": key}).
		ToSql()
	if err != nil {
		return "", err
	}

	var token string
	err = conn.QueryRow(context.Background(), sql, args...).Scan(&token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return token, nil
}
