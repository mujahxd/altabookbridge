package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mujahxd/altabookbridge/app/features/book"
	"github.com/mujahxd/altabookbridge/config"
)

func BookRoutes(e *echo.Echo, bh book.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/books", bh.GetAllBookHandler())
	e.DELETE("/books/:booksID", bh.DeLeteBookHandler(), middleware.JWT([]byte(config.Secret)))
	e.POST("/books", bh.AddBookHandler(), middleware.JWT([]byte(config.Secret)))
	e.GET("/books/:booksID", bh.GetBookByIDHandler(), middleware.JWT([]byte(config.Secret)))
	e.PUT("/books/:booksID", bh.UpdateBookHandler(), middleware.JWT([]byte(config.Secret)))
}
