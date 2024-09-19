package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/denilany/Groupie-Tracker/models"
)

func GetDates(url string) (models.Date, error) {
	resp, err := http.Get(url)
	if err != nil {
		return models.Date{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Date{}, fmt.Errorf("failed to fetch dates")
	}

	var dates models.Date
	err = json.NewDecoder(resp.Body).Decode(&dates)
	if err != nil {
		return models.Date{}, err
	}
	return dates, nil
}
