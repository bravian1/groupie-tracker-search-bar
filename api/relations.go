package api

import (
	"encoding/json"
	"net/http"

	"github.com/denilany/Groupie-Tracker/models"
)

func GetRelations(url string) (models.Relation, error) {
	resp, err := http.Get(url)
	if err != nil {
		return models.Relation{}, err
	}
	defer resp.Body.Close()

	var relations models.Relation
	err = json.NewDecoder(resp.Body).Decode(&relations)
	if err != nil {
		return relations, err
	}
	return relations, nil
}
