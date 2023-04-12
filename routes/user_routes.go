package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mujahxd/altabookbridge/app/features/user"
)

func InitRoute(e *echo.Echo, h user.Handler) {
	e.POST("/users", h.RegisterUser())
}
