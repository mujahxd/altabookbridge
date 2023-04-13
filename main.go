package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mujahxd/altabookbridge/app/features/book/handler"
	"github.com/mujahxd/altabookbridge/app/features/book/repository"
	"github.com/mujahxd/altabookbridge/app/features/book/usecase"
	"github.com/mujahxd/altabookbridge/config"
	"github.com/mujahxd/altabookbridge/routes"
	"github.com/mujahxd/altabookbridge/utils/database"
)

func main() {
	e := echo.New()
	loadConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	// database
	db := database.ConnectionDB(&loadConfig)
	database.Migrate(db)
	e.Logger.Fatal(e.Start(":8000"))

	bookModel := repository.New(db)
	bookSrv := usecase.New(bookModel)
	bookController := handler.New(bookSrv)

	routes.BookRoutes(e, bookController)
}
