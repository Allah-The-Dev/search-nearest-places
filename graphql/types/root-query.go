package types

import (
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

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"places-around": &graphql.Field{
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
						coordinates, err := httpclient.getLocationCoordinates(locationName)
						if err != nil {
							return nil, err
						}

						return getPlacesAroundGivenLocaton(coordinates)
					}
					return nil, nil
				},
			},
		},
	},
)
