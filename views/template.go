package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

// Must is a helper that wraps a call to a function returning a Template and an error.
// It panics if the error is non-nil and otherwise returns the Template.
func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

// ParseFS parses the provided patterns from the provided fs.FS and returns a Template.
// If there is an error parsing the template, an error is returned.
func ParseFS(fs fs.FS, pattern ...string) (Template, error) {
	// Create a new template and parse the provided patterns.
	tpl := template.New(pattern[0])
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("csrfField is not implemented")
			},
		},
	)

	tpl, err := tpl.ParseFS(fs, pattern...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

// Template is a wrapper around html/template to provide additional helper methods.
type Template struct {
	htmlTpl *template.Template
}

// Execute executes the template, writing the output to w.
// If there is an error executing the template, an error is logged and an http.Error is written to w.
func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	tpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Printf("cloning template: %v", err)
		http.Error(w, "There was an error processing your request.", http.StatusInternalServerError)
		return
	}
	// Add the csrfField function to the template.
	tpl = tpl.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrf.TemplateField(r)
		},
	},
	)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Create a buffer to write the template output to.
	// If there is an error writing the template output, an error is logged and an http.Error is written to w.
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}
