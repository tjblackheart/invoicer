package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"

	"github.com/tjblackheart/invoicer/pkg/models"
)

type application struct {
	config struct {
		host, dsn, secret string
		production        bool
	}
	info, err *log.Logger
	db        *gorm.DB
	models    struct {
		customers *models.Customer
		invoices  *models.Invoice
		users     *models.User
	}
}

var app application

func init() {
	app.info = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	app.err = log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	flag.StringVar(&app.config.dsn, "dsn", "./var/app.db", "SQLITE database path")
	flag.Parse()

	app.config.host = os.Getenv("APP_HOST")
	if app.config.host == "" {
		app.info.Println("APP_HOST not set: Using default value of :3000")
		app.config.host = ":3000"
	}

	app.config.secret = os.Getenv("APP_SECRET")
	if app.config.secret == "" {
		app.err.Fatalln("APP_SECRET not set: Aborting.")
	}

	app.config.production = false
	if os.Getenv("APP_ENV") == "prod" {
		app.config.production = true
	}
}

func main() {
	srv := &http.Server{
		Handler:      app.routes(),
		Addr:         app.config.host,
		ErrorLog:     app.err,
		IdleTimeout:  time.Minute,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	app.openDB()
	defer app.db.Close()

	app.models.users = &models.User{}
	app.models.users.Migrate(app.db)

	app.models.customers = &models.Customer{}
	app.models.customers.Migrate(app.db)

	app.models.invoices = &models.Invoice{}
	app.models.invoices.Migrate(app.db)

	app.info.Printf("Ready. Listening at %s, Production mode: %v\n", app.config.host, app.config.production)
	app.err.Fatal(srv.ListenAndServe())
}
