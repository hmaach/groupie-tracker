package handlers

import (
	"net/http"

	"groupie_tracker/models"
	"groupie_tracker/utils"
)

// ArtistHandler handles requests to view details of a specific artist.
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var artist models.Artist
	err := utils.Fetch("/artists/"+id, &artist)
	if err != nil {
		RenderError(w, http.StatusNotFound, "404 | Artist not found.")
		return
	}

	if err := RenderTemplate(w, "artist.html", artist); err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Failed to render artist detail page.")
	}
}

// MainHandler handles requests to the root URL and displays the list of artists.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderError(w, http.StatusNotFound, "404 | The page you are looking for does not exist.")
		return
	}

	var artists []models.Artist
	err := utils.Fetch("/artists", &artists)
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Failed to retrieve artists.")
		return
	}

	data := models.ArtistsPageData{Artists: artists}
	if err := RenderTemplate(w, "index.html", data); err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Failed to render the page.")
	}
}
