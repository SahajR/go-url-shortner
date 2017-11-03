package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"URL-Shortner/api"
	"net/http"
	"errors"
	"github.com/spf13/viper"
)

func Start() error {
	router := mux.NewRouter()
	DbHandler, err := api.NewDbHandler(getDbURL())

	if err != nil {
		fmt.Println("Unable to initialize database connection!")
		return errors.New("Database connection error")
	}

	APIHandler := api.NewAPIHandler(*DbHandler)

	router.Methods("GET").PathPrefix("/shorten").HandlerFunc(APIHandler.ShortenURLHandler)
	router.Methods("GET").PathPrefix("/{id}").HandlerFunc(APIHandler.RedirectHandler)

	return http.ListenAndServe("localhost:3003", router)
}

func getDbURL() string {
	viper.SetConfigName("db")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	viper.ReadInConfig()
	return viper.GetString("connectionString")
}
