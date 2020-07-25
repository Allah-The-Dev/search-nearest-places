package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"search-nearest-places/cache"
	"search-nearest-places/httpclient"
	"search-nearest-places/models"
)

var poiDataCache *cache.LRUCache

func init() {
	//initialize cache
	poiDataCache = cache.New(20)
}

//PlacesHandler ... returns places from here API
func PlacesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")

	location := query.Get("location")
	fmt.Println("location name is ", location)

	if isAvailableInCache, poiData := checkPOIDataInCache(location); isAvailableInCache {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(poiData)
		return
	}

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

	//update the cache
	poiDataCache.Put(location, *places)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*places)
}

func checkPOIDataInCache(location string) (bool, *models.Places) {
	poiData := poiDataCache.Get(location)
	if poiData != nil {
		return true, poiData
	}
	return false, nil
}
