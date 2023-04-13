package main

import (
	"github.com/labstack/echo/v4"
	bookhandler "github.com/mujahxd/altabookbridge/app/features/book/handler"
	bookrepo "github.com/mujahxd/altabookbridge/app/features/book/repository"
	bookusecase "github.com/mujahxd/altabookbridge/app/features/book/usecase"
	"github.com/mujahxd/altabookbridge/app/features/user/auth"
	userhandler "github.com/mujahxd/altabookbridge/app/features/user/handler"
	userrepo "github.com/mujahxd/altabookbridge/app/features/user/repository"
	userusecase "github.com/mujahxd/altabookbridge/app/features/user/usecase"
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

	authService := auth.NewService()
	userModel := userrepo.NewModel(db)
	userUsecase := userusecase.NewLogic(userModel)
	userHandler := userhandler.NewHandler(userUsecase, authService)

	bookModel := bookrepo.New(db)
	bookSrv := bookusecase.New(bookModel)
	bookController := bookhandler.New(bookSrv)

	routes.InitRoute(e, userHandler, authService, userUsecase)
	routes.BookRoutes(e, bookController)

	e.Logger.Fatal(e.Start(":8000"))
}
