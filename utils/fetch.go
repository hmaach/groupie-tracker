package utils


import (
	"encoding/json"
	"fmt"
	"net/http"
  
	"groupie_tracker/config"
  )
  
  // Fetch fetches data from the API based on the provided endpoint and unmarshals it into the given destination.
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