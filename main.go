package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	//start the serve
	go func() {
		log.Printf("**************http server listening on port %s *************", httpServerPort)

		err := httpServer.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// tap interrupt and kill signal and gracefully shutdown server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	//block until signal is received
	sig := <-c
	log.Println("Got os signal", sig)

	//gracefully shutdown server, waiting 30 second for shutting down server
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	httpServer.Shutdown(ctx)
}

func initializeHTTPRouter() *mux.Router {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/places", handlers.PlacesHandler).Methods(http.MethodGet)

	return router
}
