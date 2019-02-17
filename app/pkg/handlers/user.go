package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/tjblackheart/invoicer/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

// User defines a user
type User struct {
	db.Model
	UUID           string   `json:"uuid" gorm:"not_null;unique"`
	Email          string   `json:"email" gorm:"not null;unique"`
	Password       string   `json:"password,omitempty" gorm:"not null"`
	RepeatPassword string   `json:"repeat_password,omitempty" gorm:"-"`
	Username       string   `json:"username" gorm:"not null"`
	Settings       Settings `json:"settings" gorm:"foreignkey:UserID"`
}

// Settings holds the settings per user
type Settings struct {
	db.Model
	InvoiceNumberPrefix  string `json:"invoice_number_prefix" gorm:"not null"`
	CustomerNumberPrefix string `json:"customer_number_prefix" gorm:"not null"`
	NextInvoiceNumber    uint   `json:"next_invoice_number" gorm:"not null"`
	NextCustomerNumber   uint   `json:"next_customer_number" gorm:"not null"`
	TaxNumber            string `json:"tax_number"`
	Company              string `json:"company" gorm:"not null"`
	FirstName            string `json:"first_name" gorm:"not null"`
	LastName             string `json:"last_name" gorm:"not null"`
	Street               string `json:"street" gorm:"not null"`
	Number               string `json:"number" gorm:"not null"`
	Zip                  string `json:"zip" gorm:"not null"`
	City                 string `json:"city" gorm:"not null"`
	Country              string `json:"country" gorm:"not null"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	UserID               uint   `json:"user_id"`
}

// BeforeCreate generates an UUID for the user
func (u *User) BeforeCreate() error {
	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	u.UUID = uid.String()

	return nil
}

// UserMigrate initialises the tables
func UserMigrate() {
	db.Conn.AutoMigrate(&User{}, &Settings{})
}

// UserGet returns a single user
func UserGet(w http.ResponseWriter, r *http.Request) {
	var u User
	vars := mux.Vars(r)

	if db.Conn.Preload("Settings").Where("uuid = ?", vars["uuid"]).First(&u).RecordNotFound() {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	u.Password = ""

	json.NewEncoder(w).Encode(u)
}

// UserCreate creates a user
func UserCreate(w http.ResponseWriter, r *http.Request) {
	var u User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := u.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u.Password = string(hash)

	if err := db.Conn.Create(&u).Error; err != nil {
		e := ""
		if strings.Contains(err.Error(), "UNIQUE") && strings.Contains(err.Error(), "email") {
			e = "This email address is already in use"
		} else {
			e = err.Error()
		}

		http.Error(w, e, http.StatusBadRequest)
		return
	}

	// remove from response
	u.Password = ""
	u.RepeatPassword = ""

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// UserUpdate updates an existing user
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	var u User

	vars := mux.Vars(r)
	uuid := vars["uuid"]

	if db.Conn.Preload("Settings").Where("uuid = ?", uuid).First(&u).RecordNotFound() {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var patch User
	err := json.NewDecoder(r.Body).Decode(&patch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db.Conn.Model(&u).Updates(patch)

	json.NewEncoder(w).Encode(u)
}

// Validate validates a new user request
func (u *User) Validate() (valid bool, err error) {
	valid = false
	err = nil

	rxEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	rxUser := regexp.MustCompile("^[\\w\\s0-9_-]*$")

	if u.Password == "" {
		err = errors.New("Password cannot be empty")
		return
	}

	if len(u.Password) < 8 {
		err = errors.New("Password length should be more or equal to 8 characters")
		return
	}

	if u.Password != u.RepeatPassword {
		err = errors.New("Passwords do not match")
		return
	}

	if rxEmail.MatchString(u.Email) == false {
		err = errors.New("Please enter a valid email")
		return
	}

	if rxUser.MatchString(u.Username) == false {
		err = errors.New("Username cannot contain special chars")
		return
	}

	return
}
