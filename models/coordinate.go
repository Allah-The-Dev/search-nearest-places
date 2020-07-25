package models

//LocationCoordinate ... coordinates of a location
type LocationCoordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

//HereGeoCodeItem ... one item of geo code api
type HereGeoCodeItem struct {
	Position LocationCoordinate `json:"position"`
}

//HereGeoCodeResponse ... here geo code reponse part
type HereGeoCodeResponse struct {
	Items []HereGeoCodeItem `json:"items"`
}

//Location ... location info
type Location struct {
	name string
	long float64
	lat  float64
}
