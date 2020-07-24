package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"

	"search-nearest-places/graphql/schema"
)

// Tweak configuration values here.
const (
	httpServerPort    = ":9080"
	readHeaderTimeout = 1 * time.Second
	writeTimeout      = 10 * time.Second
	idleTimeout       = 90 * time.Second
	maxHeaderBytes    = http.DefaultMaxHeaderBytes
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
	log.Printf("receieved query is %v", requestBody)

	result := graphql.Do(graphql.Params{
		Schema:        schema.Schema,
		RequestString: requestBody.query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	json.NewEncoder(w).Encode(result)
}

func main() {

	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/graphql", doGraphQL).Methods(http.MethodPost)

	// Configure the HTTP server.
	httpServer := &http.Server{
		Addr:              httpServerPort,
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}
	log.Printf("**************http server listening on port %s *************", httpServerPort)
	log.Fatal(httpServer.ListenAndServe())

}
