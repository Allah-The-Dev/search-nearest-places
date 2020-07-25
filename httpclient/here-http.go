package httpclient

import (
	"encoding/json"
	"fmt"

	"search-nearest-places/models"
)

const (
	hereAPIKey          = "uZ9InITCNxhCIlW--t1RDnYlSplGAMkktR2UP1D_wok"
	hereBrowseAPIURL    = "https://browse.search.hereapi.com/v1/browse?at=%s3&limit=3&categories=%s&apiKey=%s"
	hereGecodeAPIURL    = "https://geocode.search.hereapi.com/v1/geocode?q=%s&apiKey=%s"
	hereRestaurentCatID = "100-1000"
)

//GetLocationCoordinates ... gives location coordinate,
// when location name is given
func GetLocationCoordinates(locationName string) (models.LocationCoordinate, error) {

	locationCoordinate := models.LocationCoordinate{}

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

	fmt.Println("here geo code response", hereGeoCodeRes)

	return hereGeoCodeRes.Items[0].Position, nil
}

//GetPlacesAroundGivenLocaton ... gives places near by to a location
//which includes restaurent, charging station, parking lot
func GetPlacesAroundGivenLocaton(coordinates models.LocationCoordinate) (*models.Places, error) {

	placesAround := &models.Places{}
	var err error

	placesAround.Restaurent, err = getNearByPlaceForACategory(coordinates, hereRestaurentCatID)
	if err != nil {
		fmt.Println("i'm here", err)
		return nil, err
	}
	return placesAround, nil
}

func getNearByPlaceForACategory(coordinates models.LocationCoordinate, categoryID string) ([]models.PlaceInfo, error) {

	placeInfoItems := models.PlaceInfoItems{}

	coordinatesStr := fmt.Sprintf("%f,%f", coordinates.Lat, coordinates.Lng)

	url := fmt.Sprintf(hereBrowseAPIURL, coordinatesStr, categoryID, hereAPIKey)

	responseBody, err := doHTTPGet(url)
	if err != nil {
		return []models.PlaceInfo{}, err
	}
	defer responseBody.Close()
	json.NewDecoder(responseBody).Decode(&placeInfoItems)

	fmt.Println("restaurent is ", placeInfoItems)
	return placeInfoItems.Items, nil

}
