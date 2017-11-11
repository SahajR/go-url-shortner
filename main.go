package main

import (
	"URL-Shortner/server"
	"fmt"
	"github.com/spf13/viper"
	"strconv"
)

func main() {

	// Init Viper for config in `db.json`
	viper.SetConfigName("db")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Please provide configuration!")
		return
	}

	config := &server.Config{
		ConnectionString:viper.GetString("connectionString"),
		DatabaseName:viper.GetString("dbName"),
		Port:viper.GetInt("port"),
	}

	fmt.Println("Server will be running on localhost:" + strconv.Itoa(config.Port))
	server.Start(config)
}
