package validator

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SubmitLinkExists(db *pgxpool.Pool, ctx context.Context, query squirrel.StatementBuilderType, link string) (bool, error) {
	conn, err := db.Acquire(ctx)
	if err != nil {
		return false, err
	}
	defer conn.Release()

	sql, args, err := query.
		Select("link").
		From("submission").
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

	if err == nil && validateLink != "" {
		return true, nil
	}

	return false, nil
}
