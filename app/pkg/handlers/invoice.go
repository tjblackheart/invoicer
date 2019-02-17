package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tjblackheart/invoicer/pkg/db"
)

//const MULTIPLIER = 100

// InvoiceMigrate initialises the tables
func InvoiceMigrate() {
	db.Conn.AutoMigrate(&Invoice{}, &InvoiceItem{})
}

// Invoice represents an invoice
type Invoice struct {
	db.Model
	Number      string        `json:"number,omitempty" gorm:"unique;not null"`
	Date        time.Time     `json:"date,omitempty"`
	TotalNet    float64       `json:"total_net" gorm:"not null;default:0"`
	Currency    string        `json:"currency" gorm:"not null;default:'EUR'"`
	Items       []InvoiceItem `json:"items,omitempty" gorm:"foreignkey:InvoiceID"`
	IsCancelled bool          `json:"is_cancelled" gorm:"not null;default:false"`
	IsPaid      bool          `json:"is_paid" gorm:"not null;default:false"`
	PaidAt      time.Time     `json:"paid_at"`
	CustomerID  uint          `json:"customer_id" gorm:"not null"`
	Customer    Customer      `json:"customer"`
	UUID        string        `json:"-"`
}

// InvoiceItem represents a single invoice item
type InvoiceItem struct {
	db.Model
	Description  string  `json:"description,omitempty"`
	Amount       float64 `json:"amount,omitempty" gorm:"not null;default:0"`
	Unit         string  `json:"unit,omitempty" gorm:"not null;default:'hrs'"`
	PricePerUnit float64 `json:"price_per_unit,omitempty" gorm:"not null;default:0"`
	VAT          uint    `json:"vat,omitempty" gorm:"not null;default:19"`
	InvoiceID    uint    `json:"-"`
}

// BeforeCreate GORM handler for invoices
/*func (i *Invoice) BeforeCreate(tx *gorm.DB) (err error) {
	if i.Date.IsZero() {
		tx.Model(i).Update("date", time.Now())
	}

	totalNet := 0.0
	for _, item := range i.Items {
		totalNet += item.Amount * item.PricePerUnit
	}

	tx.Model(i).Update("total_net", totalNet)

	return
}*/

// InvoiceGetAll returns all invoces
func InvoiceGetAll(w http.ResponseWriter, r *http.Request) {

	invoices := []Invoice{}

	uuid, err := GetUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := db.
		Conn.
		Preload("Items").
		Preload("Customer").
		Preload("Customer.Address").
		Where("uuid = ?", uuid).
		Find(&invoices).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(invoices)
}

// InvoiceGet gets a single invoice by id
func InvoiceGet(w http.ResponseWriter, r *http.Request) {

	var invoice Invoice
	vars := mux.Vars(r)
	id := vars["id"]

	uuid, err := GetUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if db.
		Conn.
		Preload("Items").
		Preload("Customer").
		Preload("Customer.Address").
		Where("uuid = ? AND id = ?", uuid, id).
		First(&invoice).RecordNotFound() {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(invoice)
}

// InvoiceCreate creates a new invoice
func InvoiceCreate(w http.ResponseWriter, r *http.Request) {

	var invoice Invoice
	err := json.NewDecoder(r.Body).Decode(&invoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = invoice.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var customer Customer
	if db.Conn.First(&customer, invoice.CustomerID).RecordNotFound() {
		http.Error(w, "No such customer", http.StatusBadRequest)
		return
	}

	uuid, err := GetUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	invoice.UUID = uuid

	if invoice.Date.IsZero() {
		invoice.Date = time.Now()
	}

	for _, item := range invoice.Items {
		invoice.TotalNet += item.Amount * item.PricePerUnit
	}

	if err := db.Conn.Create(&invoice).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var u User
	if err = db.Conn.Preload("Settings").Where("uuid = ?", uuid).First(&u).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u.Settings.NextInvoiceNumber++
	db.Conn.Save(&u)

	db.Conn.
		Preload("Items").
		Preload("Customer").
		Preload("Customer.Address").
		Where("uuid = ?", uuid).
		First(&invoice, invoice.ID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invoice)
}

// InvoiceSetPaid toggles payment status
func InvoiceSetPaid(w http.ResponseWriter, r *http.Request) {

	type Payload struct {
		ID          int    `json:"id"`
		PaymentDate string `json:"date"`
	}

	payload := Payload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uuid, err := GetUUID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var invoice Invoice

	if db.
		Conn.
		Where("uuid = ? AND id = ?", uuid, payload.ID).
		First(&invoice).RecordNotFound() {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	paymentDate, err := time.Parse("2006-01-02", payload.PaymentDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	invoice.IsPaid = true
	invoice.PaidAt = paymentDate
	db.Conn.Save(&invoice)

	invoices := []Invoice{}

	if err := db.
		Conn.
		Preload("Items").
		Preload("Customer").
		Preload("Customer.Address").
		Where("uuid = ?", uuid).
		Find(&invoices).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(invoices)
}

// Validate validates customer data
func (i *Invoice) Validate() (err error) {
	if i.Number == "" {
		err = errors.New("Invoice number can not be empty")
	}

	if len(i.Items) == 0 {
		err = errors.New("No items found")
	}

	if i.CustomerID == 0 {
		err = errors.New("No customer given")
	}

	return
}
