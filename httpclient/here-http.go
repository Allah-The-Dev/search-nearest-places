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

//LocationCoordinate ... coordinates of a location
type LocationCoordinate struct {
	lat float64
	lng float64
}

type hereGeoCodeItem struct {
	position LocationCoordinate
}

type hereGeoCodeResponse struct {
	items []hereGeoCodeItem
}

//Places ... represent place info nearby
type Places struct {
	Restaurent models.Restaurent `json:"restaurent"`
}

type location struct {
	name string
	long float64
	lat  float64
}

//GetLocationCoordinates ... gives location coordinate,
// when location name is given
func GetLocationCoordinates(locationName string) (*LocationCoordinate, error) {
	url := fmt.Sprintf(hereGecodeAPIURL, locationName, hereAPIKey)
	responseBody, err := doHTTPGet(url)
	if err != nil {
		return nil, err
	}
	var hereGeoCodeRes hereGeoCodeResponse
	json.NewDecoder(*responseBody).Decode(&hereGeoCodeRes)

	coordinates := &LocationCoordinate{
		lat: hereGeoCodeRes.items[0].position.lat,
		lng: hereGeoCodeRes.items[0].position.lng,
	}

	return coordinates, nil
}

//GetPlacesAroundGivenLocaton ... gives places near by to a location
//which includes restaurent, charging station, parking lot
func GetPlacesAroundGivenLocaton(coordinates *LocationCoordinate) (*Places, error) {

	var placesAround *Places
	var err error

	placesAround.Restaurent, err = getNearByRestaurents(coordinates, hereRestaurentCatID)
	if err != nil {
		return nil, err
	}
	return placesAround, nil
}

func getNearByRestaurents(coordinates *LocationCoordinate, categoryID string) (models.Restaurent, error) {

	restaurent := models.Restaurent{}

	coordinatesStr := fmt.Sprintf("%f,%f", coordinates.lat, coordinates.lng)

	url := fmt.Sprintf(hereBrowseAPIURL, coordinatesStr, categoryID, hereAPIKey)

	responseBody, err := doHTTPGet(url)
	if err != nil {
		return restaurent, err
	}

	json.NewDecoder(*responseBody).Decode(&restaurent)

	return restaurent, nil

}
