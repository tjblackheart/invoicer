package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

var (
	ErrUnique             = errors.New("models: unique key already exists")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrUserNotFound       = errors.New("models: no such user")
	ErrCustomerNotFound   = errors.New("models: no such customer")
	ErrInvoiceNotFound    = errors.New("models: no such invoice")

	validate = validator.New()
	db       *gorm.DB
)

// BaseModel overrides gorm.Model
type (
	BaseModel struct {
		ID        uint       `json:"id" gorm:"primary_key"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"-"`
	}

	ValidationError struct {
		Message string
	}

	translatableErrors map[string]string
	trimmed            string
)

var errList = translatableErrors{
	"required": "Please enter a value for '{field}'.",
	"email":    "Please enter a valid email.",
	"eqfield":  "{field} fields should match.",
	"gte":      "{field} should contain {param} characters or more.",
	"gt":       "{field} can not be empty.",
}

func (t *trimmed) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*t = trimmed(strings.TrimSpace(s))

	return nil
}

func (e ValidationError) Error() string {
	return fmt.Sprintf(e.Message)
}

// Init initialises the db connection and migrates the tables.
func Init(connection *gorm.DB) {
	db = connection

	db.AutoMigrate(
		&User{},
		&Settings{},
		&Invoice{},
		&InvoiceItem{},
		&Customer{},
		&Address{},
		&Contact{},
	)
}
