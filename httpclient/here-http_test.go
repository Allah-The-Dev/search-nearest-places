package httpclient

import (
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"search-nearest-places/models"
	"testing"
)

func TestGetLocationCoordinates(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Println("*** test server invoked ****")
		//validate http method
		if r.Method != http.MethodGet {
			t.Errorf("want http method %s, got %s", http.MethodGet, r.Method)
			http.Error(w, "http method should be GET", http.StatusInternalServerError)
			return
		}
		//validate http query params
		queryParams := r.URL.Query()
		queryParamQ := queryParams.Get("q")
		queryParamAPIKey := queryParams.Get("apiKey")
		if queryParamQ == "" || queryParamAPIKey == "" {
			log.Printf("didn't get query params q = %s and apiKey = %s", queryParamQ, queryParamAPIKey)
			http.Error(w, "query param q not found", http.StatusInternalServerError)
			return
		}
		log.Printf("got query params %s and %s", queryParamQ, queryParamAPIKey)
		//when request is correct
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"items\":[{\"position\":{\"lat\":52.53041,\"lng\":13.38527}}]}"))
	}
	testServer := httptest.NewServer(http.HandlerFunc(handler))
	defer testServer.Close()

	tests := []struct {
		name             string
		hereGecodeAPIURL string
		hereAPIKey       string
		locationName     string
		want             models.Position
		wantErr          bool
	}{
		{
			name:             "it should return error, when query params are not present",
			hereGecodeAPIURL: testServer.URL + "?qu=%s&apiKey=%s",
			hereAPIKey:       "",
			locationName:     "London",
			want:             models.Position{},
			wantErr:          true,
		},
		{
			name:             "it should return position, on correct request url",
			hereGecodeAPIURL: testServer.URL + "?q=%s&apiKey=%s",
			hereAPIKey:       "hereSecret",
			locationName:     "London",
			want:             models.Position{Lat: 52.53041, Lng: 13.38527},
			wantErr:          false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			hereGecodeAPIURL = tt.hereGecodeAPIURL
			hereAPIKey = tt.hereAPIKey

			got, err := GetLocationCoordinates(tt.locationName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocationCoordinates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocationCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPOIsNearALocation(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Println("*** test server invoked ****")
		//validate http method
		if r.Method != http.MethodGet {
			t.Errorf("want http method %s, got %s", http.MethodGet, r.Method)
			http.Error(w, "http method should be GET", http.StatusInternalServerError)
			return
		}
		//validate http query params
		queryParams := r.URL.Query()
		category := queryParams.Get("categories")
		switch category {
		case categoriesOfPOI[restaurent], categoriesOfPOI[evChargingStation], categoriesOfPOI[parking]:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{\"items\":[{\"position\":{\"lat\":52.53041,\"lng\":13.38527}}]}"))
		}
	}
	testServer := httptest.NewServer(http.HandlerFunc(handler))
	defer testServer.Close()

	type args struct {
		poiPosition models.Position
	}
	tests := []struct {
		name             string
		hereBrowseAPIURL string
		position         models.Position
		want             *models.Places
		wantErr          bool
	}{
		{
			name:             "it should return data for all three category",
			hereBrowseAPIURL: testServer.URL + "?at=%s3&limit=3&categories=%s&apiKey=%s",
			position:         models.Position{Lat: 123, Lng: 235},
			want: &models.Places{
				Restaurents: []models.PlaceInfo{
					{Position: models.Position{Lat: 52.53041, Lng: 13.38527}},
				},
				EvChargingStations: []models.PlaceInfo{
					{Position: models.Position{Lat: 52.53041, Lng: 13.38527}},
				},
				ParkingLots: []models.PlaceInfo{
					{Position: models.Position{Lat: 52.53041, Lng: 13.38527}},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			hereBrowseAPIURL = tt.hereBrowseAPIURL

			got, err := GetPOIsNearALocation(tt.position)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPOIsNearALocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPOIsNearALocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
