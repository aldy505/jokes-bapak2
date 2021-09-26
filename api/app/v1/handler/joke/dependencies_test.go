package joke_test

import (
	"context"
	v1 "jokes-bapak2-api/app/v1"
	"jokes-bapak2-api/app/v1/platform/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var jokesData = []interface{}{1, "https://via.placeholder.com/300/06f/fff.png", 1, 2, "https://via.placeholder.com/300/07f/fff.png", 1, 3, "https://via.placeholder.com/300/08f/fff.png", 1}
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

	j, err := db.Query(context.Background(), "DROP TABLE \"jokesbapak2\"")
	if err != nil {
		panic(err)
	}
	defer j.Close()

	a, err := db.Query(context.Background(), "DROP TABLE \"administrators\"")
	if err != nil {
		panic(err)
	}

	defer a.Close()
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

	err = database.Setup(db)
	if err != nil {
		return err
	}

	a, err := db.Query(context.Background(), "INSERT INTO \"administrators\" (\"id\", \"key\", \"token\", \"last_used\") VALUES	(1, 'test', '$argon2id$v=19$m=65536,t=16,p=4$3a08c79fbf2222467a623df9a9ebf75802c65a4f9be36eb1df2f5d2052d53cb7$ce434bd38f7ba1fc1f2eb773afb8a1f7f2dad49140803ac6cb9d7256ce9826fb3b4afa1e2488da511c852fc6c33a76d5657eba6298a8e49d617b9972645b7106', '');")
	if err != nil {
		return err
	}

	defer a.Close()

	j, err := db.Query(context.Background(), "INSERT INTO \"jokesbapak2\" (id, link, creator) VALUES ($1, $2, $3), ($4, $5, $6), ($7, $8, $9);", jokesData...)
	if err != nil {
		return err
	}

	defer j.Close()

	return nil
}
