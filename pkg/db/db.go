package db

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Conn is the db connection
var Conn *gorm.DB
var err error

// Model overrides gorm.Model
type Model struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

// InitDB initialises the db connection
func InitDB(file string, mode bool) {
	Conn, err = gorm.Open("sqlite3", file)

	if err != nil {
		log.Fatalln(err)
	}

	Conn.LogMode(mode)
}
