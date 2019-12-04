# Invoicer

A multi user invoicing application using a RESTful backend written in Go using a vue.js frontend,
because I did not find any solution that fitted me. SQLITE is used as a database backend.

## Build

### Development

You'll need docker >= 18.06 and docker-compose. Copy `.env` to `.env.local` and set `APP_SECRET`. Run `docker-compose up`. This will start both images with hot reload enabled. By default the frontend is reachable at `http://localhost:8080`.

### Production

Copy `.env` to `.env.production`, set all vars and run `docker-compose -f docker-compose.prod.yml up`. Or export all the vars and build the images with `docker build` command.

## License

[GPLv3](LICENSE)
