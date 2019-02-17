package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/tjblackheart/invoicer/pkg/db"
)

// Customer represents a customer
type Customer struct {
	db.Model
	Number    string    `json:"number,omitempty" gorm:"unique;not null"`
	Remarks   string    `json:"remarks,omitempty"`
	TaxNumber string    `json:"tax_number"`
	Address   Address   `json:"address" gorm:"foreignkey:CustomerID"`
	Contacts  []Contact `json:"contacts" gorm:"foreignkey:CustomerID"`
	UUID      string    `json:"-"`
}

// Address holds an address of a customer
type Address struct {
	db.Model
	Company    string `json:"company,omitempty" gorm:"not null"`
	FirstName  string `json:"first_name,omitempty" gorm:"not null"`
	LastName   string `json:"last_name,omitempty" gorm:"not null"`
	Street     string `json:"street,omitempty" gorm:"not null"`
	Number     string `json:"number,omitempty" gorm:"not null"`
	Zip        string `json:"zip,omitempty" gorm:"not null"`
	City       string `json:"city,omitempty" gorm:"not null"`
	Country    string `json:"country,omitempty" gorm:"not null"`
	CustomerID uint   `json:"-"`
}

// Contact holds a contact value (email, phone ...)
type Contact struct {
	db.Model
	Value      string `json:"value,omitempty" gorm:"not null"`
	Type       string `json:"type,omitempty" gorm:"not null"`
	CustomerID uint   `json:"-"`
}

// CustomerMigrate initialises the tables
func CustomerMigrate() {
	db.Conn.AutoMigrate(&Customer{}, &Address{}, &Contact{})
}

// CustomerGetAll returns a list of all customers
func CustomerGetAll(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{}
	uuid, err := GetUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.
		Conn.
		Preload("Address").
		Preload("Contacts").
		Where("uuid = ?", uuid).
		Find(&customers).
		Error

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(customers)
}

// CustomerGet returns a single customer
func CustomerGet(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	vars := mux.Vars(r)
	id := vars["id"]

	uuid, err := GetUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if db.
		Conn.
		Preload("Address").
		Preload("Contacts").
		Where("uuid = ? AND id = ?", uuid, id).
		First(&customer).
		RecordNotFound() {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

// CustomerCreate creates a new customer
func CustomerCreate(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = customer.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uuid, err := GetUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	customer.UUID = uuid

	if err := db.Conn.Create(&customer).Error; err != nil {
		e := ""
		if strings.Contains(err.Error(), "UNIQUE") && strings.Contains(err.Error(), "number") {
			e = "The customer number already exists"
		} else {
			e = err.Error()
		}

		http.Error(w, e, http.StatusBadRequest)
		return
	}

	var u User
	if err = db.Conn.Preload("Settings").Where("uuid = ?", uuid).First(&u).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u.Settings.NextCustomerNumber++
	db.Conn.Save(&u)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

// CustomerUpdate updates a customer
func CustomerUpdate(w http.ResponseWriter, r *http.Request) {
	var customer Customer

	vars := mux.Vars(r)
	id := vars["id"]

	uuid, err := GetUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if db.Conn.
		Preload("Address").
		Preload("Contacts").
		Where("uuid = ? AND id = ?", uuid, id).
		First(&customer).
		RecordNotFound() {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var patch Customer
	err = json.NewDecoder(r.Body).Decode(&patch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = patch.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Conn.Model(&customer).Updates(patch)
	// this seems to be needed to update the contact list.
	db.Conn.Model(&customer).Association("Contacts").Replace(patch.Contacts)

	json.NewEncoder(w).Encode(customer)
}

// CustomerDelete removes a customer
func CustomerDelete(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	vars := mux.Vars(r)
	id := vars["id"]

	uuid, err := GetUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if db.Conn.
		Where("uuid = ? AND id = ?", uuid, id).
		First(&customer).
		RecordNotFound() {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	db.Conn.Delete(&customer)
}

// Validate validates customer data
func (c *Customer) Validate() (err error) {
	if c.Number == "" {
		err = errors.New("Customer number can not be empty")
	}

	if c.Address.Company == "" || c.Address.City == "" || c.Address.Zip == "" || c.Address.Street == "" {
		err = errors.New("Invalid address")
	}

	return
}
