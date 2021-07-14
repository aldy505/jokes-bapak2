package database

import (
	"context"
	"log"

	"github.com/aldy505/bob"
)

// Set up the table connection, create table if not exists
func Setup() error {
	db := New()

	// Jokesbapak2 table

	// Check if table exists
	var tableJokesExists bool
	err := db.QueryRow(context.Background(), `SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE  table_schema = 'public'
		AND    table_name   = 'jokesbapak2'
		);`).Scan(&tableJokesExists)
	if err != nil {
		log.Fatalln("10 - failed on checking table: ", err)
		return err
	}

	if !tableJokesExists {
		sql, _, err := bob.CreateTable("jokesbapak2").
			Columns("id", "link").
			Types("SERIAL", "TEXT").
			ToSql()
		if err != nil {
			log.Fatalln("11 - failed on table creation: ", err)
			return err
		}

		_, err = db.Query(context.Background(), sql)
		if err != nil {
			log.Fatalln("12 - failed on table creation: ", err)
			return err
		}
		_, err = db.Query(context.Background(), "ALTER TABLE \"jokesbapak2\" ADD creator INT NOT NULL DEFAULT 0;")
		if err != nil {
			log.Fatalln("13 - failed on table creation: ", err)
			return err
		}

		_, err = db.Query(context.Background(), "ALTER TABLE \"jokesbapak2\" ADD PRIMARY KEY (\"id\")")
		if err != nil {
			log.Fatalln("14 - failed on table alteration: ", err)
			return err
		}

		_, err = db.Query(context.Background(), "ALTER TABLE \"jokesbapak2\" ADD UNIQUE (\"link\")")
		if err != nil {
			log.Fatalln("15 - failed on table alteration: ", err)
			return err
		}
	}

	// administrators table
	var tableAuthExists bool
	err = db.QueryRow(context.Background(), `SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE  table_schema = 'public'
		AND    table_name   = 'administrators'
		);`).Scan(&tableAuthExists)
	if err != nil {
		log.Fatalln("16 - failed on checking table: ", err)
		return err
	}

	if !tableAuthExists {
		sql, _, err := bob.CreateTable("administrators").
			Columns("id", "key", "token", "last_used").
			Types("SERIAL", "VARCHAR(255)", "TEXT", "VARCHAR(255)").
			ToSql()
		if err != nil {
			log.Fatalln("17 - failed on table creation: ", err)
			return err
		}

		_, err = db.Query(context.Background(), sql)
		if err != nil {
			log.Fatalln("18 - failed on table creation: ", err)
			return err
		}

		_, err = db.Query(context.Background(), "ALTER TABLE \"administrators\" ADD PRIMARY KEY (\"id\");")
		if err != nil {
			log.Fatalln("19 - failed on table alteration: ", err)
			return err
		}

		_, err = db.Query(context.Background(), "ALTER TABLE \"administrators\" ADD UNIQUE (\"key\");")
		if err != nil {
			log.Fatalln("20 - failed on table alteration: ", err)
			return err
		}

		_, err = db.Query(context.Background(), "ALTER TABLE \"jokesbapak2\" ADD CONSTRAINT fk_auth_key FOREIGN KEY (\"creator\") REFERENCES \"administrators\" (\"id\");")
		if err != nil {
			log.Fatalln("21 - failed on foreign key iteration: ", err)
			return err
		}
	}
	return nil
}
