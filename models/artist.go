package models

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

type ArtistsPageData struct {
	Artists []Artist
}

// CombinedData structure to hold all fetched data.
type CombinedData struct {
	Artists   []Artist
	Locations []Location
	Dates     []Date
	Relations []Relation
}

type Output struct {
	To_displayed CombinedData
	For_search   CombinedData
}
