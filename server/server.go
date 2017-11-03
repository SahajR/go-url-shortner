package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"URL-Shortner/api"
	"net/http"
	"errors"
	"strconv"
)

type Config struct {
	ConnectionString string
	DatabaseName string
	Port int
}

func Start(config *Config) error {
	router := mux.NewRouter()
	DbHandler, err := api.NewDbHandler(config.ConnectionString, config.DatabaseName)

	if err != nil {
		fmt.Println("Unable to initialize database connection! Please check your config.")
		return errors.New("Database connection error")
	}

	APIHandler := api.NewAPIHandler(*DbHandler)

	// Add router end-points
	router.Methods("GET").PathPrefix("/shorten").HandlerFunc(APIHandler.ShortenURLHandler)
	router.Methods("GET").PathPrefix("/{id}").HandlerFunc(APIHandler.RedirectHandler)

	// Start the server
	return http.ListenAndServe("localhost:" + strconv.Itoa(config.Port), router)
}
