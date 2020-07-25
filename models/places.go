package models

//Places ... represent place info nearby
type Places struct {
	Restaurents        []PlaceInfo `json:"restaurents,omitempty"`
	EvChargingStations []PlaceInfo `json:"evChargingStations,omitempty"`
	ParkingLots        []PlaceInfo `json:"parkingLots,omitempty"`
}

//PlaceInfoItems ... items from here API
type PlaceInfoItems struct {
	POIName string
	Items   []PlaceInfo `json:"items,omitempty"`
	Err     error
}

//PlaceInfo ... higher level restaurent info
type PlaceInfo struct {
	Title        string         `json:"title,omitempty"`
	ID           string         `json:"id,omitempty"`
	ResultType   string         `json:"resultType,omitempty"`
	Address      Address        `json:"address,omitempty"`
	Position     Position       `json:"position,omitempty"`
	Access       []Access       `json:"access,omitempty"`
	Distance     int            `json:"distance,omitempty"`
	Categories   []Categories   `json:"categories,omitempty"`
	References   []References   `json:"references,omitempty"`
	FoodTypes    []FoodTypes    `json:"foodTypes,omitempty"`
	Contacts     []Contacts     `json:"contacts,omitempty"`
	OpeningHours []OpeningHours `json:"openingHours,omitempty"`
}

// Address ... restaurent address
type Address struct {
	Label       string `json:"label,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	CountryName string `json:"countryName,omitempty"`
	State       string `json:"state,omitempty"`
	County      string `json:"county,omitempty"`
	City        string `json:"city,omitempty"`
	District    string `json:"district,omitempty"`
	Street      string `json:"street,omitempty"`
	PostalCode  string `json:"postalCode,omitempty"`
	HouseNumber string `json:"houseNumber,omitempty"`
}

//Position ... position of restaurent
type Position struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}

//Access ... access coordinates of restaurent
type Access struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}

//Categories ... restaurent categories
type Categories struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

//Supplier ... restaurent supplier
type Supplier struct {
	ID string `json:"id,omitempty"`
}

//References ... suppliers who refer this restaurent
type References struct {
	Supplier Supplier `json:"supplier,omitempty"`
	ID       string   `json:"id,omitempty"`
}

//FoodTypes ... food types offered by restaurent
type FoodTypes struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

//Phone ... phone no of place
type Phone struct {
	Value string `json:"value,omitempty"`
}

//Www ... web domain of a place
type Www struct {
	Value string `json:"value,omitempty"`
}

//Contacts ... contact info of a place
type Contacts struct {
	Phone []Phone `json:"phone,omitempty"`
	Www   []Www   `json:"www,omitempty"`
}

//Structured ... opening hours info
type Structured struct {
	Start      string `json:"start,omitempty"`
	Duration   string `json:"duration,omitempty"`
	Recurrence string `json:"recurrence,omitempty"`
}

//OpeningHours ... place opening hour
type OpeningHours struct {
	Text       []string     `json:"text,omitempty"`
	IsOpen     bool         `json:"isOpen,omitempty"`
	Structured []Structured `json:"structured,omitempty"`
}
