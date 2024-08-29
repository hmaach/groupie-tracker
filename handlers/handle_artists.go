package handlers

import (
	"net/http"
	"strconv"

	"groupie_tracker/models"
	"groupie_tracker/utils"
)

// ArtistHandler handles requests to view details of a specific artist.
// It takes an HTTP ResponseWriter to write the response and an HTTP Request.
// It fetches the artist details, locations, relations, and concert dates from the API.
// If successful, it renders the artist details page; otherwise, it renders an appropriate error page.
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RenderError(w, http.StatusMethodNotAllowed, "405 | Method Not Allowed: Use GET")
		return
	}
	id := r.PathValue("id")

	var artist models.Artist

	artist, err := utils.FetchArtist(id)
	if err != nil {
		if err.Error() == "404" {
			RenderError(w, http.StatusNotFound, "404 | Artist Not Found")
		} else {
			RenderError(w, http.StatusInternalServerError, "500 | Status Internal Server Error")
		}
		return
	}

	// Render the artist details template
	if err := RenderTemplate(w, "artist.html", http.StatusOK, artist); err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Failed to render artist detail page.")
	}
}

// MainHandler handles requests to the root URL and displays the list of artists.
// It takes an HTTP ResponseWriter to write the response and an HTTP Request.
// It fetches a summary list of all artists, assigns a type based on the number of members,
// and renders the index page. If any errors occur, it renders the appropriate error page.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderError(w, http.StatusNotFound, "404 | The page you are looking for does not exist.")
		return
	}

	if r.Method != http.MethodGet {
		RenderError(w, http.StatusMethodNotAllowed, "405 | Method Not Allowed: Use GET")
		return
	}

	var artists []models.Artist
	err := utils.Fetch("/artists", &artists)
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Failed to retrieve artists.")
		return
	}

	// Set the Type field based on the number of members
	for i := range artists {
		if len(artists[i].Members) == 1 {
			artists[i].Type = "Artist"
		} else {
			artists[i].Type = "Group of " + strconv.Itoa(len(artists[i].Members))
		}
	}

	data := models.ArtistsPageData{Artists: artists}

	if err := RenderTemplate(w, "index.html", http.StatusOK, data); err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Failed to render the page.")
		return
	}
}
