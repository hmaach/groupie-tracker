package handlers

import (
	"net/http"
	"os"
	"strings"
)

// AssetsHandler handles requests for static assets.
func AssetsHandler(w http.ResponseWriter, r *http.Request) {
	// Block direct access to the /assets/ directory
	if r.URL.Path == "/assets/" || strings.HasSuffix(r.URL.Path, "/") {
		RenderError(w, http.StatusForbidden, "403 | Access to this resource is forbidden !")
		return
	}

	// Serve the asset file
	filePath := "./assets" + strings.TrimPrefix(r.URL.Path, "/assets")
	if !isFileExists(filePath) {
		// If the requested file or path does not exist, render a styled 404 page
		RenderError(w, http.StatusNotFound, "404 | Page Not Found")
		return
	}

	// File exists, serve it
	http.ServeFile(w, r, filePath)
}

// isFileExists checks if a file exists at the given path
func isFileExists(filePath string) bool {
	if filePath == "" {
		// Prevent checking an empty path
		return false
	}

	info, err := os.Stat(filePath)
	if err != nil || info.IsDir() {
		// Return false if the file doesn't exist or it's a directory
		return false
	}
	return true
}
