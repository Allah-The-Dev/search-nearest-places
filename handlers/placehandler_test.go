package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"search-nearest-places/models"
	"testing"
)

func TestPlacesHandler_BadRequest(t *testing.T) {

	tests := []struct {
		name                  string
		url                   string
		wantStatusCode        int
		getPOIFromHereAPIFunc func(string) (*models.Places, error)
	}{
		{
			name:           "should return bad request, when query param location is not present",
			url:            "/api/v1/places",
			wantStatusCode: 400,
			getPOIFromHereAPIFunc: func(string) (*models.Places, error) {
				return &models.Places{}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			getPOIFromHereAPIFunc = tt.getPOIFromHereAPIFunc

			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(PlacesHandler)

			handler.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		})
	}
}

func TestPlacesHandler_InernalServerError(t *testing.T) {

	tests := []struct {
		name                  string
		url                   string
		wantStatusCode        int
		getPOIFromHereAPIFunc func(string) (*models.Places, error)
	}{
		{
			name:           "should return internal server error, when here api response fails",
			url:            "/api/v1/places?location=London",
			wantStatusCode: 500,
			getPOIFromHereAPIFunc: func(string) (*models.Places, error) {
				return &models.Places{}, errors.New("unable to get data from here api")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			getPOIFromHereAPIFunc = tt.getPOIFromHereAPIFunc

			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(PlacesHandler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusInternalServerError {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

		})
	}
}

func TestPlacesHandler(t *testing.T) {

	tests := []struct {
		name                  string
		url                   string
		wantStatusCode        int
		getPOIFromHereAPIFunc func(string) (*models.Places, error)
	}{
		{
			name:           "should return status ok, when request is correct",
			url:            "/api/v1/places?location=London",
			wantStatusCode: 200,
			getPOIFromHereAPIFunc: func(string) (*models.Places, error) {
				return &models.Places{}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			getPOIFromHereAPIFunc = tt.getPOIFromHereAPIFunc

			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(PlacesHandler)

			handler.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		})
	}
}
