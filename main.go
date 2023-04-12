package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mujahxd/altabookbridge/app/features/user"
	"github.com/mujahxd/altabookbridge/app/features/user/repository"
	"github.com/mujahxd/altabookbridge/app/features/user/usecase"
	"github.com/mujahxd/altabookbridge/config"
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

	userModel := repository.NewModel(db)
	userUsecase := usecase.NewLogic(userModel)
	userHandler := user.NewHandler(userUsecase)

	e.Logger.Fatal(e.Start(":8000"))
}
