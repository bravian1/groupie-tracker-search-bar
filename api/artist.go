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
		newMembers:=[]string{}
		for i, member:= range artist.Members{
			if i !=len(artist.Members)-1{

				member+=","
				
			}
			newMembers=append(newMembers, member)
		}
		newLocations:=[]string{}
		for i, member:= range place.Locations{
			if i !=len(place.Locations)-1{

				member+=","
				
			}
			newLocations=append(newLocations, member)
		}
		final = append(final, models.NewArtist{
            ID:           artist.ID,
            Image:        artist.Image,
            Name:         artist.Name,
            Members:      newMembers,
            CreationDate: artist.CreationDate,
            FirstAlbum:   artist.FirstAlbum,
            Relations:  artist.Relations,
            Dates:    artist.Dates,
            Locations:    newLocations,
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
