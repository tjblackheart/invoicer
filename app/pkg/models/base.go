package models

import (
	"errors"
	"time"
)

var (
	ErrUnique             = errors.New("models: unique key already exists")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrUserNotFound       = errors.New("models: no such user")
	ErrCustomerNotFound   = errors.New("models: no such customer")
	ErrInvoiceNotFound    = errors.New("models: no such invoice")
)

// BaseModel overrides gorm.Model
type BaseModel struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
