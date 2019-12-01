package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/tjblackheart/Invoicer/backend/app"
)

func main() {
	cfg := setup()
	app := app.New(cfg)
	app.Serve()
}

func setup() *app.Config {
	secret := os.Getenv("APP_SECRET")
	if secret == "" {
		log.Fatalln("APP_SECRET not set: Aborting.")
	}

	db := os.Getenv("APP_DB")
	if db == "" {
		db = "./var/app.db"
		log.Infof("APP_DB not set: Using default value of %s", db)
	}

	hostname := os.Getenv("APP_HOST")
	if hostname == "" {
		hostname = ":3000"
		log.Infof("APP_HOST not set: Using default value of %s", hostname)
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	return &app.Config{
		Hostname:    hostname,
		Environment: env,
		Secret:      secret,
		DB:          db,
	}
}
