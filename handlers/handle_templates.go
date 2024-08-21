package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie_tracker/models"
)

// RenderError renders an error page with a specific status code and message.
func RenderError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	RenderTemplate(w, "error.html", models.ErrorData{
		Error:   http.StatusText(statusCode),
		Code:    statusCode,
		Message: message,
	})
}

// RenderTemplate renders a given HTML template with provided data
func RenderTemplate(w http.ResponseWriter, tmpl string, data any) error {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, "500 | Internal Server Error!", http.StatusInternalServerError)
		fmt.Println("Error parsing template:", err)
		return err
	}

	err = t.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, "500 | Internal Server Error!", http.StatusInternalServerError)
		fmt.Println("Error executing template:", err)
		return err
	}

	return nil
}
