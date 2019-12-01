package models

import (
	"errors"
	"strings"
)

// Customer represents a customer
type (
	Customer struct {
		BaseModel
		Number    string    `json:"number" gorm:"unique;not null"`
		Remarks   string    `json:"remarks"`
		TaxNumber string    `json:"tax_number"`
		Address   Address   `json:"address" gorm:"foreignkey:CustomerID"`
		Contacts  []Contact `json:"contacts" gorm:"foreignkey:CustomerID"`
		UUID      string    `json:"-"`
	}

	// Address holds an address of a customer
	Address struct {
		BaseModel
		Company    string `json:"company" gorm:"not null"`
		FirstName  string `json:"first_name" gorm:"not null"`
		LastName   string `json:"last_name" gorm:"not null"`
		Street     string `json:"street" gorm:"not null"`
		Number     string `json:"number" gorm:"not null"`
		Zip        string `json:"zip" gorm:"not null"`
		City       string `json:"city" gorm:"not null"`
		Country    string `json:"country" gorm:"not null"`
		CustomerID uint   `json:"-"`
	}

	// Contact holds a contact value (email, phone ...)
	Contact struct {
		BaseModel
		Value      string `json:"value" gorm:"not null"`
		Type       string `json:"type" gorm:"not null"`
		CustomerID uint   `json:"-"`
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

func (c *Customer) validate() (err error) {
	if c.Number == "" {
		err = errors.New("Customer number can not be empty")
	}

	if c.Address.Company == "" || c.Address.City == "" || c.Address.Zip == "" || c.Address.Street == "" {
		err = errors.New("Invalid address")
	}

	return
}
