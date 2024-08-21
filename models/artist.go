package models

// Artist structure for holding artist data
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Type         string
	Location     Location
	Relation     Relation
	Date         Date
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Related_Dates     []string 
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}


type ArtistsPageData struct {
	Artists []Artist
}
