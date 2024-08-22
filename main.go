package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"groupie_tracker/config"
	"groupie_tracker/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not set
	}

	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/artist/{id}", handlers.ArtistHandler)
	fmt.Println("Starting the server on : http://localhost" + config.Port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	// err := http.ListenAndServe(config.Port, nil)
	// if err != nil {
	// 	fmt.Println("500 | Internal Server Error :", err)
	// }
}
