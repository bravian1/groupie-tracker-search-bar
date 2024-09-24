package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/denilany/Groupie-Tracker/routes"
)

var Port = ":8000"

func main() {
	http.HandleFunc("/", routes.ArtistHandler)
	http.HandleFunc("/artist/", routes.ProfileHandler)
	http.HandleFunc("/static/", routes.Static)

	fmt.Printf("Server started on http://localhost%s\n", Port)
	err := http.ListenAndServe(Port, nil)
	if err != nil {
		log.Printf("failed to listen to the port %s", Port)
	}
}
