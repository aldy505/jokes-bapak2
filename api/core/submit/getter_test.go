package submit_test

import (
	"context"
	"jokes-bapak2-api/core/schema"
	"jokes-bapak2-api/core/submit"
	"testing"
)

func TestGetSubmittedItems(t *testing.T) {
	defer Flush()

	c, err := db.Acquire(context.Background())
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer c.Release()

	_, err = c.Exec(context.Background(), "INSERT INTO submission (id, link, created_at, author, status) VALUES ($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10)", submissionData...)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	items, err := submit.GetSubmittedItems(db, context.Background(), schema.SubmissionQuery{})
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if len(items) != 2 {
		t.Error("expected 2 items, got", len(items))
	}
}

func TestGetterQueryBuilder(t *testing.T) {
	s, _, err := submit.GetterQueryBuilder(schema.SubmissionQuery{}, 0, 0, 0)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if s != "SELECT * FROM submission WHERE TRUE LIMIT 20" {
		t.Error("expected query to be", "SELECT * FROM submission WHERE TRUE LIMIT 20", "got", s)
	}

	s, i, err := submit.GetterQueryBuilder(schema.SubmissionQuery{
		Author:   "Test <example@test.com>",
		Approved: "true",
		Page:     "2",
	}, 2, 15, 10)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if s != "SELECT * FROM submission WHERE TRUE AND author = $1 AND status = $2 LIMIT 15 OFFSET 10" {
		t.Error("expected query to be", "SELECT * FROM submission WHERE TRUE AND author = $1 AND status = $2 LIMIT 15 OFFSET 15", "got:", s)
	}

	if i[0].(string) != "Test <example@test.com>" {
		t.Error("expected first arg to be Test <example@test.com>, got:", i[0].(string))
	}

	if i[1].(int) != 1 {
		t.Error("expected second arg to be 1, got:", i[1].(int))
	}
}