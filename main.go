package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie_tracker/config"
	"groupie_tracker/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/artist/{id}", handlers.ArtistHandler)
	fmt.Println("Starting the server on : http://localhost" + config.Port)
	log.Fatal(http.ListenAndServe(config.Port, nil))
}
