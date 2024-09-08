package handlers

import (
	"net/http"
	"strconv"

	"groupie_tracker/data"
)

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 | Method Not Allowed: Use GET", http.StatusMethodNotAllowed)
		return
	}

	// Parse query parameters for range filters
	// creationDateMin, _ := strconv.Atoi(r.URL.Query().Get("creationDateMin"))
	// creationDateMax, _ := strconv.Atoi(r.URL.Query().Get("creationDateMax"))
	// firstAlbumMin, _ := strconv.Atoi(r.URL.Query().Get("firstAlbumMin"))
	// firstAlbumMax, _ := strconv.Atoi(r.URL.Query().Get("firstAlbumMax"))
	// membersMin, _ := strconv.Atoi(r.URL.Query().Get("membersMin"))
	// membersMax, _ := strconv.Atoi(r.URL.Query().Get("membersMax"))

	// Fetch all data
	// data, err := utils.FetchAllData()
	// if err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, "500 | Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	// Filter data based on the provided filters
	// filteredData := utils.FilterData(data, creationDateMin, creationDateMax, firstAlbumMin, firstAlbumMax, membersMin, membersMax)

	// Marshal the filtered data into JSON
	// jsonData, err := json.Marshal(data.CombinedData)
	// if err != nil {
	// 	http.Error(w, "500 | Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	// Set response header and write JSON data
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(jsonData)

	// Set the Type field based on the number of members
	for i := range data.Artists {
		if len(data.Artists[i].Members) == 1 {
			data.Artists[i].Type = "Artist"
		} else {
			data.Artists[i].Type = "Group of " + strconv.Itoa(len(data.Artists[i].Members))
		}
	}

	if err := RenderTemplate(w, "index.html", http.StatusOK, data.Artists); err != nil {
		RenderError(w, http.StatusInternalServerError, "500 | Failed to render the page.")
		return
	}
}
