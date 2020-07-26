package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"search-nearest-places/cache"
	"search-nearest-places/httpclient"
	"search-nearest-places/models"
)

var poiDataCache cache.Cache

func init() {
	//initialize cache
	poiDataCache = cache.New(20)
}

//PlacesHandler ... returns places from here API
func PlacesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")

	locationName := query.Get("location")
	log.Printf("location name is %s", locationName)

	urlEncodedLocationName, err := getURLEncodedLocation(locationName)
	if err != nil {
		log.Printf("error:: %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("url encoded location is %s", urlEncodedLocationName)

	if ok, poiData := checkPOIDataInCache(locationName); ok {

		log.Printf("data found in cache for %s", locationName)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(poiData)

		log.Printf("update data in cache again, as it is the latest accessed node")
		poiDataCache.Put(locationName, *poiData)
		return
	}

	poiPlaces, err := getPOIFromHereAPI(urlEncodedLocationName)
	if err != nil {
		log.Printf("error:: %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//update the cache
	poiDataCache.Put(locationName, *poiPlaces)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*poiPlaces)
}

func checkPOIDataInCache(locationName string) (bool, *models.Places) {
	poiData := poiDataCache.Get(locationName)
	if poiData != nil {
		return true, poiData
	}
	return false, nil
}

func getURLEncodedLocation(locationName string) (string, error) {
	u, err := url.Parse(locationName)
	if err != nil {
		fieldName := err.Error()
		errFmt := fmt.Errorf("url encoding error for %s ...!!; %s", locationName, fieldName)
		return "", errFmt
	}
	return u.EscapedPath(), nil
}

func getPOIFromHereAPI(locationName string) (*models.Places, error) {
	log.Printf("data not found in cache, getting from here API for location %s", locationName)
	poiPosition, err := httpclient.GetLocationCoordinates(locationName)
	if err != nil {
		fieldName := err.Error()
		errFmt := fmt.Errorf("unable to get location coordinate..!! %s", fieldName)
		return nil, errFmt
	}

	places, err := httpclient.GetPOIsNearALocation(poiPosition)
	if err != nil {
		fieldName := err.Error()
		errFmt := fmt.Errorf("unable to get places ...!! %s", fieldName)
		return nil, errFmt
	}
	return places, nil
}
