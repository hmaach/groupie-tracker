package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"groupie_tracker/config"
	"groupie_tracker/models"
)

// FetchArtist fetches the artist details, locations, relations, and dates concurrently.
func FetchArtist(id string) (models.Artist, error) {
	var (
		artist models.Artist
		wg     sync.WaitGroup
		mu     sync.Mutex
		err    error
	)

	// Fetch artist details first
	if err := Fetch("/artists/"+id, &artist); err != nil {
		return models.Artist{}, err
	}

	// Handle case where artist is not found
	if artist.ID == 0 {
		return models.Artist{}, fmt.Errorf("404")
	}

	// Define a helper function to fetch data concurrently and handle errors.
	fetchData := func(endpoint string, dest interface{}) {
		defer wg.Done()
		if fetchErr := Fetch(endpoint, dest); fetchErr != nil {
			mu.Lock()
			err = fetchErr
			mu.Unlock()
		}
	}

	// Fetch related data concurrently
	wg.Add(3)
	go fetchData("/locations/"+id, &artist.Location)
	go fetchData("/relation/"+id, &artist.Relation)
	go fetchData("/dates/"+id, &artist.Date)
	wg.Wait()

	// Check if any errors occurred during concurrent fetching
	if err != nil {
		return models.Artist{}, err
	}

	return artist, nil
}

// Fetch fetches data from the API based on the provided endpoint and unmarshals it into the given destination.
// It takes an API endpoint as a string and a destination to unmarshal the JSON response into.
// The function returns an error if the request fails or if the API responds with a non-200 status code.
func Fetch(endpoint string, dest interface{}) error {
	resp, err := http.Get(config.API_URL + endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(dest)
}
