package types

import (
	"search-nearest-places/httpclient"

	"github.com/graphql-go/graphql"
)

var placesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PlacesAround",
		Fields: graphql.Fields{
			"restaurent": &graphql.Field{
				Type: restaurentType,
			},
		},
	},
)

//RootQueryType ... root query
var RootQueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"nearbyplaces": &graphql.Field{
				Type:        placesType,
				Description: "Get near by places",
				Args: graphql.FieldConfigArgument{
					"location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					locationName, ok := p.Args["location"].(string)
					if ok {
						coordinates, err := httpclient.GetLocationCoordinates(locationName)
						if err != nil {
							return nil, err
						}

						return httpclient.GetPlacesAroundGivenLocaton(coordinates)
					}
					return nil, nil
				},
			},
		},
	},
)
