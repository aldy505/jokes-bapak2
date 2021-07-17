package handler

import (
	"jokes-bapak2-api/app/v1/platform/cache"
	"jokes-bapak2-api/app/v1/platform/database"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/gojek/heimdall/v7/httpclient"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
var db = database.New()
var redis = cache.New()
var memory = cache.InMemory()
var client = httpclient.NewClient(httpclient.WithHTTPTimeout(10 * time.Second))
