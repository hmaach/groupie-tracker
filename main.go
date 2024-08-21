package main

import (
	"fmt"
	"net/http"

	"groupie_tracker/config"
	"groupie_tracker/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/artist/{id}", handlers.ArtistHandler)
	fmt.Println("Starting the server on : http://localhost:8080")
	err := http.ListenAndServe(config.Port, nil)
	if err != nil {
		fmt.Println("500 | Internal Server Error :", err)
	}
}
