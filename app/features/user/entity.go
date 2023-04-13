package user

import "github.com/labstack/echo/v4"

type Core struct {
	Name     string
	Username string
	Password string
}

type Repository interface {
	Save(user Core) (Core, error)
}

type UseCase interface {
	RegisterUser(user Core) error
}

type Handler interface {
	RegisterUser() echo.HandlerFunc
}
