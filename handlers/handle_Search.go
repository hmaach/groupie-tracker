package handlers

import (
	"fmt"
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
	Key := r.FormValue("Search")
	var Newdata models.CombinedData
	for _, l := range data.Artists {
		if (strings.Contains(Key, l.Name)) || (strings.Contains(l.FirstAlbum, Key)) || (strings.Contains(Key, strconv.Itoa(l.CreationDate))) {
			if !Double(l.ID, Newdata) {
				fmt.Println(l.ID)
				Newdata.Artists = append(Newdata.Artists, l)
			}
		}
		for _, M := range l.Members {
			if strings.Contains(M, Key) && !Double(l.ID, Newdata) {
				fmt.Println(l.ID)
				Newdata.Artists = append(Newdata.Artists, l)
			}
		}
	}

	id := []int{}
	for _, j := range data.Locations.Index {
		for _, J := range j.Locations {
			if strings.Contains(J, Key) {
				if !exist(id, j.ID, Newdata) {
					id = append(id, j.ID)
				}
			}
		}
	}
	fmt.Println(id)
	for _, ids := range id {
		if !Double(ids, Newdata) {
			new, err := utils.FetchArtist(strconv.Itoa(ids))
			if err != nil {
				RenderError(w, http.StatusInternalServerError, "500 | Failed to render artist detail page.")
			}
			Newdata.Artists = append(Newdata.Artists, new)
		}
	}
	if err := RenderTemplate(w, "index.html", http.StatusOK, Newdata); err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Failed to render the page.")
		return
	}
}

func exist(ids []int, nb int, data models.CombinedData) bool {
	for _, id := range ids {
		if id == nb || Double(id, data) {
			return true
		}
	}
	return false
}

func Double(id int, data models.CombinedData) bool {
	for _, artist := range data.Artists {
		if artist.ID == id {
			return true
		}
	}
	return false
}
