package controllers

import (
	"html/template"
	"net/http"
)

// StaticHandler returns a handler that will render the static page
func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
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
			Question: "How do I delete a photo?",
			Answer:   "To delete a photo, you must first be logged in. Once you are logged in, navigate to the photo you want to delete and click the 'Delete' button.",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>`,
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}
}
