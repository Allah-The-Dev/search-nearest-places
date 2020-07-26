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
		return
	}
	testServer := httptest.NewServer(http.HandlerFunc(handler))
	testServerURL := testServer.URL

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
			hereGecodeAPIURL: testServerURL + "?qu=%s&apiKey=%s",
			hereAPIKey:       "",
			locationName:     "London",
			want:             models.Position{},
			wantErr:          true,
		},
		{
			name:             "it should return position, on correct request url",
			hereGecodeAPIURL: testServerURL + "?q=%s&apiKey=%s",
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
	testServer.Close()
}

func TestGetPOIsNearALocation(t *testing.T) {
	type args struct {
		poiPosition models.Position
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Places
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPOIsNearALocation(tt.args.poiPosition)
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

func Test_getNearByPlaceForACategory(t *testing.T) {
	type args struct {
		poi poiMetaData
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getNearByPlaceForACategory(tt.args.poi)
		})
	}
}
