package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"URL-Shortner/api"
	"net/http"
	"errors"
)

func Start() error {
	router := mux.NewRouter()
	DbHandler, err := api.NewDbHandler("mongodb://sahajr_url_shortner:TWEEKxCRAIG@ds143754.mlab.com:43754/sahajr-website")

	if err != nil {
		fmt.Println("Unable to initialize database connection!")
		return errors.New("Database connection error")
	}

	APIHandler := api.NewAPIHandler(*DbHandler)

	apirouter := router.PathPrefix("/").Subrouter()
	apirouter.Methods("GET").Path("/{id}").HandlerFunc(APIHandler.RedirectHandler)
	apirouter.Methods("GET").PathPrefix("/shorten").HandlerFunc(APIHandler.ShortenURLHandler)

	return http.ListenAndServe("localhost:3003", router)
}
