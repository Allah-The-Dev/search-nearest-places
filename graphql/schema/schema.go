package schema

import (
	"github.com/graphql-go/graphql"

	"search-nearest-places/graphql/types"
)

//Schema ... GraphQL root schema
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: types.QueryType,
	},
)
