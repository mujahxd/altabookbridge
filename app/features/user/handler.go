package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mujahxd/altabookbridge/app/features/user/data"
	"github.com/mujahxd/altabookbridge/helper"
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
			errors := helper.FormatValidationError(err)
			errorMessage := echo.Map{"errors": errors}
			response := helper.APIResponse("Register account failed, field must be filled", http.StatusUnprocessableEntity, "error", errorMessage)
			return c.JSON(http.StatusUnprocessableEntity, response)
		}
		_, err = h.service.RegisterUser(input)
		if err != nil {
			response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}
		response := helper.APIResponse("Account has been registered", http.StatusCreated, "success", nil)
		return c.JSON(http.StatusOK, response)
	}
}
