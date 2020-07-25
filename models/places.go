package models

//Places ... represent place info nearby
type Places struct {
	Restaurents        []PlaceInfo `json:"restaurents"`
	EvChargingStations []PlaceInfo `json:"evChargingStations"`
	ParkingLots        []PlaceInfo `json:"parkingLots"`
}

//PlaceInfoItems ... items from here API
type PlaceInfoItems struct {
	POIName string
	Items   []PlaceInfo `json:"items"`
	Err     error
}

//PlaceInfo ... higher level restaurent info
type PlaceInfo struct {
	Title        string         `json:"title"`
	ID           string         `json:"id"`
	ResultType   string         `json:"resultType"`
	Address      Address        `json:"address"`
	Position     Position       `json:"position"`
	Access       []Access       `json:"access"`
	Distance     int            `json:"distance"`
	Categories   []Categories   `json:"categories"`
	References   []References   `json:"references"`
	FoodTypes    []FoodTypes    `json:"foodTypes"`
	Contacts     []Contacts     `json:"contacts"`
	OpeningHours []OpeningHours `json:"openingHours"`
}

// Address ... restaurent address
type Address struct {
	Label       string `json:"label"`
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`
	State       string `json:"state"`
	County      string `json:"county"`
	City        string `json:"city"`
	District    string `json:"district"`
	Street      string `json:"street"`
	PostalCode  string `json:"postalCode"`
	HouseNumber string `json:"houseNumber"`
}

//Position ... position of restaurent
type Position struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

//Access ... access coordinates of restaurent
type Access struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

//Categories ... restaurent categories
type Categories struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Primary bool   `json:"primary"`
}

//Supplier ... restaurent supplier
type Supplier struct {
	ID string `json:"id"`
}

//References ... suppliers who refer this restaurent
type References struct {
	Supplier Supplier `json:"supplier"`
	ID       string   `json:"id"`
}

//FoodTypes ... food types offered by restaurent
type FoodTypes struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Primary bool   `json:"primary"`
}

//Phone ... phone no of place
type Phone struct {
	Value string `json:"value"`
}

//Www ... web domain of a place
type Www struct {
	Value string `json:"value"`
}

//Contacts ... contact info of a place
type Contacts struct {
	Phone []Phone `json:"phone"`
	Www   []Www   `json:"www"`
}

//Structured ... opening hours info
type Structured struct {
	Start      string `json:"start"`
	Duration   string `json:"duration"`
	Recurrence string `json:"recurrence"`
}

//OpeningHours ... place opening hour
type OpeningHours struct {
	Text       []string     `json:"text"`
	IsOpen     bool         `json:"isOpen"`
	Structured []Structured `json:"structured"`
}
