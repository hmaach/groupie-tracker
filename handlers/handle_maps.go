package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"groupie_tracker/models"
	"groupie_tracker/utils"
)

func GeocodeLocations(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var coordinatesOfArtist models.CoordinatesOfArtist
	var locations models.Location

	// Fetch artist locations
	if err := utils.Fetch("/locations/"+id, &locations); err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Status Internal Server Error")
		return
	}

	for _, location := range locations.Locations {
		var coordinates models.Coordinates
		var err error

		// Attempt to fetch coordinates using the full location
		coordinates, err = utils.Geocode(location)
		if err != nil {
			// If an error occurs, trim the location to remove everything after the last hyphen
			if lastHyphen := strings.LastIndex(location, "-"); lastHyphen != -1 {
				location = location[:lastHyphen]
				// Try fetching coordinates again with the trimmed location
				coordinates, err = utils.Geocode(location)
			}
		}

		// If both attempts fail, continue to the next location
		if err != nil {
			fmt.Println("no data found for the given location:", location)
			continue
		}

		// Append the successfully geocoded location to the coordinatesOfArtist
		coordinatesOfArtist.Coordinates = append(coordinatesOfArtist.Coordinates, coordinates)
	}

	// Encode the coordinatesOfArtist to JSON and write to the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(coordinatesOfArtist); err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Status Internal Server Error")
		return
	}
}
