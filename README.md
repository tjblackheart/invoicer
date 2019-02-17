# Invoicer

A simple multiuser invoicing application using a RESTful backend written in Go and a vue.js frontend,
because I did not find any solution that fitted me. SQLITE is used as a database backend.

This is a work in progress, see [TODO](#todo) for things not working yet. Use at your own risk - although nothing should really break.

## Install

Copy `env/prod.env.dist` to `env/prod.env` and add the missing `APP_SECRET`. Check out `docker-compose.yml`, edit it to your liking and run `docker-compose up`.

## Develop

Checkout the repo, `cd` into `app/` and run `cp .env.dist .env`. Edit this file to your liking - at least set the `APP_SECRET`, else the app won't run. When done, run `go run cmd/*` to start the dev backend with logging, and in another tab run `yarn --cwd=ui && yarn --cwd ui serve` to start the dev server.

By default the frontend is reachable at `http://localhost:8000`.

## <a name="todo"></a> TODO

* PDF creation
* --A dockerized version--
* Frontend bugs
* Cleanup

## License

[GPLv3](LICENSE)
