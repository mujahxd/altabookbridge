package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mujahxd/altabookbridge/app/features/user/auth"
	"github.com/mujahxd/altabookbridge/app/features/user/handler"
	"github.com/mujahxd/altabookbridge/app/features/user/usecase"
)

func InitRoute(e *echo.Echo, h handler.Handler, authService auth.Service, userService usecase.UseCase) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/users", h.RegisterUser())
	e.POST("/login", h.Login())
	e.GET("/users", h.GetProfileUser())
	e.DELETE("/users", h.DeleteActiveUser())
	e.PUT("/users", h.UpdateProfileUser())

}
