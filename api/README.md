# Jokes Bapak2 API

Still work in progress

## Development

```bash
# Install modules
$ go mod download
# or
$ go mod vendor

# run the local server
$ go run main.go

# build everything
$ go build main.go
```

## Used packages

| Name | Version | Type |
| --- | --- | --- |
| [gofiber/fiber](https://github.com/gofiber/fiber) | `v2.14.0` | Framework |
| [jackc/pgx](https://github.com/jackc/pgx) | `v4.11.0` | Database |
| [go-redis/redis](https://github.com/go-redis/redis) | `v8.11.0` | Cache |
| [joho/godotenv](https://github.com/joho/godotenv) | `v1.3.0` | Config |
| [getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `v0.11.0` | Logging |
| [aldy505/phc-crypto](https://github.com/aldy505/phc-crypto) | `v1.1.0` | Utils |
| [Masterminds/squirrel](https://github.com/Masterminds/squirrel ) | `v1.5.0` | Utils |
| [aldy505/bob](https://github.com/aldy505/bob) | `v0.0.1` | Utils |

## Directory structure

```
└-- /app
    └---- /v1
          └---- /handler
          └---- /middleware             folder for add middleware 
          └---- /models
          └---- /platform
                └--------- /cache       folder with in-memory cache setup functions
                └--------- /database    folder with database setup functions 
          └---- /routes                 folder for describe routes
          └---- /utils                  folder with utility functions 
```
## `.env` configuration

```ini
ENV=development
PORT=5000

DATABASE_URL=postgres://postgres:password@localhost:5432/jokesbapak2
REDIS_URL=redis://@localhost:6379

SENTRY_DSN=
```