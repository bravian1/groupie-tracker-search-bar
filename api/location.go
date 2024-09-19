package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/denilany/Groupie-Tracker/models"
)

func GetLocations(url string) (models.Location, error) {
	if url == "" {
		return models.Location{}, errors.New("URL cannot be empty")
	}
	resp, err := http.Get(url)
	if err != nil {
		return models.Location{}, fmt.Errorf("failed to get URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Location{}, fmt.Errorf("failed to fetchlocations: %s", resp.Status)
	}

	var locations models.Location
	err = json.NewDecoder(resp.Body).Decode(&locations)
	if err != nil {
		return models.Location{}, fmt.Errorf("failed to decode location response body: %v", err)
	}

	return locations, nil
}
