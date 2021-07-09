package cache

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

// Connect to the database
func New() *redis.Client {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	rdb := redis.NewClient(opt)
	return rdb
}
