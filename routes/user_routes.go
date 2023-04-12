package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mujahxd/altabookbridge/app/features/user"
)

func InitRoute(e *echo.Echo, h user.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.POST("/users", h.RegisterUser())
}
