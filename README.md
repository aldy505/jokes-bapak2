# Jokes Bapak2 API

👋 Hey there! Still work in progress, if you'd like to contribute this while this repo is still growing, that would be so great!

## Brief explanation of what is this

This project will be a website like icanhazdadjokes but in Indonesian version and it's not text, it's images. Dad jokes in Indonesia is somewhat a bit different than in US/UK because I guess here, it's a lot dumber.

## Project Directories

* `api` - REST API service. Created with Go.
* `client` - Front facing website (front end). Created with [Svelte Kit](https://kit.svelte.dev/).

Anyway, later you can consume this API via a website (that will be created later on when this is finished) with a few endpoints:

 * `/v1/` - Random jokes bapak2
 * `/v1/id/[number]` - Jokes bapak2 based on ID
 * `/v1/today` - Jokes bapak2 of the day

Currently I'm searching for an alternative for AWS S3 that I can use for free.

## Tech stacks

 * Go (for `api`/back end)
 * Node.js (for `client`/front end)
 * Postgres
 * Redis

That's it.

## Development

Two ways of doing this:
  1. Install all the tech stack on your local machine
  2. Use docker-compose (TODO)

See README files on each project directory for further instruction on how to run the development environment.

## License

Jokes Bapak2 API is licensed under [GNU GENERAL PUBLIC LICENSE v3 license](./LICENSE)