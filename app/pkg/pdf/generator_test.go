package pdf

import (
	"os"
	"path"
	"regexp"
	"runtime"
	"testing"
	"time"

	"github.com/tjblackheart/invoicer/pkg/models"
	"github.com/tjblackheart/invoicer/pkg/money"
)

var mockInvoice = models.Invoice{
	Number:     "T0001",
	Date:       time.Now(),
	TotalNet:   100.00,
	TotalGross: 119.00,
	Currency:   "EUR",
	Items: []models.InvoiceItem{
		models.InvoiceItem{
			Description:  "Testitem",
			Amount:       2,
			Unit:         "hrs",
			PricePerUnit: money.Money(50),
			VAT:          19.0,
			InvoiceID:    1,
		},
	},
	IsCancelled: false,
	IsPaid:      false,
	Customer:    models.Customer{},
}

var mockUser = models.User{
	Settings: models.Settings{
		Company:   "Testcompany",
		FirstName: "Test",
		LastName:  "Tester",
		Street:    "Teststreet",
		Number:    "123a",
		City:      "Testing",
		Zip:       "12345",
		Bank:      "Testbank",
		IBAN:      "TT1234",
		BIC:       "TESTFF",
		TaxNumber: "TEST465",
		Email:     "test@test.com",
		Phone:     "12345",
	},
}

// https://brandur.org/fragments/testing-go-project-root
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../")
	if err := os.Chdir(dir); err != nil {
		panic(err)
	}
}

func TestGenerate(t *testing.T) {
	g := Generator{Invoice: &mockInvoice, User: &mockUser}
	fname, err := g.Generate()
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

	if fname != "T0001.pdf" {
		t.Errorf("wrong filename, have: %s (%T), want: T0001.pdf (string)", fname, fname)
	}

	if _, err := os.Stat("var/out/" + fname); err != nil {
		t.Errorf("file var/out/%s does not exist", fname)
	}

	os.Remove("var/out/" + fname)
}

func TestBase64(t *testing.T) {
	g := Generator{Invoice: &mockInvoice, User: &mockUser}
	fname, err := g.Generate()
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

	b64, err := g.Base64(fname)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

	matched, _ := regexp.MatchString(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`, b64)
	if matched == false {
		t.Errorf("not a Base64 encoded string: %s", b64)
	}
}
