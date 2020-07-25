package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"search-nearest-places/httpclient"
)

//PlacesHandler ... returns places from here API
func PlacesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")

	location := query.Get("location")
	fmt.Println("location name is ", location)

	locationCoordinates, err := httpclient.GetLocationCoordinates(location)
	if err != nil {
		fieldName := err.Error()
		msg := fmt.Sprintf("unable to get location coordinate %s", fieldName)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	places, err := httpclient.GetPOINearALocation(locationCoordinates)
	if err != nil {
		fieldName := err.Error()
		msg := fmt.Sprintf("unable to get places %s", fieldName)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Printf("%v", places)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*places)
}
