package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestPlacesHandler_Sucess(t *testing.T) {

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

func TestPlacesHandler_CachedData(t *testing.T) {

	tests := []struct {
		name                    string
		url                     string
		wantStatusCode          int
		getPOIDataFromCacheFunc func(string) (bool, *models.Places)
		getPOIFromHereAPIFunc   func(string) (*models.Places, error)
		wantPOIData             string
	}{
		{
			name:           "should return status ok, when request is correct",
			url:            "/api/v1/places?location=London",
			wantStatusCode: 200,
			getPOIDataFromCacheFunc: func(string) (bool, *models.Places) {
				return true, &models.Places{ParkingLots: []models.PlaceInfo{{Title: "title1"}}}
			},
			getPOIFromHereAPIFunc: func(string) (*models.Places, error) {
				return &models.Places{Restaurents: []models.PlaceInfo{{Title: "title1"}}}, nil
			},
			wantPOIData: `["parkingLots": [{"title":"title1"}]]`,
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

			if reflect.DeepEqual(rr.Body.String(), tt.wantPOIData) {
				t.Errorf("place POI data: got %v want %v", rr.Body.String(), tt.wantPOIData)
			}
		})
	}
}

func Test_getPOIFromHereAPI(t *testing.T) {

	tests := []struct {
		name                     string
		locationName             string
		getLocCoordiantesFunc    func(string) (models.Position, error)
		getPOIsNearALocationFunc func(models.Position) (*models.Places, error)
		want                     *models.Places
		wantErr                  bool
	}{
		{
			name:         "should return error, when getLocCoordinateFunc returns error",
			locationName: "London",
			getLocCoordiantesFunc: func(string) (models.Position, error) {
				return models.Position{}, errors.New("unable to get coordinates")
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:         "should return error, when getPOIsNearALocationFunc returns error",
			locationName: "London",
			getLocCoordiantesFunc: func(string) (models.Position, error) {
				return models.Position{}, nil
			},
			getPOIsNearALocationFunc: func(models.Position) (*models.Places, error) {
				return nil, errors.New("unable to get coordinates")
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:         "should return places, when http client package returns proper response",
			locationName: "London",
			getLocCoordiantesFunc: func(string) (models.Position, error) {
				return models.Position{}, nil
			},
			getPOIsNearALocationFunc: func(models.Position) (*models.Places, error) {
				return &models.Places{}, nil
			},
			want:    &models.Places{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			getLocCoordiantesFunc = tt.getLocCoordiantesFunc
			getPOIsNearALocationFunc = tt.getPOIsNearALocationFunc

			got, err := getPOIFromHereAPI(tt.locationName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPOIFromHereAPI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPOIFromHereAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
