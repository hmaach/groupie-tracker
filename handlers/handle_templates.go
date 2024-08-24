package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie_tracker/models"
)

// RenderError renders an error page with a specific status code and message.
func RenderError(w http.ResponseWriter, statusCode int, message string) {
	err := RenderTemplate(w, "error.html", models.ErrorData{
		Error:   http.StatusText(statusCode),
		Code:    statusCode,
		Message: message,
	})
	if err != nil {
		http.Error(w, "500 | Internal Server Error!", http.StatusInternalServerError)
		fmt.Println("Error parsing template:", err)
		return
	}
	w.WriteHeader(statusCode)
}

// RenderTemplate renders a given HTML template with provided data
func RenderTemplate(w http.ResponseWriter, tmpl string, data any) error {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		return err
	}
	err = t.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		return err
	}
	return nil
}
