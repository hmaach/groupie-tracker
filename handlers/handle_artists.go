package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"groupie_tracker/models"
	"groupie_tracker/utils"
)

// ArtistHandler handles requests to view details of a specific artist.
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var artist models.Artist
	var location models.Location
	var relation models.Relation
	var date models.Date

	// Fetch artist details
	if err := utils.Fetch("/artists/"+id, &artist); err != nil {
		RenderError(w, http.StatusNotFound, "404 | Artist not found.")
		return
	}

	// Fetch artist locations
	if err := utils.Fetch("/locations/"+id, &location); err != nil {
		fmt.Println(err)
		RenderError(w, http.StatusInternalServerError, "500 | Failed to retrieve locations.")
		return
	}

	// Fetch artist relations
	if err := utils.Fetch("/relation/"+id, &relation); err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Failed to retrieve relations.")
		return
	}

	// Fetch concert dates
	if err := utils.Fetch("/dates/"+id, &date); err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Failed to retrieve dates.")
		return
	}

	// Normalize location names
	location.Locations = utils.NormalizeLocations(location.Locations)
	relation.DatesLocations = utils.NormalizeDatesLocations(relation.DatesLocations)

	// Add fetched data to the artist
	artist.Location = location
	artist.Relation = relation
	artist.Date = date

	// Render the artist details template
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

	var artists []models.ArtistSummary
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

	if err := RenderTemplate(w, "index.html", data); err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Failed to render the page.")
	}
}
