package pdf

import (
	"bufio"
	"bytes"
	"encoding/base64"
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

// Generate prints an invoice to PDF and returns the filename
func Generate(i *models.Invoice, u *models.User) (string, error) {
	html, err := toHTML(i, u)
	if err != nil {
		return "", err
	}

	filename := i.Number + ".pdf"

	if err = toPDF(html, filename); err != nil {
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
	r.Read(buf)
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

func toPDF(html string, filename string) error {
	// use file if it's already there
	// if _, err := os.Stat("var/out/" + filename); err == nil {
	// 	return nil
	// }

	gen, err := wk.NewPDFGenerator()
	if err != nil {
		return err
	}

	gen.Grayscale.Set(true)
	gen.Dpi.Set(300)

	page := wk.NewPageReader(strings.NewReader(html))
	page.FooterRight.Set("[page]/[topage]")
	page.FooterFontSize.Set(10)

	gen.AddPage(page)

	if err = gen.Create(); err != nil {
		return err
	}

	if err = gen.WriteFile("var/out/" + filename); err != nil {
		return err
	}

	return nil
}
