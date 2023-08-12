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

There is a placeholder data ready for you to query it manually in `/platform/database/placeholder.sql`. Have a good time
developing!

## Used packages

| Name                                                             | Version                              | Type      |
|------------------------------------------------------------------|--------------------------------------|-----------|
| [gofiber/fiber](https://github.com/gofiber/fiber)                | `v2.21.0`                            | Framework |
| [jackc/pgx](https://github.com/jackc/pgx)                        | `v4.13.0`                            | Database  |
| [go-redis/redis](https://github.com/go-redis/redis)              | `v8.11.4`                            | Cache     |
| [allegro/bigcache](https://github.com/allegro/bigcache)          | `v3.0.1`                             | Cache     |
| [joho/godotenv](https://github.com/joho/godotenv)                | `v1.4.0`                             | Config    |
| [getsentry/sentry-go](https://github.com/getsentry/sentry-go)    | `v0.11.0`                            | Logging   |
| [aldy505/phc-crypto](https://github.com/aldy505/phc-crypto)      | `v1.1.0`                             | Utils     |
| [Masterminds/squirrel](https://github.com/Masterminds/squirrel ) | `v1.5.1`                             | Utils     |
| [aldy505/bob](https://github.com/aldy505/bob)                    | `v0.0.4`                             | Utils     |
| [gojek/heimdall](https://github.com/gojek/heimdall)              | `v7.0.2`                             | Utils     |
| [georgysavva/scany](https://github.com/georgysavva/scany)        | `v0.2.9`                             | Utils     |
| [pquerna/ffjson](https://github.com/pquerna/ffjson)              | `v0.0.0-20190930134022-aa0246cd15f7` | Utils     |

## Directory structure

```
.
├── core                  - Pure business logic
│  ├── administrator
│  ├── joke
│  ├── schema
│  ├── submit
│  └── validator
├── Dockerfile            - Docker image for API
├── documentation.json    - Swagger documentation
├── documentation.yaml    - Swagger documentation
├── favicon.png
├── go.mod                - Module declaration
├── go.sum                - Checksum for modules
├── handler               - Route handlers
│  ├── health
│  ├── joke
│  └── submit
├── main.go               - Application entry point
├── middleware            - Route middlewares
├── platform              - Third party packages
│  └── database
├── README.md             - You are here
├── routes                - Route definitions
└── utils                 - Utility functions
```

## `.env` configuration

```ini
ENV=development
PORT=5000

DATABASE_URL=postgres://postgres:password@localhost:5432/jokesbapak2
REDIS_URL=redis://@localhost:6379

SENTRY_DSN=
```