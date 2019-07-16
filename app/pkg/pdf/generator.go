package pdf

import (
	"bytes"
	"html/template"
	"strings"

	wk "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/tjblackheart/invoicer/pkg/models"
)

// Generate prints an invoice to PDF and returns the filename
func Generate(i *models.Invoice, u *models.User) (string, error) {
	html, err := toHTML(i, u)
	if err != nil {
		return "", err
	}

	filename := i.Number + ".pdf"

	err = toPDF(html, filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func toHTML(i *models.Invoice, u *models.User) (string, error) {
	name := "invoice"
	data := map[string]string{
		"Title": i.Number,
		// TODO
	}

	ts, err := template.New(name).ParseFiles("tpl/invoice.gohtml")
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = ts.ExecuteTemplate(buf, name, &data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func toPDF(html string, filename string) error {
	gen, err := wk.NewPDFGenerator()
	if err != nil {
		return err
	}

	gen.Grayscale.Set(true)
	gen.Dpi.Set(600)

	page := wk.NewPageReader(strings.NewReader(html))
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)

	gen.AddPage(page)

	err = gen.Create()
	if err != nil {
		return err
	}

	err = gen.WriteFile("out/" + filename)
	if err != nil {
		return err
	}

	return nil
}
