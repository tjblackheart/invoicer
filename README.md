# Invoicer

A simple multiuser invoicing application using a RESTful backend written in Go and a vue.js frontend,
because I did not find any solution that fitted me. SQLITE is used as a database backend.

This is a work in progress, see [TODO](#todo) for things not working yet. Use at your own risk - although nothing should really break.

## Install

You'll need:

* a working Go-1.11 installation
* make
* yarn

Build:

```
git clone https://github.com/tjblackheart/invoicer
cd invoicer
cp .env.dist .env
```

If you want to use the .env file, edit it to your liking. At the bare minimum you'll have to add a signing secret.
If not, export all these variables into your ENV.

When done, type `make all` to build the binary, then run the app: `./invoicer`.
By default it is available at localhost:8000.

## <a name="todo"></a> TODO

* PDF creation
* A dockerized version
* Frontend bugs
* Cleanup

# License

[GPLv3](LICENSE)
