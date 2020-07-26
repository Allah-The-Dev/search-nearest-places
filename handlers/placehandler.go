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

	location := query.Get("location")
	log.Printf("location name is %s", location)

	urlEncodedLocation, err := getURLEncodedLocation(location)
	if err != nil {
		log.Printf("error:: %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("url encoded location is %s", urlEncodedLocation)

	if ok, poiData := checkPOIDataInCache(location); ok {

		log.Printf("data found in cache for %s", location)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(poiData)

		log.Printf("update data in cache again, as it is the latest accessed node")
		poiDataCache.Put(location, *poiData)
		return
	}

	poiPlaces, err := getPOIFromHereAPI(urlEncodedLocation)
	if err != nil {
		log.Printf("error:: %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//update the cache
	poiDataCache.Put(location, *poiPlaces)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*poiPlaces)
}

func checkPOIDataInCache(location string) (bool, *models.Places) {
	poiData := poiDataCache.Get(location)
	if poiData != nil {
		return true, poiData
	}
	return false, nil
}

func getURLEncodedLocation(location string) (string, error) {
	u, err := url.Parse(location)
	if err != nil {
		fieldName := err.Error()
		errFmt := fmt.Errorf("url encoding error for %s ...!!; %s", location, fieldName)
		return "", errFmt
	}
	return u.EscapedPath(), nil
}

func getPOIFromHereAPI(location string) (*models.Places, error) {
	log.Printf("data not found in cache, getting from here API for location %s", location)
	locationCoordinates, err := httpclient.GetLocationCoordinates(location)
	if err != nil {
		fieldName := err.Error()
		errFmt := fmt.Errorf("unable to get location coordinate..!! %s", fieldName)
		return nil, errFmt
	}

	places, err := httpclient.GetPOINearALocation(locationCoordinates)
	if err != nil {
		fieldName := err.Error()
		errFmt := fmt.Errorf("unable to get places ...!! %s", fieldName)
		return nil, errFmt
	}
	return places, nil
}
