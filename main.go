package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie_tracker/config"
	"groupie_tracker/handlers"
)

func main() {
	// Serve static files from the "./assets" directory
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Parse all HTML templates before starting the server
	handlers.ParseTemplates()

	// Route handlers
	http.HandleFunc("/", handlers.MainHandler)                             // Root route (home page)
	http.HandleFunc("/artist/{id}", handlers.ArtistHandler)                // Artist detail page

	// Start the server
	serverPort := config.Port
	fmt.Println("Starting the server on http://localhost" + serverPort)
	log.Println(http.ListenAndServe(serverPort, nil))
}