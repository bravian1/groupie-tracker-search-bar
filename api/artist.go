package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/denilany/Groupie-Tracker/models"
)

var artistAPI = "https://groupietrackers.herokuapp.com/api/artists"

func GetArtists() ([]models.Artist, error) {
	resp, err := http.Get(artistAPI)
	if err != nil {
		return nil, fmt.Errorf("failed to get the artist URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch artists: %s", resp.Status)
	}

	var artists []models.Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, fmt.Errorf("failed to decode artist response: %v", err)
	}

	// fmt.Println("artists: ",artists)
	return artists, nil
}

func GetNewArtists() ([]models.NewArtist, error) {
	fmt.Println("here")
	artists,err:=GetArtists()
	if err!=nil{
        return nil,err
    }
	var final []models.NewArtist
	for _, artist := range artists {
		place, err:=GetLocations(artist.Locations)
		if err!=nil{
			return nil,err
		}
		final = append(final, models.NewArtist{
            ID:           artist.ID,
            Image:        artist.Image,
            Name:         artist.Name,
            Members:      artist.Members,
            CreationDate: artist.CreationDate,
            FirstAlbum:   artist.FirstAlbum,
            Relations:  artist.Relations,
            Dates:    artist.Dates,
            Locations:    place.Locations,
        })
        
	}
	fmt.Println("final: ",final)
	return final, nil
}

func GetArtistByID(artistID int) (*models.Artist, error) {
	artists, err := GetArtists()
	if err != nil {
		return nil, err
	}

	for _, artist := range artists {
		if artist.ID == artistID {
			return &artist, nil
		}
	}
	return nil, errors.New("artist not found")
}
