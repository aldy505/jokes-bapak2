package health

import (
	"context"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
)

type Dependencies struct {
	Bucket *minio.Client
	Cache  *redis.Client
}

func (d *Dependencies) Health(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*15)
	defer cancel()

	var bucketOk = true
	var cacheOk = true

	cancel, err := d.Bucket.HealthCheck(time.Second * 15)
	if err != nil {
		bucketOk = false
	}

	if cancel != nil {
		cancel()
	}

	_, err = d.Cache.Ping(ctx).Result()
	if err != nil {
		cacheOk = false
	}

	if !bucketOk || !cacheOk {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}
