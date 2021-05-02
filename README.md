# Jokes Bapak2 API

⚠ Still work in progress

## Brief explanation of what is this

This project will be a website like icanhazdadjokes but in Indonesian version and it's not text, it's images. Dad jokes in Indonesia is somewhat a bit different than in US/UK because I guess here, it's a lot dumber.

## Project Directories

* `api` - REST API service. Created with Go.
* `client` - Front facing website (front end). Created with Vite.js

Anyway, later you can consume this API via a website (that will be created later on when this is finished) with a few endpoints:

 * `/v1/` - Random jokes bapak2
 * `/v1/id/[number]` - Jokes bapak2 based on ID
 * `/v1/today` - Jokes bapak2 of the day

Currently I'm searching for an alternative for AWS S3 that I can use for free.

## Tech stacks

 * Go
 * 

That's it.

## License

Copyright 2021-present Jokes Bapak2 API Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.