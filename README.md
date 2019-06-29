# Invoicer

A simple multiuser invoicing application using a RESTful backend written in Go and a vue.js frontend,
because I did not find any solution that fitted me. SQLITE is used as a database backend.

This is a work in progress, see [TODO](#todo) for things not working yet. Use at your own risk - although nothing should really break.

## Install

Copy `app/env.dist` to `app/.env` and add the missing `APP_SECRET`. Check out `docker-compose.yml`, edit it to your liking and run `docker-compose up`.

By default the frontend is reachable at `http://localhost:3000`.

## <a name="todo"></a> TODO

* PDF creation
* ~~A dockerized version~~
* Frontend bugs
* ~~Cleanup~~

## License

[GPLv3](LICENSE)
