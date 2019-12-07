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
		Username       trimmed  `json:"username" gorm:"not null" validate:"required"`
		Email          trimmed  `json:"email" gorm:"not null;unique" validate:"required,email"`
		Password       trimmed  `json:"password" gorm:"not null" validate:"required,gte=8,eqfield=RepeatPassword"`
		RepeatPassword trimmed  `json:"repeat_password,omitempty" gorm:"-"`
		Settings       Settings `json:"settings" gorm:"foreignkey:UserID"`
	}

	// Settings holds the settings per user
	Settings struct {
		BaseModel
		InvoiceNumberPrefix  trimmed `json:"invoice_number_prefix" gorm:"not null"`
		CustomerNumberPrefix trimmed `json:"customer_number_prefix" gorm:"not null"`
		NextInvoiceNumber    uint    `json:"next_invoice_number" gorm:"not null"`
		NextCustomerNumber   uint    `json:"next_customer_number" gorm:"not null"`
		TaxNumber            trimmed `json:"tax_number" gorm:"not null"`
		Company              trimmed `json:"company" gorm:"not null"`
		FirstName            trimmed `json:"first_name" gorm:"not null"`
		LastName             trimmed `json:"last_name" gorm:"not null"`
		Street               trimmed `json:"street"`
		Number               trimmed `json:"number"`
		Zip                  trimmed `json:"zip"`
		City                 trimmed `json:"city"`
		Country              trimmed `json:"country"`
		Email                trimmed `json:"email"`
		Phone                trimmed `json:"phone"`
		Bank                 trimmed `json:"bank"`
		IBAN                 trimmed `json:"iban"`
		BIC                  trimmed `json:"bic"`
		UserID               uint    `json:"user_id"`
	}

	// Credentials holds login form data
	Credentials struct {
		Email    trimmed `json:"email"`
		Password trimmed `json:"password"`
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

	u.Password = trimmed(string(hash))

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
