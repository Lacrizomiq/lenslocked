package controllers

import "net/http"

// Template is a type that is capable of executing a template and writing the output to a http.ResponseWriter.
type Template interface {
	Execute(w http.ResponseWriter, data interface{})
}
