package utils

import (
	"fmt"
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
) ([]models.Artist, []models.Location) {
	var (
		filteredArtists   []models.Artist
		filteredLocations []models.Location
	)
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

		if !Exist(id, artist.ID) {
			id = append(id, artist.ID)
		}
	}
	// Filter locations
	if ExistData(location) {
		for _, places := range data.Locations {
			for _, place := range places.Locations {
				if strings.Contains(strings.ToLower(place), strings.ToLower(location)) {
					fmt.Println("location:", location)
					artist, err := FetchArtist(strconv.Itoa(places.ID))
					if err != nil {
						fmt.Println("Error fetching artist:", err)
						continue
					}
					if !Exist(id, artist.ID) {
						id = append(id, artist.ID)
					}
					continue
				}
			}
		}
	}
	fmt.Println(filteredArtists)
	return filteredArtists, filteredLocations
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
