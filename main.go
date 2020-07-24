package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"

	"search-nearest-places/graphql/schema"
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
		Schema:        schema.Schema,
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
