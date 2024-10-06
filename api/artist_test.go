package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/denilany/Groupie-Tracker/models"
)

var mockArtists = []models.Artist{
	{
		ID:           0,
		Image:        "image1.png",
		Name:         "Migos",
		Members:      []string{"Quavo", "Takeoff", "Offset"},
		CreationDate: 2008,
		FirstAlbum:   "Yung Rich Nation",
	},
	{
		ID:           1,
		Image:        "image2.png",
		Name:         "Destiny's Child",
		Members:      []string{"Beyonce Knowles", "Kelly Rowland", "Michelle Williams", "Farrah Franklin", "LeToya Luckett", "LaTavia Roberson"},
		CreationDate: 1990,
		FirstAlbum:   "Destiny's Child",
	},
}

func TestGetArtists_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockArtists)
	}))
	defer ts.Close()

	oldArtistAPI := artistAPI
	artistAPI = ts.URL
	defer func() { artistAPI = oldArtistAPI }()

	artists, err := GetArtists()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(artists) != len(mockArtists) {
		t.Fatalf("expected %d artists, got %d", len(mockArtists), len(artists))
	}

	for i, artist := range artists {
		if !compareArtist(artist, mockArtists[i]) {
			t.Errorf("expected artist %v, got %v", mockArtists[i], artist)
		}
	}
}

func TestGetArtists_NetworkError(t *testing.T) {
	oldArtistAPI := artistAPI
	artistAPI = "http://invalidURL"
	defer func() { artistAPI = oldArtistAPI }()

	_, err := GetArtists()
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	expectedError := "failed to get the artist URL: "
	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf("expected error message to contain %q, got %q", expectedError, err.Error())
	}
}

func TestGetArtists_NonOKStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	oldArtistAPI := artistAPI
	artistAPI = ts.URL
	defer func() { artistAPI = oldArtistAPI }()

	_, err := GetArtists()
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	expectedError := "failed to fetch artists: "
	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf("expected error message to contain %q, got %q", expectedError, err.Error())
	}
}

func compareArtist(a, b models.Artist) bool {
	if a.ID != b.ID ||
		a.Image != b.Image ||
		a.Name != b.Name ||
		a.CreationDate != b.CreationDate ||
		a.FirstAlbum != b.FirstAlbum ||
		len(a.Members) != len(b.Members) {
		return false
	}

	for i := range a.Members {
		if a.Members[i] != b.Members[i] {
			return false
		}
	}

	return true
}

// func TestGetArtistByID_Success(t *testing.T) {
// 	mockArtists := []models.Artist{
// 		{ID: 1, Name: "Artist 1"},
// 		{ID: 2, Name: "Artist 2"},
// 	}

// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(mockArtists)
// 	}))
// 	defer ts.Close()

// 	oldArtistAPI := artistAPI
// 	artistAPI = ts.URL
// 	defer func() { artistAPI = oldArtistAPI }()

// 	artist, err := GetArtistByID(1)
// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}

// 	if artist == nil {
// 		t.Fatal("expected artist, got nil")
// 	}

// 	if artist.ID != 1 || artist.Name != "Artist 1" {
// 		t.Errorf("expected Artist 1, got %+v", artist)
// 	}
// }

// func TestGetArtistByID_NotFound(t *testing.T) {
// 	mockArtists := []models.Artist{
// 		{ID: 1, Name: "Artist 1"},
// 		{ID: 2, Name: "Artist 2"},
// 	}

// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(mockArtists)
// 	}))
// 	defer ts.Close()

// 	oldArtistAPI := artistAPI
// 	artistAPI = ts.URL
// 	defer func() { artistAPI = oldArtistAPI }()

// 	artist, err := GetArtistByID(3)
// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}

// 	if artist != nil {
// 		t.Errorf("expected nil artist, got %+v", artist)
// 	}

// 	if err.Error() != "artist not found" {
// 		t.Errorf("expected 'artist not found' error, got %v", err)
// 	}
// }

// func TestGetArtistByID_GetArtistsError(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}))
// 	defer ts.Close()

// 	oldArtistAPI := artistAPI
// 	artistAPI = ts.URL
// 	defer func() { artistAPI = oldArtistAPI }()

// 	artist, err := GetArtistByID(1)
// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}

// 	if artist != nil {
// 		t.Errorf("expected nil artist, got %+v", artist)
// 	}

// 	expectedError := "failed to fetch artists: 500 Internal Server Error"
// 	if err.Error() != expectedError {
// 		t.Errorf("expected error %q, got %q", expectedError, err.Error())
// 	}
// }
