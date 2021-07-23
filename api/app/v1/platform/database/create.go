package database

import (
	"context"
	"log"

	"github.com/aldy505/bob"
)

// Setup the table connection, create table if not exists
func Setup() error {
	db := New()

	// administrators table
	var tableAuthExists bool
	err := db.QueryRow(context.Background(), `SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE  table_schema = 'public'
		AND    table_name   = 'administrators'
		);`).Scan(&tableAuthExists)
	if err != nil {
		log.Fatalln("16 - failed on checking table: ", err)
		return err
	}

	if !tableAuthExists {
		sql, _, err := bob.
			CreateTable("administrators").
			AddColumn(bob.ColumnDef{Name: "id", Type: "SERIAL", Extras: []string{"PRIMARY KEY"}}).
			StringColumn("key", "UNIQUE").
			TextColumn("token").
			StringColumn("last_used").
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
	}

	// Jokesbapak2 table

	// Check if table exists
	var tableJokesExists bool
	err = db.QueryRow(context.Background(), `SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE  table_schema = 'public'
		AND    table_name   = 'jokesbapak2'
		);`).Scan(&tableJokesExists)
	if err != nil {
		log.Fatalln("10 - failed on checking table: ", err)
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
			log.Fatalln("11 - failed on table creation: ", err)
			return err
		}

		_, err = db.Query(context.Background(), sql)
		if err != nil {
			log.Fatalln("12 - failed on table creation: ", err)
			return err
		}
	}

	return nil
}
