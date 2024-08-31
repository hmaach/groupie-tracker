package utils

import (
	"strconv"
	"strings"

	"groupie_tracker/models"
)

func FilterData(data models.CombinedData,
	creationDateMin,
	creationDateMax,
	firstAlbumMin,
	firstAlbumMax,
	membersMin,
	membersMax int,
) []models.Artist {
	var filteredArtists []models.Artist

	for _, artist := range data.Artists {
		// Filter by creation date
		if artist.CreationDate < creationDateMin || artist.CreationDate > creationDateMax {
			continue
		}

		// Filter by first album date
		firstAlbumYear, err := strconv.Atoi(strings.Split(artist.FirstAlbum, "-")[0])
		if err != nil || firstAlbumYear < firstAlbumMin || firstAlbumYear > firstAlbumMax {
			continue
		}

		// Filter by number of members
		if len(artist.Members) < membersMin || len(artist.Members) > membersMax {
			continue
		}

		filteredArtists = append(filteredArtists, artist)
	}

	return filteredArtists
}
