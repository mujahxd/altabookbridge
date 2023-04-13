package handler

import "github.com/labstack/echo/v4"

type Handler interface {
	RegisterUser() echo.HandlerFunc
}
