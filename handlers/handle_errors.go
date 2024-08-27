package handlers

import (
	"fmt"
	"net/http"

	"groupie_tracker/models"
)

var Error500 bool

// RenderError renders an error page with a specific status code and message.
// It takes an HTTP ResponseWriter to write the response, an HTTP status code, 
// and a message to display on the error page. If rendering the template fails, 
// it falls back to a generic 500 Internal Server Error page.
func RenderError(w http.ResponseWriter, statusCode int, message string) {
	
	err := RenderTemplate(w, "error.html", statusCode, models.ErrorData{
		Error:   http.StatusText(statusCode),
		Code:    statusCode,
		Message: message,
	})

	if err != nil {
		http.Error(w, "500 | Internal Server Error!", http.StatusInternalServerError)
		fmt.Println("Error parsing template:", err)
		return
	}
}