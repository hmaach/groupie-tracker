package models

type GeocodeResponse []struct {
	Lat         string `json:"lat"`
	Lng         string `json:"lon"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

// Coordinates for geocoded locations
type Coordinates struct {
	Lat          float64
	Lng          float64
	Name         string
	LocationName string
}

type CoordinatesOfArtist struct {
	Coordinates []Coordinates
}
