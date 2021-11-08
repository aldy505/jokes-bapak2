package joke_test

import (
	"context"
	"jokes-bapak2-api/core/joke"
	"jokes-bapak2-api/core/schema"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
)

func TestSetAllJSONJoke(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	defer Flush()

	conn, err := db.Acquire(ctx)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer conn.Release()

	err = conn.BeginFunc(ctx, func(t pgx.Tx) error {
		_, err := t.Exec(
			ctx,
			`INSERT INTO "administrators"
				(id, key, token, last_used)
				VALUES
				($1, $2, $3, $4),
				($5, $6, $7, $8)`,
			administratorsData...,
		)
		if err != nil {
			return err
		}
		_, err = t.Exec(
			ctx,
			`INSERT INTO "jokesbapak2" 
				(id, link, creator)
				VALUES
				($1, $2, $3),
				($4, $5, $6),
				($7, $8, $9)`,
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

	err = joke.SetAllJSONJoke(db, ctx, memory)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
}

func TestSetTotalJoke(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()
	
	defer Flush()

	conn, err := db.Acquire(ctx)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer conn.Release()

	err = conn.BeginFunc(ctx, func(t pgx.Tx) error {
		_, err := t.Exec(
			ctx,
			`INSERT INTO "administrators"
				(id, key, token, last_used)
				VALUES
				($1, $2, $3, $4),
				($5, $6, $7, $8)`,
			administratorsData...,
		)
		if err != nil {
			return err
		}
		_, err = t.Exec(
			ctx,
			`INSERT INTO "jokesbapak2" 
				(id, link, creator)
				VALUES
				($1, $2, $3),
				($4, $5, $6),
				($7, $8, $9)`,
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

	err = joke.SetTotalJoke(db, ctx, memory)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
}

func TestInsertJokeIntoDB(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()
	
	defer Flush()

	data := schema.Joke{
		ID:      1,
		Link:    "link1",
		Creator: 1,
	}
	err := joke.InsertJokeIntoDB(db, ctx, data)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
}

func TestDeleteSingleJoke(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()
	
	defer Flush()

	conn, err := db.Acquire(ctx)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer conn.Release()

	err = conn.BeginFunc(ctx, func(t pgx.Tx) error {
		_, err := t.Exec(
			ctx,
			`INSERT INTO "administrators"
				(id, key, token, last_used)
				VALUES
				($1, $2, $3, $4),
				($5, $6, $7, $8)`,
			administratorsData...,
		)
		if err != nil {
			return err
		}
		_, err = t.Exec(
			ctx,
			`INSERT INTO "jokesbapak2" 
				(id, link, creator)
				VALUES
				($1, $2, $3),
				($4, $5, $6),
				($7, $8, $9)`,
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

	err = joke.DeleteSingleJoke(db, ctx, 1)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
}

func TestUpdateJoke(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()
	
	defer Flush()

	conn, err := db.Acquire(ctx)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	defer conn.Release()

	err = conn.BeginFunc(ctx, func(t pgx.Tx) error {
		_, err := t.Exec(
			ctx,
			`INSERT INTO "administrators"
				(id, key, token, last_used)
				VALUES
				($1, $2, $3, $4),
				($5, $6, $7, $8)`,
			administratorsData...,
		)
		if err != nil {
			return err
		}
		_, err = t.Exec(
			ctx,
			`INSERT INTO "jokesbapak2" 
				(id, link, creator)
				VALUES
				($1, $2, $3),
				($4, $5, $6),
				($7, $8, $9)`,
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

	err = joke.UpdateJoke(db, ctx, newJoke)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
}
