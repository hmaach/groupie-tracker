package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"groupie_tracker/data"
	"groupie_tracker/models"
	"groupie_tracker/utils"
)

func Search(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		RenderError(w, http.StatusNotFound, "404 | The page you are looking for does not exist.")
		return
	}

	if r.Method != http.MethodGet {
		RenderError(w, http.StatusMethodNotAllowed, "405 | Method Not Allowed: Use GET")
		return
	}
	id := []int{}
	Key := r.FormValue("Search")
	Key = strings.TrimSpace(Key)
	Key = strings.ToLower(Key)
	var Newdata models.CombinedData
	for _, l := range data.Artists {
		if (strings.Contains(strings.ToLower(l.Name), Key)) || strings.ToLower(l.FirstAlbum) == Key || strconv.Itoa(l.CreationDate) == Key {
			if !exist(id, l.ID) {
				id = append(id, l.ID)
			}
		}
		for _, M := range l.Members {
			if strings.Contains(strings.ToLower(M), Key) && !exist(id, l.ID) {
				id = append(id, l.ID)
			}
		}
	}

	for _, j := range data.Locations.Index {
		for _, J := range j.Locations {
			if strings.Contains(J, Key) {
				if !exist(id, j.ID) {
					id = append(id, j.ID)
				}
			}
		}
		for _, j := range data.Dates.Index {
			for _, i := range j.Dates {
				if strings.ToLower(i) == Key {
					if !exist(id, j.ID) {
						id = append(id, j.ID)
					}
				}
			}
		}
	}
	for _, ids := range id {
		new, err := utils.FetchArtist(strconv.Itoa(ids))
		if err != nil {
			RenderError(w, http.StatusInternalServerError, "500 | Failed to render artist detail page.")
		}
		Newdata.Artists = append(Newdata.Artists, new)
	}

	type Output struct {
		To_displayed models.CombinedData
		For_search   models.CombinedData
	}

	// Create a variable of type Output and initialize it
	affiche := Output{
		To_displayed: Newdata,           // Ensure Newdata is of type models.CombinedData
		For_search:   data.CombinedData, // Ensure data.CombinedData is of type models.CombinedData
	}

	if err := RenderTemplate(w, "index.html", http.StatusOK, affiche); err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Failed to render the page.")
		return
	}
}

func exist(ids []int, nb int) bool {
	for _, id := range ids {
		if id == nb {
			return true
		}
	}
	return false
}
