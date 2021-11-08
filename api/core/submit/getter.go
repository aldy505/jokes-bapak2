package submit

import (
	"context"
	"jokes-bapak2-api/core/schema"
	"net/url"
	"strconv"
	"strings"

	"github.com/aldy505/bob"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetSubmittedItems(db *pgxpool.Pool, ctx context.Context, queries schema.SubmissionQuery) ([]schema.Submission, error) {
	var err error
	var limit int
	var offset int
	var approved bool

	if queries.Limit != "" {
		limit, err = strconv.Atoi(queries.Limit)
		if err != nil {
			return []schema.Submission{}, err

		}
	}
	if queries.Page != "" {
		page, err := strconv.Atoi(queries.Page)
		if err != nil {
			return []schema.Submission{}, err

		}
		offset = (page - 1) * 20
	}

	if queries.Approved != "" {
		approved, err = strconv.ParseBool(queries.Approved)
		if err != nil {
			return []schema.Submission{}, err

		}
	}

	var status int

	if approved {
		status = 1
	} else {
		status = 0
	}

	sql, args, err := GetterQueryBuilder(queries, status, limit, offset)
	if err != nil {
		return []schema.Submission{}, err

	}

	conn, err := db.Acquire(ctx)
	if err != nil {
		return []schema.Submission{}, err
	}
	defer conn.Release()

	var submissions []schema.Submission
	results, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return []schema.Submission{}, err
	}
	defer results.Close()

	err = pgxscan.ScanAll(&submissions, results)
	if err != nil {
		return []schema.Submission{}, err
	}

	return submissions, nil
}

func GetterQueryBuilder(queries schema.SubmissionQuery, status, limit, offset int) (string, []interface{}, error) {
	var sql string
	var args []interface{}
	var sqlQuery strings.Builder

	sqlQuery.WriteString("SELECT * FROM submission WHERE TRUE")

	if queries.Author != "" {
		sqlQuery.WriteString(" AND author = ?")
		escapedAuthor, err := url.QueryUnescape(queries.Author)
		if err != nil {
			return sql, args, err

		}
		args = append(args, escapedAuthor)
	}

	if queries.Approved != "" {
		sqlQuery.WriteString(" AND status = ?")
		args = append(args, status)
	}

	if limit > 0 {
		sqlQuery.WriteString(" LIMIT " + strconv.Itoa(limit))
	} else {
		sqlQuery.WriteString(" LIMIT 20")
	}

	if queries.Page != "" {
		sqlQuery.WriteString(" OFFSET " + strconv.Itoa(offset))
	}

	sql = bob.ReplacePlaceholder(sqlQuery.String(), bob.Dollar)

	return sql, args, nil
}
