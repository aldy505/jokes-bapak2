package database

import (
	"context"
	"log"

	"github.com/aldy505/bob"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Setup the table connection, create table if not exists
func Setup(db *pgxpool.Pool, ctx *context.Context) error {
	conn, err := db.Acquire(*ctx)
	if err != nil {
		log.Fatalln("30 - err here")
		return err
	}
	defer conn.Release()

	err = setupAuthTable(conn, ctx)
	if err != nil {
		return err
	}

	conn2, err := db.Acquire(*ctx)
	if err != nil {
		log.Fatalln("32 - err here")
		return err
	}
	defer conn2.Release()

	err = setupJokesTable(conn2, ctx)
	if err != nil {
		return err
	}

	conn3, err := db.Acquire(*ctx)
	if err != nil {
		return err
	}
	defer conn3.Release()

	err = setupSubmissionTable(conn3, ctx)
	if err != nil {
		return err
	}

	return nil
}

func setupAuthTable(conn *pgxpool.Conn, ctx *context.Context) error {
	// Check if table exists
	var tableAuthExists bool
	err := conn.QueryRow(*ctx, `SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE  table_schema = 'public'
		AND    table_name   = 'administrators'
		);`).Scan(&tableAuthExists)
	if err != nil {
		return err
	}

	if !tableAuthExists {
		sql, _, err := bob.
			CreateTable("administrators").
			AddColumn(bob.ColumnDef{Name: "id", Type: "SERIAL", Extras: []string{"PRIMARY KEY"}}).
			StringColumn("key", "NOT NULL", "UNIQUE").
			TextColumn("token").
			StringColumn("last_used").
			ToSql()
		if err != nil {
			return err
		}

		q, err := conn.Query(*ctx, sql)
		if err != nil {
			return err
		}
		defer q.Close()
	}
	return nil
}

func setupJokesTable(conn *pgxpool.Conn, ctx *context.Context) error {
	// Check if table exists
	var tableJokesExists bool
	err := conn.QueryRow(*ctx, `SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE  table_schema = 'public'
		AND    table_name   = 'jokesbapak2'
		);`).Scan(&tableJokesExists)
	if err != nil {
		return err
	}

	if !tableJokesExists {
		sql, _, err := bob.
			CreateTable("jokesbapak2").
			AddColumn(bob.ColumnDef{Name: "id", Type: "SERIAL", Extras: []string{"PRIMARY KEY"}}).
			TextColumn("link", "UNIQUE").
			AddColumn(bob.ColumnDef{Name: "creator", Type: "INT", Extras: []string{"NOT NULL", "REFERENCES \"administrators\" (\"id\")"}}).
			ToSql()
		if err != nil {
			return err
		}

		q, err := conn.Query(*ctx, sql)
		if err != nil {
			return err
		}
		defer q.Close()
	}

	return nil
}

func setupSubmissionTable(conn *pgxpool.Conn, ctx *context.Context) error {
	//Check if table exists
	var tableSubmissionExists bool
	err := conn.QueryRow(*ctx, `SELECT EXISTS (
	SELECT FROM information_schema.tables 
	WHERE  table_schema = 'public'
	AND    table_name   = 'submission'
	);`).Scan(&tableSubmissionExists)
	if err != nil {
		return err
	}

	if !tableSubmissionExists {
		sql, _, err := bob.
			CreateTable("submission").
			AddColumn(bob.ColumnDef{Name: "id", Type: "SERIAL", Extras: []string{"PRIMARY KEY"}}).
			TextColumn("link", "UNIQUE", "NOT NULL").
			StringColumn("created_at").
			StringColumn("author", "NOT NULL").
			AddColumn(bob.ColumnDef{Name: "status", Type: "SMALLINT", Extras: []string{"DEFAULT 0"}}).
			ToSql()
		if err != nil {
			return err
		}

		q, err := conn.Query(*ctx, sql)
		if err != nil {
			return err
		}
		defer q.Close()
	}
	return nil
}
