package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie_tracker/config"
	"groupie_tracker/data"
	"groupie_tracker/handlers"
	"groupie_tracker/utils"
)

func init() {
	// Initialize all data before starting the server
	var err error
	data.CombinedData, err = utils.FetchAllData()
	if err != nil {
		fmt.Printf("Failed to fetch data: %v", err)
	}
}

func main() {
	// Handle requests for assets using the custom handler
	http.HandleFunc("/assets/", handlers.AssetsHandler)

	// Parse all HTML templates before starting the server
	handlers.ParseTemplates()

	// Route handlers
	http.HandleFunc("/", handlers.MainHandler)              // Root route (home page)
	http.HandleFunc("/artist/{id}", handlers.ArtistHandler) // Artist detail page
	http.HandleFunc("/search", handlers.Search)
	// Start the server
	serverPort := config.Port
	fmt.Println("Starting the server on http://localhost" + serverPort)
	log.Println(http.ListenAndServe(serverPort, nil))
}
