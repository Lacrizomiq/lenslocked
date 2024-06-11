package controllers

import (
	"net/http"

	"github.com/joncalhoun/lenslocked/views"
)

// Static is a controller that renders static pages
type Static struct {
	Template views.Template
}

// NewStatic returns a new Static
func (static Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}

// StaticHandler returns a handler that will render the static page
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
