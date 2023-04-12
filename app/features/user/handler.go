package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mujahxd/altabookbridge/app/features/user/data"
)

type handler struct {
	service UseCase
}

func NewHandler(service UseCase) *handler {
	return &handler{service}
}

func (h *handler) RegisterUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input data.RegisterUserInput
		err := c.Bind(&input)
		if err != nil {
			c.Logger().Error("error bind: ", err.Error())
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusOK, nil)
	}
}
