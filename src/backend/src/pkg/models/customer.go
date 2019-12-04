package models

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// Customer represents a customer
type (
	Customer struct {
		BaseModel
		Number    trimmed   `json:"number" gorm:"unique;not null" validate:"required"`
		Remarks   string    `json:"remarks"`
		TaxNumber string    `json:"tax_number"`
		Address   Address   `json:"address" gorm:"foreignkey:CustomerID" validate:"dive"`
		Contacts  []Contact `json:"contacts" gorm:"foreignkey:CustomerID" validate:"dive"`
		UUID      string    `json:"-"`
	}

	// Address holds an address of a customer
	Address struct {
		BaseModel
		Company    trimmed `json:"company" gorm:"not null" validate:"required"`
		FirstName  string  `json:"first_name" gorm:"not null"`
		LastName   string  `json:"last_name" gorm:"not null"`
		Street     trimmed `json:"street" gorm:"not null" validate:"required"`
		Number     trimmed `json:"number" gorm:"not null" validate:"required"`
		Zip        trimmed `json:"zip" gorm:"not null" validate:"required"`
		City       trimmed `json:"city" gorm:"not null" validate:"required"`
		Country    string  `json:"country" gorm:"not null"`
		CustomerID uint    `json:"-"`
	}

	// Contact holds a contact value (email, phone ...)
	Contact struct {
		BaseModel
		Value      trimmed `json:"value" gorm:"not null" validate:"required"`
		Type       trimmed `json:"type" gorm:"not null" validate:"required"`
		CustomerID uint    `json:"-"`
	}
)

// CustomerGetAll returns a slice of all available customers
func CustomerGetAll(uuid string) (*[]Customer, error) {
	var customers []Customer

	err := db.
		Preload("Address").
		Preload("Contacts").
		Where("uuid = ?", uuid).
		Order("created_at DESC").
		Find(&customers).
		Error

	if err != nil {
		return nil, err
	}

	return &customers, nil
}

// CustomerGet fetches a single customer
func CustomerGet(uuid string, id string) (*Customer, error) {
	var customer Customer

	if db.
		Preload("Address").
		Preload("Contacts").
		Where("uuid = ? AND id = ?", uuid, id).
		First(&customer).
		RecordNotFound() {
		return nil, ErrCustomerNotFound
	}

	return &customer, nil
}

// CustomerCreate creates a new customer
func CustomerCreate(uuid string, c *Customer) (*Customer, error) {
	u, err := FindUser(uuid)
	if err != nil {
		return nil, err
	}

	c.UUID = uuid

	if err := c.validate(); err != nil {
		return nil, err
	}

	if err := db.Create(&c).Error; err != nil {
		if strings.Contains(err.Error(), "UNIQUE") && strings.Contains(err.Error(), "number") {
			return nil, ErrUnique
		}

		return nil, err
	}

	u.Settings.NextCustomerNumber++
	db.Save(&u)

	return c, nil
}

// CustomerDelete removes a customer and it's relations
func CustomerDelete(uuid string, id string) error {
	var c Customer

	if db.
		Where("uuid = ? AND id = ?", uuid, id).
		First(&c).
		RecordNotFound() {
		return ErrCustomerNotFound
	}

	db.Delete(&c)

	return nil
}

// Update patches an existing customer
func (c *Customer) Update(patch *Customer) error {
	if err := patch.validate(); err != nil {
		return err
	}

	db.Model(c).Updates(patch)
	// this seems to be needed to update the contact list.
	db.Model(c).Association("Contacts").Replace(patch.Contacts)

	return nil
}

func (c *Customer) validate() error {
	if err := validate.Struct(c); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			r := strings.NewReplacer("{field}", v.Field(), "{param}", v.Param())
			e := r.Replace(errList[v.Tag()])

			// return first error found.
			return ValidationError{e}
		}
	}

	return nil
}
