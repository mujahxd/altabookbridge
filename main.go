package main

import (
	"log"

	"github.com/mujahxd/altabookbridge/config.go"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	// database
	db := config.ConnectionDB(&loadConfig)
}
