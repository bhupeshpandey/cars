package main

import (
	"github.com/bhupeshpandey/cars/cache"
	"github.com/bhupeshpandey/cars/config"
	"github.com/bhupeshpandey/cars/db"
	"github.com/bhupeshpandey/cars/server"
	"log"
)

func main() {
    // TODO call your code here.
	appConfig, err := config.ReadConfig()

	if err != nil {
		log.Fatal(err)
	}

	db, err := db.New(appConfig.DBConfig)

	if err != nil {
		log.Fatal(err)
	}
	s := server.New(appConfig.ServerConfig, cache.New(appConfig.CacheConfig), db)

	if s == nil {
		log.Fatal("Unable to create server")
	}

	s.Start()
}
