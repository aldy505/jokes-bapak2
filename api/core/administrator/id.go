package administrator

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetUserID(db *pgxpool.Pool, ctx context.Context, key string) (int, error) {
	var query = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	c, err := db.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer c.Release()

	tx, err := c.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	sql, args, err := query.
		Update("administrators").
		Set("last_used", time.Now().UTC().Format(time.RFC3339)).
		ToSql()
	if err != nil {
		return 0, err
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	sql, args, err = query.
		Select("id").
		From("administrators").
		Where(squirrel.Eq{"key": key}).
		ToSql()
	if err != nil {
		return 0, err
	}

	var id int
	err = tx.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return id, nil
}
