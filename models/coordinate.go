package models

//Position ... coordinates of a location
type Position struct {
	Name string  `json:"name,omitempty"`
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
}

//HereGeoCodeItem ... one item of geo code api
type HereGeoCodeItem struct {
	Position Position `json:"position"`
}

//HereGeoCodeResponse ... here geo code reponse part
type HereGeoCodeResponse struct {
	Items []HereGeoCodeItem `json:"items"`
}
