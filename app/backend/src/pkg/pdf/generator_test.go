package pdf

import (
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/tjblackheart/Invoicer/backend/pkg/models"
	"github.com/tjblackheart/Invoicer/backend/pkg/money"
)

var mockInvoice = models.Invoice{
	Number:     "T0001",
	Date:       time.Now(),
	TotalNet:   money.Money(10000),
	TotalGross: money.Money(11900),
	Currency:   "EUR",
	Items: []models.InvoiceItem{
		models.InvoiceItem{
			Description:  "Testitem",
			Amount:       2,
			Unit:         "hrs",
			PricePerUnit: money.Money(5000),
			VAT:          19.0,
		},
	},
	Customer: models.Customer{
		Number:    "C100",
		TaxNumber: "TAX123",
		Address: models.Address{
			Company:   "Test Corp",
			FirstName: "Cus",
			LastName:  "Tomer",
			Street:    "Customerstreet",
			Number:    "999",
			Zip:       "11111",
			City:      "City",
			Country:   "Germany",
		},
	},
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
		TaxNumber: "TEST465",
		Email:     "test@test.com",
		Phone:     "01234567",
	},
}

var g = Generator{Invoice: &mockInvoice, User: &mockUser}

// https://brandur.org/fragments/testing-go-project-root
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../")
	if err := os.Chdir(dir); err != nil {
		panic(err)
	}
}

func TestGenerate(t *testing.T) {
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

func TestHtml(t *testing.T) {
	html, err := g.toHTML()
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

	tests := []struct {
		needle string
	}{
		{mockInvoice.Number},
		{mockInvoice.Date.Format("02.01.2006")},
		{mockInvoice.TotalNet.Format()},
		{mockInvoice.TotalGross.Format()},
		{mockInvoice.Currency},
		{mockInvoice.Customer.TaxNumber},
		{mockInvoice.Customer.Address.Company},
		{mockInvoice.Customer.Address.FirstName},
		{mockInvoice.Customer.Address.LastName},
		{mockInvoice.Customer.Address.Street},
		{mockInvoice.Customer.Address.City},
		{mockInvoice.Customer.Address.Zip},
		{mockUser.Settings.Company},
		{mockUser.Settings.FirstName},
		{mockUser.Settings.LastName},
		{mockUser.Settings.City},
		{mockUser.Settings.Street},
		{mockUser.Settings.Zip},
	}

	for _, tt := range tests {
		if strings.Contains(html, tt.needle) == false {
			t.Errorf("want string: %s - not found", tt.needle)
		}
	}
}

func TestFooter(t *testing.T) {
	g.generateFooter()

	tests := []struct {
		haystack string
		needle   string
	}{
		{g.footer.Left, mockUser.Settings.FirstName},
		{g.footer.Left, mockUser.Settings.LastName},
		{g.footer.Left, mockUser.Settings.Street},
		{g.footer.Left, mockUser.Settings.Number},
		{g.footer.Left, mockUser.Settings.Zip},
		{g.footer.Center, mockUser.Settings.Bank},
		{g.footer.Center, mockUser.Settings.IBAN},
		{g.footer.Center, mockUser.Settings.BIC},
		{g.footer.Right, mockUser.Settings.TaxNumber},
		{g.footer.Right, mockUser.Settings.Email},
		{g.footer.Right, mockUser.Settings.Phone},
	}

	for _, tt := range tests {
		if strings.Contains(tt.haystack, tt.needle) == false {
			t.Errorf("have: %s, want: %s", tt.haystack, tt.needle)
		}
	}
}
