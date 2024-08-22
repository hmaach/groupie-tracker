package models

// Artist structure for holding artist data
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     Location
	Relation     Relation
	Date         Date
}

type Location struct {
	ID            int      `json:"id"`
	Locations     []string `json:"locations"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// this data is for the index page only

type ArtistSummary struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Members []string `json:"members"`
	Type    string
}

type ArtistsPageData struct {
	Artists []ArtistSummary
}
