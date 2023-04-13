package user

import (
	"github.com/labstack/echo/v4"
	"github.com/mujahxd/altabookbridge/app/features/user/data"
)

type Core struct {
	Name     string
	Username string
	Password string
}

type Repository interface {
	Save(user Core) (Core, error)
	FindByUsername(username string) (Core, error)
}

type UseCase interface {
	RegisterUser(input data.RegisterUserInput) (Core, error)
}

type Handler interface {
	RegisterUser() echo.HandlerFunc
}
