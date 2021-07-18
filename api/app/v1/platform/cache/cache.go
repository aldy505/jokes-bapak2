package cache

import (
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
)

func InMemory() *bigcache.BigCache {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(6 * time.Hour))
	if err != nil {
		log.Fatalln(err)
	}
	return cache
}
