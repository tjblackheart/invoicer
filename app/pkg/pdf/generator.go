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
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/tjblackheart/invoicer/pkg/models"
	"github.com/tjblackheart/invoicer/pkg/money"
)

// Generator represents a generator instance
type Generator struct {
	Invoice *models.Invoice
	User    *models.User
	footer  struct {
		Left, Center, Right string
	}
}

// Generate prints an invoice to PDF and returns the filename
func (g *Generator) Generate() (string, error) {
	html, err := g.toHTML()
	if err != nil {
		return "", err
	}

	filename := g.Invoice.Number + ".pdf"

	if err = g.toPDF(html, filename); err != nil {
		return "", err
	}

	return filename, nil
}

// Base64 converts a file to a Base64 encoded string.
func (g *Generator) Base64(filename string) (string, error) {
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

func (g *Generator) toHTML() (string, error) {
	extensions := parser.HardLineBreak | parser.Strikethrough
	parser := parser.NewWithExtensions(extensions)

	funcs := template.FuncMap{
		"add": func(i int) int {
			return i + 1
		},
		"itemNet": func(net money.Money, amount float64) string {
			return net.Multiply(amount).Format()
		},
		"itemGross": func(perUnit money.Money, amount, vat float64) string {
			tax := perUnit.Multiply(vat / 100)
			single := perUnit + tax

			return single.Multiply(amount).Format()
		},
		"tax": func(gross, net money.Money) string {
			tax := gross - net
			return tax.Format()
		},
		"markdown": func(s string) template.HTML {
			html := markdown.ToHTML([]byte(s), parser, nil)
			return template.HTML(html)
		},
	}

	name := "invoice"
	ts, err := template.New(name).Funcs(funcs).ParseFiles("tpl/invoice-en.html")
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = ts.ExecuteTemplate(buf, name, g); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g *Generator) toPDF(html string, filename string) error {
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

	g.generateFooter()
	page.FooterFontSize.Set(8)
	page.FooterLeft.Set(g.footer.Left)
	page.FooterCenter.Set(g.footer.Center)
	page.FooterRight.Set(g.footer.Right)

	gen.AddPage(page)

	if err = gen.Create(); err != nil {
		return err
	}

	if err = gen.WriteFile("var/out/" + filename); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateFooter() {
	s := g.User.Settings
	f := &g.footer

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
}
