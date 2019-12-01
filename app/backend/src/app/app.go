package app

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tjblackheart/Invoicer/backend/pkg/models"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type (
	Config struct {
		Hostname, Environment string
		DB                    string
		Secret                string
	}

	Application struct {
		cfg *Config
		db  *gorm.DB
	}
)

func New(cfg *Config) *Application {
	app := &Application{cfg: cfg}

	app.openDB()
	models.Init(app.db)

	return app
}

func (app Application) Serve() {
	srv := &http.Server{
		Handler:      app.routes(),
		Addr:         app.cfg.Hostname,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	defer app.db.Close()

	log.Infof("API listening at %s, Environment: %v", app.cfg.Hostname, app.cfg.Environment)
	log.Fatal(srv.ListenAndServe())
}
