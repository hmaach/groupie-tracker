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
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// File does not exist, use custom 404 page
		RenderError(w, http.StatusNotFound, "404 | Page Not Found")
		return
	}

	// File exists, serve it
	http.ServeFile(w, r, filePath)
}
