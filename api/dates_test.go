package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/denilany/Groupie-Tracker/models"
)

func compareDates(a, b models.Date) bool {
	if len(a.Dates) != len(b.Dates) {
		return false
	}
	return true
}

func TestGetDates_Success(t *testing.T) {
	mockDates := models.Date{
		ID:    0,
		Dates: []string{"*23-08-2019", "*20-08-2019", "*28-01-2020", "*07-02-2020"},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockDates)
	}))
	defer ts.Close()

	dates, err := GetDates(ts.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !compareDates(dates, mockDates) {
		t.Fatalf("expected dates %+v, got %+v", mockDates, dates)
	}
}

func TestGetDates_NetworkError(t *testing.T) {
	invalidURL := "http://invalid-URL"

	_, err := GetDates(invalidURL)
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	expectedError := "Get \"http://invalid-URL\": dial tcp: lookup invalid-URL on 127.0.0.53:53: server misbehaving"
	if err.Error() != expectedError {
		t.Fatalf("expected error to contain %q, got %q", expectedError, err.Error())
	}
}

func TestGetDates_EmptyResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte{})
	}))
	defer ts.Close()

	_, err := GetDates(ts.URL)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	expectedError := "EOF"
	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf("expected error to contain %q, got %q", expectedError, err.Error())
	}
}
