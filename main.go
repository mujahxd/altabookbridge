package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mujahxd/altabookbridge/config"
)

func main() {
	e := echo.New()
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	e.Logger.Fatal(e.Start(":8000"))
	// database
	// db:= database.ConnectionDB(&loadConfig)
}
