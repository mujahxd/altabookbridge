package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mujahxd/altabookbridge/app/features/user"
	"github.com/mujahxd/altabookbridge/app/features/user/repository"
	"github.com/mujahxd/altabookbridge/app/features/user/usecase"
	"github.com/mujahxd/altabookbridge/config"
	"github.com/mujahxd/altabookbridge/routes"
	"github.com/mujahxd/altabookbridge/utils/database"
)

func main() {
	e := echo.New()
	loadConfig := config.InitConfig()
	db := database.ConnectionDB(loadConfig)

	// database
	database.Migrate(db)

	userModel := repository.NewModel(db)
	userUsecase := usecase.NewLogic(userModel)
	userHandler := user.NewHandler(userUsecase)

	routes.InitRoute(e, userHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
