package pdf

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"os"
	"strings"

	wk "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/tjblackheart/invoicer/pkg/models"
	"github.com/tjblackheart/invoicer/pkg/money"
)

type tplData struct {
	Invoice *models.Invoice
	User    *models.User
}

type footer struct {
	Left, Center, Right string
}

// Generate prints an invoice to PDF and returns the filename
func Generate(i *models.Invoice, u *models.User) (string, error) {
	html, err := toHTML(i, u)
	if err != nil {
		return "", err
	}

	filename := i.Number + ".pdf"
	footerData := generateFooter(&u.Settings)

	if err = toPDF(html, filename, footerData); err != nil {
		return "", err
	}

	return filename, nil
}

// Base64 converts a file to a Base64 encoded string.
func Base64(filename string) (string, error) {
	file, err := os.Open("var/out/" + filename)
	defer file.Close()

	if err != nil {
		return "", err
	}

	stat, err := file.Stat()
	if err != nil {
		return "", err
	}

	size := stat.Size()
	buf := make([]byte, size)

	r := bufio.NewReader(file)
	if _, err = r.Read(buf); err != nil {
		return "", err
	}

	s := base64.StdEncoding.EncodeToString(buf)

	return s, nil
}

func toHTML(i *models.Invoice, u *models.User) (string, error) {
	funcs := template.FuncMap{
		"add": func(i int) int {
			return i + 1
		},
		"itemNet": func(p money.Money, a float64) string {
			return p.Multiply(a).Format()
		},
		"itemGross": func(perUnit money.Money, amount float64, vat float64) string {
			tax := perUnit.Multiply(vat / 100)
			single := perUnit + tax

			return single.Multiply(amount).Format()
		},
		"tax": func(gross money.Money, net money.Money) string {
			tax := gross - net
			return tax.Format()
		},
	}

	name := "invoice"
	ts, err := template.New(name).Funcs(funcs).ParseFiles("tpl/invoice-en.html")
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = ts.ExecuteTemplate(buf, name, &tplData{i, u}); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func toPDF(html string, filename string, footerData footer) error {
	// use file if it's already there
	if _, err := os.Stat("var/out/" + filename); err == nil {
		return nil
	}

	gen, err := wk.NewPDFGenerator()
	if err != nil {
		return err
	}

	gen.Grayscale.Set(true)
	gen.Dpi.Set(300)

	page := wk.NewPageReader(strings.NewReader(html))

	page.FooterFontSize.Set(8)
	page.FooterLeft.Set(footerData.Left)
	page.FooterCenter.Set(footerData.Center)
	page.FooterRight.Set(footerData.Right)

	gen.AddPage(page)

	if err = gen.Create(); err != nil {
		return err
	}

	if err = gen.WriteFile("var/out/" + filename); err != nil {
		return err
	}

	return nil
}

func generateFooter(s *models.Settings) footer {
	f := footer{}

	f.Left = fmt.Sprintf(`
        %s
        %s %s
        %s %s
        %s, %s
    `, s.Company, s.FirstName, s.LastName, s.Street, s.Number, s.City, s.Zip)

	f.Center = fmt.Sprintf(`
        Bank: %s
        IBAN: %s
        BIC: %s
    `, s.Bank, s.IBAN, s.BIC)

	f.Right = fmt.Sprintf(`
		VAT ID: %s
        %s
        %s
        Page [page]/[topage]
    `, s.TaxNumber, s.Email, s.Phone)

	return f
}
