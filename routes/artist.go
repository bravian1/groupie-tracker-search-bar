package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/denilany/Groupie-Tracker/api"

	"github.com/denilany/Groupie-Tracker/models"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("Wrong method, %s, expected %s\n", r.Method, http.MethodGet)
		errorMsg(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		log.Printf("wrong path: %s\n", r.URL.Path)
		errorMsg(w, "Page not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("frontend/index.html")
	if err != nil {
		log.Printf("failed to parse template: %v\n", err)
		errorMsg(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	artists, err := api.GetNewArtists()
	if err != nil {
		log.Printf("failed to fetch artist: %v\n", err)
		errorMsg(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, artists)
}

func getArtistbyName(name string) (models.Artist, map[string][]string) {
	artists, err := api.GetArtists()
	if err != nil {
		log.Printf("failed to fetch artist: %v", err)
		return models.Artist{}, make(map[string][]string)
	}

	var final models.Artist
	var relations models.Relation

	for _, artist := range artists {
		if artist.Name == name {
			relations, err = api.GetRelations(artist.Relations)
			if err != nil {
				log.Printf("failed to fetch relations: %v", err)
				return models.Artist{}, make(map[string][]string)
			}
			final = artist
		}
	}

	return final, relations.DatesLocations
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorMsg(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	t, err := template.ParseFiles("frontend/profile.html")
	if err != nil {
		log.Printf("failed to parse template: %v", err)
		errorMsg(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	templ := template.Must(t, err)
	name := strings.Split(r.URL.Path, "/")[2]
	fmt.Println(name)
	artist, relations := getArtistbyName(name)

	locations, err := api.GetLocations(artist.Locations)
	if err != nil {
		log.Printf("failed to fetch locations: %v", err)
		errorMsg(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	dates, _ := api.GetDates(locations.Dates)

	if artist.Name == "" {
		errorMsg(w, "Artist not found", http.StatusNotFound)
		return
	}

	concert := make(map[string][]string)
	for key, val := range relations {
		temp := strings.Replace(key, "-", ", ", -1)
		temp = strings.Replace(temp, "_", " ", -1)
		temp = capitalize(temp)
		concert[temp] = val
	}

	final := struct {
		Name         string
		Concerts     map[string][]string
		Image        string
		Members      []string
		CreationDate int
		FirstAlbum   string
		Location     []string
		Dates        []string
		ID           int
	}{
		Name:         artist.Name,
		Concerts:     concert,
		Image:        artist.Image,
		Members:      artist.Members,
		CreationDate: artist.CreationDate,
		FirstAlbum:   artist.FirstAlbum,
		Location:     locations.Locations,
		Dates:        dates.Dates,
		ID:           artist.ID,
	}

	templ.Execute(w, final)
}

func capitalize(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		// Capitalize the first letter and concatenate it with the rest of the lowercase word
		words[i] = string(unicode.ToUpper(rune(word[0]))) + strings.ToLower(word[1:])
	}
	return strings.Join(words, " ")
}
