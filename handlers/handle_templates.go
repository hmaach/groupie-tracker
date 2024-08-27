package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var Templates *template.Template

// ParseTemplates parses all HTML templates from the "templates" directory 
// and stores them in the global variable "Templates".
// If parsing fails, it logs  error.
func ParseTemplates() {
	var err error
	Templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Printf("Failed to parse templates: %v", err)
	}
}

// RenderTemplate renders a pre-parsed template with the provided data.
// It takes an HTTP ResponseWriter to write the response, the name of the template to render (tmpl),
// the HTTP status code, and the data to pass to the template.
// If the template is not found, it returns a 500 error. Otherwise, it returns any error encountered during rendering.
func RenderTemplate(w http.ResponseWriter, tmpl string, statusCode int, data any) error {
	t := Templates.Lookup(tmpl)
	if t == nil {
		http.Error(w, "500 | Internal Server Error", http.StatusInternalServerError)
		return fmt.Errorf("template %s not found", tmpl)
	}

	w.WriteHeader(statusCode)
	return t.Execute(w, data)
}
