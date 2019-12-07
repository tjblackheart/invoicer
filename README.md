# Invoicer

A multi user invoicing application using a RESTful backend written in Go using a vue.js frontend,
because I did not find any solution that fitted me. SQLITE is used as a database backend.

## Build

The easy way: `./start [dev|prod] [args]`

### Development

You'll need docker >= 18.06 and docker-compose. Copy `.env` to `.env.local` and set `APP_SECRET`. Run `docker-compose -f docker-compose.dev.yml up`. This will start both backend and frontend with a hot reload mechanism enabled. By default the frontend is reachable at `http://localhost:8080`.

Run the tests in a running container: `docker exec -t invoicer_backend_dev go test ./...`

### Production

Copy `.env` to `.env.production`, set all vars and run `docker-compose -f docker-compose.prod.yml up`. Or export all the vars and build the images with `docker build` command.

## License

[GPLv3](LICENSE)
