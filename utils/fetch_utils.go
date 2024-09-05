package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"groupie_tracker/config"
	"groupie_tracker/data"
	"groupie_tracker/models"
)

// FetchArtist fetches the artist details, locations, relations, and dates concurrently.
func FetchAllData() (models.CombinedData, error) {
	var (
		wg  sync.WaitGroup
		mu  sync.Mutex
		err error
	)

	// Define a helper function to fetch data concurrently and handle errors.
	fetchData := func(endpoint string, dest interface{}) {
		defer wg.Done()
		if fetchErr := Fetch(endpoint, dest); fetchErr != nil {
			mu.Lock()
			err = fmt.Errorf("error fetching data from %s: %v", endpoint, fetchErr)
			mu.Unlock()
		}
	}

	// Fetch related data concurrently
	wg.Add(4)
	go fetchData("/artists", &data.Artists)
	go fetchData("/dates", &data.Dates)
	go fetchData("/locations", &data.Locations)
	go fetchData("/relation", &data.Relations)
	wg.Wait()

	// Check if any errors occurred during concurrent fetching
	if err != nil {
		return models.CombinedData{}, err
	}

	return models.CombinedData{
		Artists:   data.Artists,
		Dates:     data.Dates.Index,
		Locations: data.Locations.Index,
		Relations: data.Relations.Index,
	}, nil
}

// FetchArtist fetches the artist details, locations, relations, and dates concurrently.
func FetchArtist(id string) (models.Artist, error) {
	newid, err := strconv.Atoi(id)
	if err != nil {
		return models.Artist{}, errors.New("404 | The page you are looking for does not exist")
	}
	var artist models.Artist
	for _, v := range data.Artists {
		if v.ID == newid {
			artist.ID = v.ID
			artist.CreationDate = v.CreationDate
			artist.FirstAlbum = v.FirstAlbum
			artist.Image = v.Image
			artist.Members = v.Members
			artist.Name = v.Name
			artist.Type = v.Type
		}
	}
	var loca models.Location
	for _, loc := range data.Locations.Index {
		if loc.ID == newid {
			loca.Locations = loc.Locations
		}
	}
	var date models.Date
	for _, dat := range data.Dates.Index {
		if dat.ID == newid {
			date.Dates = dat.Dates
		}
	}
	var rel models.Relation
	for _, rela := range data.Relations.Index {
		if rela.ID == newid {
			rel.DatesLocations = rela.DatesLocations
		}
	}
	artist.Location = loca
	artist.Date = date
	artist.Relation = rel
	if artist.ID == 0 {
		return models.Artist{}, errors.New("404 | The page you are looking for does not exist")
	}
	return artist, nil
}

// Fetch fetches data from the API based on the provided endpoint and unmarshals it into the given destination.
// It takes an API endpoint as a string and a destination to unmarshal the JSON response into.
// The function returns an error if the request fails or if the API responds with a non-200 status code.
func Fetch(endpoint string, dest interface{}) error {
	resp, err := http.Get(config.ARTISTS_API_URL + endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(dest)
}
