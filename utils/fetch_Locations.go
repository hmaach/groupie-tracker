package utils

import (
	"strings"

	"groupie_tracker/data"
	"groupie_tracker/models"
)

func FetchLocations() []models.Location {
	var arr []string
	for _, v := range data.Locations.Index {
		for _, place := range v.Locations {
			place = strings.ReplaceAll(place, "-", ", ")
			if !CheckRepeatLocs(place, arr) {
				arr = append(arr, place)
			} else {
				continue
			}
		}
	}
	// fmt.Println("arr", arr)
	loca := models.Location{
		ID:        0,
		Locations: arr,
	}
	var loats []models.Location
	loats = append(loats, loca)
	return loats
}

func CheckRepeatLocs(s string, arr []string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == s {
			return true
		}
	}
	return false
}
