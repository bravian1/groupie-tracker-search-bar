package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/denilany/Groupie-Tracker/models"
)

func TestGetRelations_NetworkError(t *testing.T) {
	_, err := GetRelations("http://invalid-url")
	if err == nil {
		t.Fatal("Expected network error, got nil")
	}
}

func TestGetRelations_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	_, err := GetRelations(server.URL)
	if err == nil {
		t.Fatal("Expected JSON decoding error, got nil")
	}
}

func TestGetRelations_NonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	_, err := GetRelations(server.URL)
	if err == nil {
		t.Fatal("Expected error due to non-OK status, got nil")
	}
}

func TestGetRelations_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.Relation{})
	}))
	defer server.Close()

	relations, err := GetRelations(server.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !reflect.DeepEqual(relations, models.Relation{}) {
		t.Errorf("Expected empty Relation, got %+v", relations)
	}
}
