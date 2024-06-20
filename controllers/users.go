package controllers

import (
	"fmt"
	"net/http"
)

// Users is used to handle the signup process.
type Users struct {
	Templates struct {
		New Template
	}
}

// New is used to render the signup form when a user visits /signup
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}

// Create is used to process the signup form when a user submits it.
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Email:", r.FormValue("email"))
	fmt.Fprint(w, "Password:", r.FormValue("password"))
}

// I can use Gorrila Schema to validate the form data, here just show the basic usage of form data with standard library
