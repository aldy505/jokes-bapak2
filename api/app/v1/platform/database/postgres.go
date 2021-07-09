package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Connect to the database
func New() *pgxpool.Pool {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Unable to create pool config", err)
	}
	poolConfig.MaxConns = 18
	poolConfig.MinConns = 2

	conn, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalln("Unable to create connection", err)
	}

	return conn
}
