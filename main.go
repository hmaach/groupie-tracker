package main

import (
	"fmt"
	"net/http"

	web "groupie_tracker/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", web.MainHandler)
	http.HandleFunc("/artist", web.ArtistDetailHandler)
	fmt.Println("Starting the server on : http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("500 | Internal Server Error :", err)
	}
}
