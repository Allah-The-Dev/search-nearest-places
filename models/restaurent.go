package models

//RestaurentItems ... items from here API
type RestaurentItems struct {
	Items []Restaurent `json:"items"`
}

//Restaurent ... higher level restaurent info
type Restaurent struct {
	Title      string       `json:"title"`
	ID         string       `json:"id"`
	ResultType string       `json:"resultType"`
	Address    Address      `json:"address"`
	Position   Position     `json:"position"`
	Access     []Access     `json:"access"`
	Distance   int          `json:"distance"`
	Categories []Categories `json:"categories"`
	References []References `json:"references"`
	FoodTypes  []FoodTypes  `json:"foodTypes"`
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
