package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"groupie_tracker/data"
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
	if r.URL.Path != "/" && r.URL.Path != "/filter" {
		RenderError(w, http.StatusNotFound, "404 | The page you are looking for does not exist.")
		return
	}

	if r.Method != http.MethodGet {
		RenderError(w, http.StatusMethodNotAllowed, "405 | Method Not Allowed: Use GET")
		return
	}

	var allData models.Output
	allData.For_search = data.CombinedData

	if r.URL.Path == "/filter" {
		// Parse query parameters for range filters with error handling
		creationDate1, err := strconv.Atoi(r.URL.Query().Get("creation-date-1"))
		if err != nil {
			creationDate1 = 1950
		}

		creationDate2, err := strconv.Atoi(r.URL.Query().Get("creation-date-2"))
		if err != nil {
			creationDate2 = 2024
		}

		firstAlbum1, err := strconv.Atoi(r.URL.Query().Get("first-album-1"))
		if err != nil {
			firstAlbum1 = 1950
		}

		firstAlbum2, err := strconv.Atoi(r.URL.Query().Get("first-album-2"))
		if err != nil {
			firstAlbum2 = 2024
		}
		if creationDate1 > creationDate2 {
			creationDate1, creationDate2 = creationDate2, creationDate1
		}
		if firstAlbum1 > firstAlbum2 {
			firstAlbum1, firstAlbum2 = firstAlbum2, firstAlbum1
		}

		// Get members filter
		membersStr := r.URL.Query()["members"]
		var members []int
		for _, ms := range membersStr {
			memberVal, err := strconv.Atoi(ms)
			if err == nil {
				members = append(members, memberVal)
			}
		}

		location := r.URL.Query().Get("location")
		location = strings.ReplaceAll(location, ", ", "-")
		// Filter the data using the provided criteria
		filteredData := utils.FilterData(data.CombinedData, creationDate1, creationDate2, firstAlbum1, firstAlbum2, location, members)

		// Create a new CombinedData structure for To_displayed
		allData.To_displayed = models.CombinedData{
			Artists:   filteredData,           // Set filtered artists
			Locations: utils.FetchLocations(), // Keep original locations, dates, relations
			Dates:     data.CombinedData.Dates,
			Relations: data.CombinedData.Relations,
		}
	} else {
		allData.To_displayed = data.CombinedData
	}

	// Set the Type field based on the number of members
	for i := range data.Artists {
		if len(data.Artists[i].Members) == 1 {
			data.Artists[i].Type = "Artist"
		} else {
			data.Artists[i].Type = "Group of " + strconv.Itoa(len(data.Artists[i].Members))
		}
	}

	if err := RenderTemplate(w, "index.html", http.StatusOK, allData); err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Failed to render the page.")
		return
	}
}
