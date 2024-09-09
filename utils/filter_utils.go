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
