package models

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/nu7hatch/gouuid"
)

// User defines a user
type User struct {
	BaseModel
	UUID           string   `json:"uuid" gorm:"not_null;unique"`
	Email          string   `json:"email" gorm:"not null;unique"`
	Password       string   `json:"password,omitempty" gorm:"not null"`
	RepeatPassword string   `json:"repeat_password,omitempty" gorm:"-"`
	Username       string   `json:"username" gorm:"not null"`
	Settings       Settings `json:"settings" gorm:"foreignkey:UserID"`
}

// Settings holds the settings per user
type Settings struct {
	BaseModel
	InvoiceNumberPrefix  string `json:"invoice_number_prefix" gorm:"not null"`
	CustomerNumberPrefix string `json:"customer_number_prefix" gorm:"not null"`
	NextInvoiceNumber    uint   `json:"next_invoice_number" gorm:"not null"`
	NextCustomerNumber   uint   `json:"next_customer_number" gorm:"not null"`
	TaxNumber            string `json:"tax_number" gorm:"not null"`
	Company              string `json:"company" gorm:"not null"`
	FirstName            string `json:"first_name" gorm:"not null"`
	LastName             string `json:"last_name" gorm:"not null"`
	Street               string `json:"street"`
	Number               string `json:"number"`
	Zip                  string `json:"zip"`
	City                 string `json:"city"`
	Country              string `json:"country"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	Bank                 string `json:"bank"`
	IBAN                 string `json:"iban"`
	BIC                  string `json:"bic"`
	UserID               uint   `json:"user_id"`
}

// Credentials holds login form data
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// FindUser finds a user by UUID
func FindUser(uuid string) (*User, error) {
	var u User

	if db.Preload("Settings").Where("uuid = ?", uuid).First(&u).RecordNotFound() {
		return nil, ErrUserNotFound
	}

	return &u, nil
}

// Authenticate performs a login attempt
func Authenticate(c *Credentials) (*User, error) {
	user := &User{}

	if db.Preload("Settings").Where("email = ?", c.Email).First(&user).RecordNotFound() {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
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

// Create creates a user
func (u *User) Create() error {
	if _, err := u.validate(); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}
	u.Password = string(hash)

	if err := db.Create(&u).Error; err != nil {
		return err
	}

	return nil
}

// Update updates an existing user
func (u *User) Update(patch *User) (*User, error) {
	db.Model(&u).Updates(patch)

	return u, nil
}

func (u *User) validate() (valid bool, err error) {
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
