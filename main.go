package main

import (
	"log"

	"github.com/mujahxd/altabookbridge/config"
	"github.com/mujahxd/altabookbridge/utils/database"
)

func main() {
	loadConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	// database
	db:= database.ConnectionDB(&loadConfig)
}
