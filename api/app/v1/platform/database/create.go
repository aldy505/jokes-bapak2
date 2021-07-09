package database

import (
	"context"
	"log"
	"strings"

	"github.com/aldy505/bob"
)

// Set up the table connection, create table if not exists
func Setup() error {
	db := New()

	// Jokesbapak2 table & data
	sql, _, err := bob.CreateTableIfNotExists("jokesbapak2").
		Columns("id", "link", "key").
		Types("SERIAL", "TEXT", "VARCHAR(255)").
		Primary("id").ToSql()
	if err != nil {
		log.Fatalln("10 - failed on table creation: ", err)
	}

	splitSql := strings.Split(sql, ";")
	for i := range splitSql {
		_, err = db.Query(context.Background(), splitSql[i])
		if err != nil {
			log.Fatalln("11 - failed on table creation: ", err)
			return err
		}
	}

	// Authorization
	sql, _, err = bob.CreateTableIfNotExists("authorization").
		Columns("id", "token", "key").
		Types("SERIAL", "TEXT", "VARCHAR(255)").
		Primary("id").
		Unique("token").
		ToSql()
	if err != nil {
		log.Fatalln("14 - failed on table creation: ", err)
		return err
	}

	splitSql = strings.Split(sql, ";")
	for i := range splitSql {
		_, err = db.Query(context.Background(), splitSql[i])
		if err != nil {
			log.Fatalln("15 - failed on table creation: ", err)
			return err
		}
	}

	_, err = db.Query(context.Background(), "ALTER TABLE jokesbapak2 ADD CONSTRAINT fk_jokesbapak2_key FOREIGN KEY (key) REFERENCES authorization (id)")
	if err != nil {
		log.Fatalln("16 - failed on foreign key iteration: ", err)
	}
	return nil
}
