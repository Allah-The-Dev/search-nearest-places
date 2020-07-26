package httpclient

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"search-nearest-places/models"
)

const (
	hereAPIKey        = "uZ9InITCNxhCIlW--t1RDnYlSplGAMkktR2UP1D_wok"
	hereBrowseAPIURL  = "https://browse.search.hereapi.com/v1/browse?at=%s3&limit=3&categories=%s&apiKey=%s"
	hereGecodeAPIURL  = "https://geocode.search.hereapi.com/v1/geocode?q=%s&apiKey=%s"
	restaurent        = "restaurent"
	evChargingStation = "evChargingStation"
	parking           = "parking"
)

type poiMetaData struct {
	categoryID   string
	categoryName string
	coordinates  models.Position
	dataChannel  chan models.PlaceInfoItems
	waitGroup    *sync.WaitGroup
}

var categoriesOfPOI map[string]string

func init() {
	categoriesOfPOI = map[string]string{
		restaurent:        "100-1000",
		evChargingStation: "700-7600-0322",
		parking:           "800-8500",
	}
}

//GetLocationCoordinates ... gives location coordinate,
// when location name is given
func GetLocationCoordinates(locationName string) (models.Position, error) {

	locationCoordinate := models.Position{}

	url := fmt.Sprintf(hereGecodeAPIURL, locationName, hereAPIKey)

	responseBody, err := doHTTPGet(url)
	if err != nil {
		return locationCoordinate, err
	}
	defer responseBody.Close()

	var hereGeoCodeRes models.HereGeoCodeResponse
	err = json.NewDecoder(responseBody).Decode(&hereGeoCodeRes)
	if err != nil {
		return locationCoordinate, err
	}

	log.Println("here geo code response received")

	return hereGeoCodeRes.Items[0].Position, nil
}

//GetPOIsNearALocation ... gives places near by to a location
//which includes restaurent, charging station, parking lot
func GetPOIsNearALocation(poiPosition models.Position) (*models.Places, error) {

	placesAround := &models.Places{}

	poiDataChannel := make(chan models.PlaceInfoItems)

	var wg sync.WaitGroup

	for poiCategoryName, poiCategoryID := range categoriesOfPOI {
		wg.Add(1) // This tells the waitgroup, that there is now 1 pending operation here
		poiMetaData := poiMetaData{
			categoryID:   poiCategoryID,
			categoryName: poiCategoryName,
			coordinates:  poiPosition,
			dataChannel:  poiDataChannel,
			waitGroup:    &wg,
		}
		go getNearByPlaceForACategory(poiMetaData)
	}

	go func() {
		wg.Wait()
		close(poiDataChannel)
	}() // This calls itself

	for poiData := range poiDataChannel {
		if poiData.Err != nil {
			return nil, poiData.Err
		}
		switch poiData.POIName {
		case restaurent:
			placesAround.Restaurents = poiData.Items
		case evChargingStation:
			placesAround.EvChargingStations = poiData.Items
		case parking:
			placesAround.ParkingLots = poiData.Items
		}
	}
	return placesAround, nil
}

func getNearByPlaceForACategory(poi poiMetaData) {

	defer poi.waitGroup.Done()

	placeInfoItems := models.PlaceInfoItems{}

	coordinatesStr := fmt.Sprintf("%f,%f", poi.coordinates.Lat, poi.coordinates.Lng)

	url := fmt.Sprintf(hereBrowseAPIURL, coordinatesStr, poi.categoryID, hereAPIKey)

	responseBody, err := doHTTPGet(url)
	if err != nil {
		fieldName := err.Error()
		msg := fmt.Errorf("unable to get %s ; %s", poi.categoryName, fieldName)
		poi.dataChannel <- models.PlaceInfoItems{
			POIName: poi.categoryName,
			Items:   []models.PlaceInfo{},
			Err:     msg,
		}
		return
	}
	defer responseBody.Close()
	json.NewDecoder(responseBody).Decode(&placeInfoItems)

	poi.dataChannel <- models.PlaceInfoItems{
		POIName: poi.categoryName,
		Items:   placeInfoItems.Items,
		Err:     nil,
	}
}
