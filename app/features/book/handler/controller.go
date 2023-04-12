package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/mujahxd/altabookbridge/app/features/book"
	"github.com/mujahxd/altabookbridge/helper"
)

func response(code int, msg string, status string, data any) (int, map[string]any) {
	res := map[string]any{}
	res["code"] = code
	res["message"] = msg
	res["status"] = status
	if data != nil {
		res["data"] = data
	}

	return code, res
}

type bookController struct {
	srv book.UseCase
}

func New(s book.UseCase) book.Handler {
	return &bookController{
		srv: s,
	}
}

func (bc *bookController) GetAllBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		offset, err := strconv.Atoi(c.QueryParam("offset"))
		if err != nil {
			c.Logger().Error("error on converting offset")
			return c.JSON(response(http.StatusBadRequest, "previous data or value are not allowed", "failed", nil))
		}

		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			c.Logger().Error("error on converting offset")
			return c.JSON(response(http.StatusBadRequest, "previous data or value are not allowed", "failed", nil))
		}

		res, err := bc.srv.GetAllBookLogic(offset, limit)
		if err != nil {
			c.Logger().Error("error on calling bookmodel in handler GetAllBook", err.Error())
			return c.JSON(response(http.StatusBadRequest, "previous data or value are not allowed", "error", nil))
		}

		return c.JSON(response(http.StatusOK, "succes to find all book", "succes", res))
	}
}

func (bc *bookController) DeLeteBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		validate := helper.NewService()

		token, err := validate.ValidateToken(c.Request().Header.Get("Authorization"))
		if err != nil {
			c.Logger().Error("error from validating jwt")
			return c.JSON(response(http.StatusBadRequest, "invalid token", "error", nil))
		}
		username := validate.DecodeToken(token)

		bookID, err := strconv.Atoi(c.Param("bookID"))
		if err != nil {
			return c.JSON(response(http.StatusBadRequest, "invalid book ID", "error", nil))
		}

		if err := bc.srv.DeleteBookLogic(username, uint(bookID)); err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Println("error occurs on deletebookHandler, in FINDING book(bad user)", err.Error())
				return c.JSON(response(http.StatusBadRequest, "book not found", "error", nil))
			}
			log.Println("error occurs in in deleteHandler for delete method", err.Error())
			return c.JSON(response(http.StatusInternalServerError, "inter server error", "error", nil))
		}

		return c.JSON(response(http.StatusOK, "succes to delete book", "succes", nil))
	}
}
