package database

import (
	"context"
	"log"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/aldy505/bob"
)

// Set up the table connection, create table if not exists
func Setup() error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	db := New()

	// Jokesbapak2 table & data
	// Check if table exists
	sql, args, err := bob.HasTable("jokesbapak2").PlaceholderFormat(bob.Dollar).ToSQL()
	if err != nil {

		log.Fatalln("failed on checking database table:", err)
	}

	var hasTableJokes bool
	err = db.QueryRow(context.Background(), sql, args...).Scan(&hasTableJokes)
	if err != nil {
		if err.Error() == "no rows in result set" {
			hasTableJokes = false
		} else {
			log.Fatalln("failed on checking database table:", err)
		}
	}

	if !hasTableJokes {
		sql, _, err = bob.CreateTable("jokesbapak2").
			Columns("id", "link").
			Types("SERIAL", "VARCHAR(255)").
			Primary("id").ToSQL()
		if err != nil {
			log.Fatalln("failed on table creation:", err)
		}

		splitSql := strings.Split(sql, ";")
		for i := range splitSql {
			_, err = db.Query(context.Background(), splitSql[i])
			if err != nil {
				log.Println(sql)
				log.Fatalln("Failed on table creation: ", err)
				return err
			}
		}

		insertQuery, args, err := psql.Insert("jokesbapak2").
			Columns("link").
			Values("https://i.ibb.co/19pntdQ/Ea-p8-BWU8-AAtbjp.jpg").
			ToSql()
		if err != nil {
			log.Fatalln("Failed on query creation: ", err)
			return err
		}
		_, err = db.Query(context.Background(), insertQuery, args...)
		if err != nil {
			log.Fatalln("Failed on table insertion: ", err)
			return err
		}
	}

	// Authorization
	// Check if table exists
	sql, args, err = bob.HasTable("authorization").PlaceholderFormat(bob.Dollar).ToSQL()
	if err != nil {
		log.Fatalln("failed on checking database table:", err)
	}

	var hasTableAuth bool
	err = db.QueryRow(context.Background(), sql, args...).Scan(&hasTableAuth)
	if err != nil {
		if err.Error() == "no rows in result set" {
			hasTableAuth = false
		} else {
			log.Fatalln("failed on checking database table:", err)
		}
	}

	if !hasTableAuth {
		sql, _, err = bob.CreateTable("authorization").
			Columns("id", "token", "key").
			Types("SERIAL", "VARCHAR(255)", "VARCHAR(255)").
			Primary("id").
			Unique("token").
			ToSQL()
		if err != nil {
			log.Fatalln("Failed on table creation: ", err)
			return err
		}

		splitSql := strings.Split(sql, ";")
		for i := range splitSql {
			_, err = db.Query(context.Background(), splitSql[i])
			if err != nil {
				log.Fatalln("Failed on table creation: ", err)
				return err
			}
		}
	}

	return nil
}
