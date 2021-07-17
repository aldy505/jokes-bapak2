package v1

import (
	"jokes-bapak2-api/app/v1/core"
	"jokes-bapak2-api/app/v1/platform/cache"
	"jokes-bapak2-api/app/v1/platform/database"
	"jokes-bapak2-api/app/v1/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	gocache "github.com/patrickmn/go-cache"
)

var memory = cache.InMemory()
var db = database.New()

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableKeepalive: true,
		CaseSensitive:    true,
	})

	checkCache := core.CheckJokesCache(memory)

	if !checkCache {
		jokes, err := core.GetAllJSONJokes(db)
		if err != nil {
			log.Fatalln(err)
		}
		memory.Set("jokes", jokes, gocache.NoExpiration)
		if err != nil {
			log.Fatalln(err)
		}
	}

	routes.Health(app)
	routes.Joke(app)

	return app
}
