package joke

import (
	"github.com/allegro/bigcache/v3"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
)

// Dependencies provides a struct for dependency injection
// on joke package
type Dependencies struct {
	Redis  *redis.Client
	Memory *bigcache.BigCache
	Bucket *minio.Client
}
