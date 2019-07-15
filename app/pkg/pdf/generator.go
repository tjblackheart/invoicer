package pdf

import (
	"bytes"
	"fmt"
	"html/template"

	wkp "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/tjblackheart/invoicer/pkg/models"
)

// Generate prints an invoice to PDF and returns the filename
func Generate(i *models.Invoice) (string, error) {
	html, err := toHTML(i)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// TODO
	fmt.Println(html)

	return "", nil
}

func toHTML(i *models.Invoice) (string, error) {
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

func toPDF(html string) (string, error) {
	generator, err := wkp.NewPDFGenerator()
	if err != nil {
		return "", err
	}

	// TODO
	fmt.Println(generator)

	return "", nil
}
