package submit_test

import (
	"context"
	v1 "jokes-bapak2-api/app/v1"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var submissionData = []interface{}{1, "https://via.placeholder.com/300/01f/fff.png", "2021-08-03T18:20:38Z", "Test <test@example.com>", 0, 2, "https://via.placeholder.com/300/02f/fff.png", "2021-08-04T18:20:38Z", "Test <test@example.com>", 1}
var app *fiber.App = v1.New()
var db *pgxpool.Pool

func cleanup() {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Unable to create pool config", err)
	}
	poolConfig.MaxConns = 5
	poolConfig.MinConns = 2

	db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalln("Unable to create connection", err)
	}

	s, err := db.Query(context.Background(), "DROP TABLE \"submission\"")
	if err != nil {
		panic(err)
	}
	defer s.Close()
}

func setup() error {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Unable to create pool config", err)
	}
	poolConfig.MaxConns = 15
	poolConfig.MinConns = 2

	db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalln("Unable to create connection", err)
	}

	s, err := db.Query(context.Background(), "INSERT INTO \"submission\" (id, link, created_at, author, status) VALUES ($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10);", submissionData...)
	if err != nil {
		return err
	}

	defer s.Close()

	return nil
}
