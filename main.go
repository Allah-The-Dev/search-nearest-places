package main

import (
	"log"
	"net/http"
	"search-nearest-places/handlers"
	"time"

	"github.com/gorilla/mux"
)

// Tweak configuration values here.
const (
	httpServerPort    = ":9080"
	readHeaderTimeout = 1 * time.Second
	writeTimeout      = 10 * time.Second
	idleTimeout       = 90 * time.Second
	maxHeaderBytes    = http.DefaultMaxHeaderBytes
)

func main() {

	router := initializeHTTPRouter()

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

func initializeHTTPRouter() *mux.Router {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/places", handlers.PlacesHandler).Methods(http.MethodGet)

	return router
}
