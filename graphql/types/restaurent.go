package types

import "github.com/graphql-go/graphql"

var addressType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Restaurent Address Info",
		Fields: graphql.Fields{
			"label": &graphql.Field{
				Type: graphql.String,
			},
			"countryCode": &graphql.Field{
				Type: graphql.String,
			},
			"countryName": &graphql.Field{
				Type: graphql.String,
			},
			"state": &graphql.Field{
				Type: graphql.String,
			},
			"county": &graphql.Field{
				Type: graphql.String,
			},
			"city": &graphql.Field{
				Type: graphql.String,
			},
			"district": &graphql.Field{
				Type: graphql.String,
			},
			"street": &graphql.Field{
				Type: graphql.String,
			},
			"postalCode": &graphql.Field{
				Type: graphql.String,
			},
			"houseNumber": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var positionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Restaurent Info",
		Fields: graphql.Fields{
			"lat": &graphql.Field{
				Type: graphql.String,
			},
			"lng": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var categoryAndFoodType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Restaurent Info",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"primary": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var restaurentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Restaurent Info",
		Fields: graphql.Fields{
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: addressType,
			},
			"position": &graphql.Field{
				Type: positionType,
			},
			"distance": &graphql.Field{
				Type: graphql.Int,
			},
			"categories": &graphql.Field{
				Type: graphql.NewList(categoryAndFoodType),
			},
			"foodTypes": &graphql.Field{
				Type: graphql.NewList(categoryAndFoodType),
			},
		},
	},
)
