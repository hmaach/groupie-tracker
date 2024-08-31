package handlers

import (
	"encoding/json"
	"net/http"

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
	jsonData, err := json.Marshal(data.CombinedData)
	if err != nil {
		http.Error(w, "500 | Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set response header and write JSON data
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
