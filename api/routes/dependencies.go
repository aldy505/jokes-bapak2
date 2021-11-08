package routes

import (
	"github.com/Masterminds/squirrel"
	"github.com/allegro/bigcache/v3"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Dependencies struct {
	DB     *pgxpool.Pool
	Redis  *redis.Client
	Memory *bigcache.BigCache
	HTTP   *httpclient.Client
	Query  squirrel.StatementBuilderType
	App    *fiber.App
}
