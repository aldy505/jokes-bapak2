package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

func InMemory() *gocache.Cache {
	cache := gocache.New(6*time.Hour, 6*time.Hour)
	return cache
}
