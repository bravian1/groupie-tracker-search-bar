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
	return artists, nil
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
