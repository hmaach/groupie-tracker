package models

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Dates struct {
    Index []Date `json:"index"`
}