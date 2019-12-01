package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	log "github.com/sirupsen/logrus"
)

func (app *Application) openDB() {
	db, err := gorm.Open("sqlite3", app.cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.DB().Ping(); err != nil {
		log.Fatal("Error opening database: ", err)
	}

	db.LogMode(app.cfg.Environment == "dev")
	app.db = db

	log.Info("SQLITE connection established: ", app.cfg.DB)
}
