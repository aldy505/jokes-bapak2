package handler

import (
	"jokes-bapak2-api/app/v1/platform/cache"
	"jokes-bapak2-api/app/v1/platform/database"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/gojek/heimdall/v7/httpclient"
)

var Psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
var Db = database.New()
var Redis = cache.New()
var Memory = cache.InMemory()
var Client = httpclient.NewClient(httpclient.WithHTTPTimeout(10 * time.Second))
