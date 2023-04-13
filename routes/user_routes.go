package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mujahxd/altabookbridge/app/features/user/handler"
)

func InitRoute(e *echo.Echo, h handler.Handler) {
	e.POST("/users", h.RegisterUser())
	e.POST("/login", h.Login())
}
