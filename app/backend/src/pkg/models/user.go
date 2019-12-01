package models

import (
	"strings"

	"github.com/go-playground/validator/v10"
	uuid "github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	// User defines a user
	User struct {
		BaseModel
		UUID           string   `json:"uuid" gorm:"not_null;unique"`
		Username       string   `json:"username" gorm:"not null" validate:"required"`
		Email          string   `json:"email" gorm:"not null;unique" validate:"required,email"`
		Password       string   `json:"password" gorm:"not null" validate:"required,gte=8,eqfield=RepeatPassword"`
		RepeatPassword string   `json:"repeat_password,omitempty" gorm:"-"`
		Settings       Settings `json:"settings" gorm:"foreignkey:UserID"`
	}

	// Settings holds the settings per user
	Settings struct {
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
	Credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

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
	if err := u.validate(); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	if err := db.Create(&u).Error; err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return ErrUnique
		}
		return err
	}

	return nil
}

// Update updates an existing user
func (u *User) Update(patch *User) (*User, error) {
	db.Model(&u).Updates(patch)

	return u, nil
}

func (u *User) validate() error {
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)
	u.Password = strings.TrimSpace(u.Password)

	if err := validate.Struct(u); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			r := strings.NewReplacer("{field}", v.Field(), "{param}", v.Param())
			e := r.Replace(errList[v.Tag()])

			// return first error found.
			return ValidationError{e}
		}
	}

	return nil
}
