package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetLocations_URLEmpty(t *testing.T) {
	locations, err := GetLocations("")
	if err == nil {
		t.Errorf("expected error for empty URL, got nil")
	}

	if err.Error() != "URL cannot be empty" {
		t.Errorf("expected 'URL cannot be empty' error, got %q", err.Error())
	}

	if len(locations.Locations) != 0 {
		t.Errorf("expected empty Location, got %+v", locations)
	}
}

func TestGetLocations_NetworkError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Close the connection to simulate a network error
		conn, _, _ := w.(http.Hijacker).Hijack()
		conn.Close()
	}))
	defer ts.Close()

	_, err := GetLocations(ts.URL)
	if err == nil {
		t.Fatal("expected network error, got nil")
	}

	expectedError := "failed to get URL: "
	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf("expected error message to contain %q, got %q", expectedError, err.Error())
	}
}

func TestGetLocations_MalformedJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	_, err := GetLocations(server.URL)
	if err == nil {
		t.Fatal("expected an error for malformed JSON, got nil")
	}

	expectedError := "failed to decode location response body: "
	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf("expected error message to contain %q, got %q", expectedError, err.Error())
	}
}

func TestGetLocations_NonOKStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	_, err := GetLocations(ts.URL)
	if err == nil {
		t.Fatalf("expected an error for non-200 status code, got nil")
	}

	expectedError := "failed to fetchlocations: 400 Bad Request"
	if err.Error() != expectedError {
		t.Fatalf("expected error %q, got %q", expectedError, err.Error())
	}
}
