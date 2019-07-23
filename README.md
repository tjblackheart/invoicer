# Invoicer

A simple multi user invoicing application using a RESTful backend written in Go using a vue.js frontend,
because I did not find any solution that fitted me. SQLITE is used as a database backend.

## Usage

### With Docker

Copy `app/env.dist` to `app/.env` and add the missing `APP_SECRET`. Check `docker-compose.yml`, edit it to your liking and run `docker-compose up`. By default the frontend is reachable at `http://localhost:3000`.

### Without Docker / Development version

* Copy `app/env.dist` to `app/.env` and add the missing `APP_SECRET`
* `cd` into `app/ui/` and run `yarn && yarn serve`
* `cd` into `app` and run `go run cmd/web/*`

By default the development frontend is reachable at `http://localhost:8080`.

## License

[GPLv3](LICENSE)
