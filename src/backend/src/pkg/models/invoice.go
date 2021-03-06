package models

import (
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/tjblackheart/Invoicer/backend/pkg/money"
)

// Invoice represents an invoice
type (
	Invoice struct {
		BaseModel
		Number      trimmed       `json:"number" gorm:"unique;not null" validate:"required"`
		Date        time.Time     `json:"date" validate:"required"`
		TotalNet    money.Money   `json:"total_net" gorm:"not null;default:0"`
		TotalGross  money.Money   `json:"total_gross" gorm:"not null;default:0"`
		Currency    string        `json:"currency" gorm:"not null;default:'EUR'" validate:"required"`
		Items       []InvoiceItem `json:"items" gorm:"foreignkey:InvoiceID" validate:"gt=0"`
		IsCancelled bool          `json:"is_cancelled" gorm:"not null;default:false"`
		IsPaid      bool          `json:"is_paid" gorm:"not null;default:false"`
		PaidAt      time.Time     `json:"paid_at"`
		CustomerID  uint          `json:"customer_id" gorm:"not null" validate:"required"`
		Customer    Customer      `json:"customer" validate:"-"`
		DueDays     int           `json:"due_days" gorm:"default:10"`
		UUID        string        `json:"-"`
	}

	// InvoiceItem represents a single invoice item
	InvoiceItem struct {
		BaseModel
		Description  string      `json:"description"`
		Amount       float64     `json:"amount" gorm:"not null;default:0"`
		Unit         string      `json:"unit" gorm:"not null;default:'hrs'"`
		PricePerUnit money.Money `json:"price_per_unit" gorm:"not null;default:0"`
		VAT          float64     `json:"vat" gorm:"not null;default:19"`
		InvoiceID    uint        `json:"-"`
	}

	// PaidPayload holds a json body for payment requests
	PaidPayload struct {
		ID          int    `json:"id"`
		PaymentDate string `json:"date"`
	}
)

// FindInvoice returns a single invoice
func FindInvoice(uuid string, id string) (*Invoice, error) {
	var invoice Invoice

	if db.
		Preload("Items").
		Preload("Customer").
		Preload("Customer.Address").
		Where("uuid = ? AND id = ?", uuid, id).
		First(&invoice).RecordNotFound() {
		return nil, ErrNotFound{Message: "Invoice not found"}
	}

	return &invoice, nil
}

// InvoiceGetAll finds all invoices for a user
func InvoiceGetAll(uuid string) (*[]Invoice, error) {
	invoices := []Invoice{}

	if err := db.
		Preload("Items").
		Preload("Customer").
		Preload("Customer.Address").
		Where("uuid = ?", uuid).
		Order("date DESC").
		Find(&invoices).Error; err != nil {
		return nil, err
	}

	return &invoices, nil
}

// Create creates a new invoice
func (i *Invoice) Create(uuid string) error {
	u, err := FindUser(uuid)
	if err != nil {
		return err
	}

	i.UUID = u.UUID

	if i.Date.IsZero() {
		i.Date = time.Now()
	}

	for _, item := range i.Items {
		net := item.PricePerUnit.Multiply(item.Amount)
		i.TotalNet += net

		tax := net.Multiply(item.VAT / 100)
		i.TotalGross += net + tax
	}

	if err := i.validate(); err != nil {
		return err
	}

	if err := db.Create(i).Error; err != nil {
		return err
	}

	u.Settings.NextInvoiceNumber++
	db.Save(&u)

	return nil
}

// InvoiceSetPaid sets payment flag
func InvoiceSetPaid(uuid string, payload *PaidPayload) error {
	paymentDate, err := time.Parse("2006-01-02", payload.PaymentDate)
	if err != nil {
		return err
	}

	id := strconv.Itoa(payload.ID)
	invoice, err := FindInvoice(uuid, id)
	if err != nil {
		return err
	}

	invoice.IsPaid = true
	invoice.PaidAt = paymentDate
	db.Save(&invoice)

	return nil
}

func (i Invoice) validate() (err error) {
	if err := validate.Struct(i); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			r := strings.NewReplacer("{field}", v.Field(), "{param}", v.Param())
			e := r.Replace(errList[v.Tag()])

			// return first error found.
			return ErrValidation{Message: e}
		}
	}

	return nil
}
