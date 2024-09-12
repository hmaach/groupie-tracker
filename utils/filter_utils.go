package utils

import (
	"strconv"
	"strings"

	"groupie_tracker/models"
)

// FilterData applies filters to the artist data
func FilterData(
	data models.CombinedData,
	creationDateMin, creationDateMax,
	firstAlbumMin, firstAlbumMax int,
	location string,
	members []int,
) []models.Artist {
	var filteredArtists []models.Artist
	id := []int{}
	for _, artist := range data.Artists {
		// Filter by creation date, only if the range is valid
		if creationDateMax >= creationDateMin {
			if artist.CreationDate < creationDateMin || artist.CreationDate > creationDateMax {
				continue
			}
		}

		// Filter by first album date, only if the range is valid
		firstAlbumYear, err := strconv.Atoi(strings.Split(artist.FirstAlbum, "-")[2])
		if err == nil && firstAlbumMax >= firstAlbumMin {
			if firstAlbumYear < firstAlbumMin || firstAlbumYear > firstAlbumMax {
				continue
			}
		}
		// Filter by number of members
		if len(members) > 0 && !intInSlice(len(artist.Members), members) {
			continue
		}

		if !Exist(id, artist.ID) && Checklocation(artist.ID, location, id, data.Locations) {
			id = append(id, artist.ID)
		} else {
			continue
		}

	}
	// Filter locations
	for _, id2 := range id {
		artist, _ := FetchArtist(strconv.Itoa(id2))
		filteredArtists = append(filteredArtists, artist)
	}
	// fmt.Println(filteredArtists)
	return filteredArtists
}

// intInSlice checks if an integer is in a slice
func intInSlice(value int, list []int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func Exist(ids []int, nb int) bool {
	for _, id := range ids {
		if id == nb {
			return true
		}
	}
	return false
}

func Checklocation(id1 int, location string, ids []int, locations []models.Location) bool {
	if location == "" {
		return true
	}
	for _, artist := range locations {
		// for _, v := range artist.ID {
		if artist.ID == id1 {
			for _, place := range artist.Locations {
				if strings.Contains(strings.ToLower(place), strings.ToLower(location)) {
					if !Exist(ids, artist.ID) {
						return true
					}
					continue
				}
			}
		}
		// }
	}
	return false
}

// if artist.ID == id1 {

// }
