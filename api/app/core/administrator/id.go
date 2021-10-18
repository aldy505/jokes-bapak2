package administrator

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetUserID(db *pgxpool.Pool, ctx context.Context, key string) (int, error) {
	var query = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	c1, err := db.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer c1.Release()

	sql, args, err := query.
		Update("administrators").
		Set("last_used", time.Now().UTC().Format(time.RFC3339)).
		ToSql()
	if err != nil {
		return 0, err
	}

	r, err := c1.Query(context.Background(), sql, args...)
	if err != nil {
		return 0, err
	}
	defer r.Close()

	c2, err := db.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer c2.Release()

	sql, args, err = query.
		Select("id").
		From("administrators").
		Where(squirrel.Eq{"key": key}).
		ToSql()
	if err != nil {
		return 0, err
	}

	var id int
	err = c2.QueryRow(context.Background(), sql, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
