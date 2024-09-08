package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"groupie_tracker/config"
	"groupie_tracker/models"
)

func Geocode(location string) (models.Coordinates, error) {
	var coordinates models.Coordinates
	u, _ := url.Parse(config.MAP_API_URL)
	q := u.Query()
	q.Set("q", location)
	q.Set("format", "jsonv2")
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return coordinates, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return coordinates, err
	}

	var geocodeResponse models.GeocodeResponse
	err = json.Unmarshal(body, &geocodeResponse)
	if err != nil {
		return coordinates, err
	}

	if len(geocodeResponse) == 0 {
		return coordinates, fmt.Errorf("no results found for location: %s", location)
	}

	// Get the first result
	result := geocodeResponse[0]
	coordinates.Lat, err = strconv.ParseFloat(result.Lat, 64)
	if err != nil {
		return coordinates, fmt.Errorf("failed to parse latitude: %v", err)
	}
	coordinates.Lng, err = strconv.ParseFloat(result.Lng, 64)
	if err != nil {
		return coordinates, fmt.Errorf("failed to parse longitude: %v", err)
	}

	coordinates.Name = result.Name
	coordinates.LocationName = result.DisplayName

	return coordinates, nil
}
