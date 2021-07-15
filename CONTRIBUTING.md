# Contributing Guide

First of all. Thank you for considering to contribute on Jokes Bapak2 API project. I hope this project will get better and we will become more bapak2 than ever.

This project is a monorepo, meaning the backend, frontend, and Github CI are all in one place (one repository). Before you do anything, if you're going to do some breaking change or you'll write (or remove) large numbers of LOC (line of codes), please open an issue first and let us know about it. So that our work won't bother you and you'll have a breeze on developing this.

## Project Prerequisites && Setup

### Front End (`./client`)

You'll have to install:
* Node.js LTS (preferably with [fnm](https://github.com/Schniz/fnm) or [nvm](https://github.com/nvm-sh/nvm))
* Yarn v1

See the [README](./client/README.md) on client for detailed project setup.

### Back End (`./api`)

You'll have to install:
* Go v1.16.x
* (Optional) [Fiber CLI](https://github.com/gofiber/cli) for ease of development

See the [README](./api/README.md) on client for detailed project setup.

### With Docker Compose

If you're just developing the front end and too lazy installing Go and such (or the other way around), you can use `docker-compose` file specified on the main page.

You'll have to install:
* Docker (preferably with Docker Desktop if you're on Windows or Mac)
* Docker Compose

```bash
# Create a docker container but don't start it yet.
$ docker-compose up --no-start

# Or if you want to create the docker container and start it right away
$ docker-compose up

# If you want to have it running in the background
$ docker-compose up --detach

# Start existing container
$ docker-compose start

# Stop running container
$ docker-compose stop

# Destroy current container
$ docker-compose down
```

## Before submitting PR

### Front End (`./client`)

Please run these:
* `yarn lint`
* `yarn format`
* `yarn build`

If those command didn't pass, please fix the problem first. Please recheck your changes, make sure NOT to leave any secret token/keys behind.

### Back End (`./api`)

Please run these:
* `go fmt`
* `go build main.go`
* `go test -v -race -coverprofile=coverage.out -covermode=atomic ./...`

If those command didn't pass, please fix the problem first. Please recheck your changes, make sure NOT to leave any secret token/keys behind.

## One more thing..

Oh my God, thank you so much!!! Working on an open source project is interesting right?? 😆