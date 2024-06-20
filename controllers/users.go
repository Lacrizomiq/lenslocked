package controllers

import (
	"net/http"

	"github.com/joncalhoun/lenslocked/views"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// I need a view to render
	u.Templates.New.Execute(w, nil)
}
