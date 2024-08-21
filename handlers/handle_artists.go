package groupie_tracker

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Data struct to hold data for error
type ErrorData struct {
	Error   string
	Code    int
	Message string
}

// Artist structure for holding artist data
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type ArtistsPageData struct {
	Artists []Artist
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is for the root URL
	if r.URL.Path != "/" {
		RenderTemplate(w, "error.html", ErrorData{
			Error:   "404 Not Found",
			Code:    http.StatusNotFound,
			Message: "The page you are looking for does not exist.",
		})
		return
	}

	data, err := getArtists()
	if err != nil {
		RenderTemplate(w, "error.html", ErrorData{
			Error:   "500 Internal Server Error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to render the page.",
		})
		return
	}
	// Pass the data to the template
	err = RenderTemplate(w, "index.html", data)
	if err != nil {
		RenderTemplate(w, "error.html", ErrorData{
			Error:   "500 Internal Server Error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to render the page.",
		})
	}
}

func ArtistDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Extract artist ID from URL
	id := r.URL.Query().Get("id")

	artist, err := getArtistByID(id)
	if err != nil {
		RenderTemplate(w, "error.html", ErrorData{
			Error:   "404 Not Found",
			Code:    http.StatusNotFound,
			Message: "Artist not found.",
		})
		return
	}

	// Render artist detail page
	err = RenderTemplate(w, "artist.html", artist)
	if err != nil {
		RenderTemplate(w, "error.html", ErrorData{
			Error:   "500 Internal Server Error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to render artist detail page.",
		})
	}
}

func getArtists() (ArtistsPageData, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return ArtistsPageData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ArtistsPageData{}, errors.New("API returned non-200 status code")
	}

	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return ArtistsPageData{}, err
	}

	return ArtistsPageData{Artists: artists}, nil
}

func getArtistByID(id string) (Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		return Artist{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Artist{}, errors.New("artist not found")
	}

	var artist Artist
	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		return Artist{}, err
	}

	return artist, nil
}
