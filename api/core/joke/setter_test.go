package joke_test

import (
	"context"
	"jokes-bapak2-api/core/joke"
	"jokes-bapak2-api/core/schema"
	"testing"

	"github.com/jackc/pgx/v4"
)

func TestSetAllJSONJoke(t *testing.T) {
	t.Cleanup(func() { Flush() })

	conn, err := db.Acquire(context.Background())
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer conn.Release()

	err = conn.BeginFunc(context.Background(), func(t pgx.Tx) error {
		_, err := t.Exec(
			context.Background(),
			`INSERT INTO "administrators"
				(id, key, token, last_used)
				VALUES
				($1, $2, $3, $4),
				($5, $6, $7, $8);`,
			administratorsData...,
		)
		if err != nil {
			return err
		}
		_, err = t.Exec(
			context.Background(),
			`INSERT INTO "jokesbapak2" 
				(id, link, creator)
				VALUES
				($1, $2, $3),
				($4, $5, $6),
				($7, $8, $9);`,
			jokesData...,
		)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	err = joke.SetAllJSONJoke(db, context.Background(), memory)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
}

func TestSetTotalJoke(t *testing.T) {
	t.Cleanup(func() { Flush() })

	conn, err := db.Acquire(context.Background())
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer conn.Release()

	err = conn.BeginFunc(context.Background(), func(t pgx.Tx) error {
		_, err := t.Exec(
			context.Background(),
			`INSERT INTO "administrators"
				(id, key, token, last_used)
				VALUES
				($1, $2, $3, $4),
				($5, $6, $7, $8);`,
			administratorsData...,
		)
		if err != nil {
			return err
		}
		_, err = t.Exec(
			context.Background(),
			`INSERT INTO "jokesbapak2" 
				(id, link, creator)
				VALUES
				($1, $2, $3),
				($4, $5, $6),
				($7, $8, $9);`,
			jokesData...,
		)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	err = joke.SetTotalJoke(db, context.Background(), memory)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
}

func TestInsertJokeIntoDB(t *testing.T) {
	t.Cleanup(func() { Flush() })

	data := schema.Joke{
		ID:      1,
		Link:    "link1",
		Creator: 1,
	}
	err := joke.InsertJokeIntoDB(db, context.Background(), data)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
}

func TestDeleteSingleJoke(t *testing.T) {
	t.Cleanup(func() { Flush() })

	conn, err := db.Acquire(context.Background())
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer conn.Release()

	err = conn.BeginFunc(context.Background(), func(t pgx.Tx) error {
		_, err := t.Exec(
			context.Background(),
			`INSERT INTO "administrators"
				(id, key, token, last_used)
				VALUES
				($1, $2, $3, $4),
				($5, $6, $7, $8);`,
			administratorsData...,
		)
		if err != nil {
			return err
		}
		_, err = t.Exec(
			context.Background(),
			`INSERT INTO "jokesbapak2" 
				(id, link, creator)
				VALUES
				($1, $2, $3),
				($4, $5, $6),
				($7, $8, $9);`,
			jokesData...,
		)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	err = joke.DeleteSingleJoke(db, context.Background(), 1)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
}

func TestUpdateJoke(t *testing.T) {
	t.Cleanup(func() { Flush() })

	conn, err := db.Acquire(context.Background())
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer conn.Release()

	err = conn.BeginFunc(context.Background(), func(t pgx.Tx) error {
		_, err := t.Exec(
			context.Background(),
			`INSERT INTO "administrators"
				(id, key, token, last_used)
				VALUES
				($1, $2, $3, $4),
				($5, $6, $7, $8);`,
			administratorsData...,
		)
		if err != nil {
			return err
		}
		_, err = t.Exec(
			context.Background(),
			`INSERT INTO "jokesbapak2" 
				(id, link, creator)
				VALUES
				($1, $2, $3),
				($4, $5, $6),
				($7, $8, $9);`,
			jokesData...,
		)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	newJoke := schema.Joke{
		ID:      1,
		Link:    "link1",
		Creator: 1,
	}

	err = joke.UpdateJoke(db, context.Background(), newJoke)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
}
