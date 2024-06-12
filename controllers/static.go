package controllers

import (
	"html/template"
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

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "What is LensLocked?",
			Answer:   "LensLocked is a photo gallery application that allows you to view photos and create your own account to upload your own photos.",
		},
		{
			Question: "How do I create an account?",
			Answer:   "To create an account, click the 'Sign Up' link in the top right corner of the page and fill out the form.",
		},
		{
			Question: "How do I upload a photo?",
			Answer:   "To upload a photo, you must first create an account. Once you have an account, click the 'Upload' link in the top right corner of the page and follow the instructions.",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>`,
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
