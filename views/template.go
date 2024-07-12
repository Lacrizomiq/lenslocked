package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
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
	// Create a new template and add a function to generate a CSRF field.
	// This is a contrived example to show how to add functions to templates.
	tpl := template.New(pattern[0])
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return `<input type="hidden" />`
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

// Parse parses the provided file and returns a Template.
// If there is an error parsing the template, an error is returned.
// func Parse(filepath string) (Template, error) {
// 	htmlTpl, err := template.ParseFiles(filepath)
// 	if err != nil {
// 		return Template{}, fmt.Errorf("parsing template: %w", err)
// 	}
// 	return Template{
// 		htmlTpl: htmlTpl,
// 	}, nil
// }

// Template is a wrapper around html/template to provide additional helper methods.
type Template struct {
	htmlTpl *template.Template
}

// Execute executes the template, writing the output to w.
// If there is an error executing the template, an error is logged and an http.Error is written to w.
func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}
