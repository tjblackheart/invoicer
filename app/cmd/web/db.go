package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func (app *application) openDB() {
	db, err := gorm.Open("sqlite3", app.config.dsn)
	if err != nil {
		app.err.Fatal(err)
	}

	db.LogMode(!app.config.production)

	app.db = db
}
