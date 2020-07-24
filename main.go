package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
)

const (
	hereAPIKey          = "uZ9InITCNxhCIlW--t1RDnYlSplGAMkktR2UP1D_wok"
	hereBrowseAPIURL    = "https://browse.search.hereapi.com/v1/browse?at=%s3&limit=3&categories=%s&apiKey=%s"
	hereGecodeAPIURL    = "https://geocode.search.hereapi.com/v1/geocode?q=%s&apiKey=%s"
	hereRestaurentCatID = "100-1000"
)

type placeInfo struct {
	name     string `json:"name"`
	distance int    `json:"distance`
}

type places struct {
	restaurent placeInfo `json:"restaurent`
}

type location struct {
	name string
	long float64
	lat  float64
}

var placeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PlacesAround",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"distance": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single product by id
			   http://localhost:8080/product?query={product(id:1){name,info,price}}
			*/
			"places-around": &graphql.Field{
				Type:        placeType,
				Description: "Get near by places",
				Args: graphql.FieldConfigArgument{
					"location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					locationName, ok := p.Args["location"].(string)
					if ok {
						locationInfo, err := getLocationCoordinates(locationName)
						if err != nil {
							return nil, err
						}

						return getPlacesAroundGivenLocaton(locationInfo)
					}
					return nil, nil
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func doGraphQL(w http.ResponseWriter, r *http.Request) {
	//decode body
	var requestBody struct {
		query string
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		log.Printf("%v", err)
		http.Error(w, fmt.Sprintf("request body is not in correct format"), http.StatusBadRequest)
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: requestBody.query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
}

func main() {

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("graphql", doGraphQL).Methods(http.MethodGet)

}

type hereGeoCodeItem struct {
	position struct {
		lat float64
		lng float64
	}
}

type hereGeoCodeResponse struct {
	items []hereGeoCodeItem
}

func getLocationCoordinates(locationName string) (*location, error) {
	url := fmt.Sprintf(hereBrowseAPIURL, locationName, hereAPIKey)
	responseBody, err := doHTTPGet(url)
	if err != nil {
		return nil, err
	}
	var hereGeoCodeRes hereGeoCodeResponse
	json.NewDecoder(*responseBody).Decode(&hereGeoCodeRes)
	return nil, err
}

func getPlacesAroundGivenLocaton(locationInfo *location) (*places, error) {
	return nil, nil
}

func doHTTPGet(url string) (*io.ReadCloser, error) {

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("here geocode api returned status code " + response.StatusCode)
	}

	return &response.Body, nil
}
