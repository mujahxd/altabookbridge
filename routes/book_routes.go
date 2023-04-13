package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mujahxd/altabookbridge/app/features/book"
)

func BookRoutes(e *echo.Echo, bh book.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/books", bh.GetAllBookHandler())
	e.DELETE("/books", bh.DeLeteBookHandler())
	e.POST("/books", bh.AddBookHandler())
}
