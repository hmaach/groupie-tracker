package groupie_tracker

import (
	"fmt"
	"html/template"
	"net/http"
)

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
