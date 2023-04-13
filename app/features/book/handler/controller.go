package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/mujahxd/altabookbridge/app/features/book"
	"github.com/mujahxd/altabookbridge/helper"
)

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
			response := helper.APIResponse("previous data or value are not allowed", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
		}

		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			c.Logger().Error("error on converting offset")
			response := helper.APIResponse("previous data or value are not allowed", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
		}

		res, err := bc.srv.GetAllBookLogic(offset, limit)
		if err != nil {
			c.Logger().Error("error on calling bookmodel in handler GetAllBook", err.Error())
			response := helper.APIResponse("previous data or value are not allowed", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
		}
		response := helper.APIResponse("succes to get all data", http.StatusOK, "error", res)
		return c.JSON(http.StatusOK, response)
	}
}

func (bc *bookController) DeLeteBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := helper.DecodeToken(c.Get("username").(*jwt.Token))

		bookID, err := strconv.Atoi(c.Param("bookID"))
		if err != nil {
			response := helper.APIResponse("invalid book ID", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := bc.srv.DeleteBookLogic(username, uint(bookID)); err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Println("error occurs on deletebookHandler, in FINDING book(bad user)", err.Error())
				response := helper.APIResponse("book not found", http.StatusNotFound, "error", nil)
				return c.JSON(http.StatusNotFound, response)
			}
			log.Println("error occurs in in deleteHandler for delete method", err.Error())
			response := helper.APIResponse("server error", http.StatusInternalServerError, "error", nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
		response := helper.APIResponse("succes to delete book", http.StatusOK, "succes", nil)
		return c.JSON(http.StatusOK, response)
	}
}

func (bc *bookController) AddBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// username := helper.DecodeToken(c.Get("username").(*jwt.Token))
		username := c.FormValue("username")

		description := c.FormValue("description")
		title := c.FormValue("title")

		bookImageFile, err := c.FormFile("book_image")
		if err != nil {
			log.Println("error occurs on reading form image")
			response := helper.APIResponse("bad request", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
		}

		err = bc.srv.AddBookLogic(username, description, title, bookImageFile)
		if err != nil {
			if strings.Contains(err.Error(), "server cloudinary") {
				c.Logger().Error("error from calling upload third party server")
				response := helper.APIResponse("server error", http.StatusInternalServerError, "error", nil)
				return c.JSON(http.StatusInternalServerError, response)
			}
			c.Logger().Error("error from calling addbooklogic bad request")
			response := helper.APIResponse("bad request", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
		}

		response := helper.APIResponse("server error", http.StatusOK, "error", nil)
		return c.JSON(http.StatusOK, response)

		// 	var newBook BookRequest
		// 	form, err := c.MultipartForm()
		// 	if err != nil {
		// 		response := helper.APIResponse("failed to read form data", http.StatusBadRequest, "error", nil)
		// 		c.JSON(http.StatusBadRequest, response)
		// 	}

		// 	description := c.FormValue("description")
		// 	title := c.FormValue("title")
		// 	bookImageFile, err := c.FormFile("book_image")
		// 	if err != nil {
		// 		response := helper.APIResponse("failed to read book image", http.StatusBadRequest, "error", nil)
		// 		return c.JSON(http.StatusBadRequest, response)
		// 	}
		// 	src, err := bookImageFile.Open()
		// 	if err != nil {
		// 		response := helper.APIResponse("file is corrupt", http.StatusBadRequest, "error", nil)
		// 		return c.JSON(http.StatusBadRequest, response)
		// 	}
		// 	defer src.Close()

		// 	fileByte, err := io.ReadAll(src)
		// 	if err != nil {
		// 		response := helper.APIResponse("file is corrupt (EOF)", http.StatusBadRequest, "error", nil)
		// 		return c.JSON(http.StatusBadRequest, response)
		// 	}

		// 	fileType := http.DetectContentType(fileByte)
		// 	if fileType != "image/jpeg" && fileType != "image/png" {
		// 		response := helper.APIResponse("file is not jpg or png", http.StatusBadRequest, "failed", nil)
		// 		return c.JSON(http.StatusBadRequest, response)
		// 	}

		// 	fileName := "app/feature/book/uploads/" + strconv.FormatInt(time.Now().Unix(), 10)
		// 	// ini harusnya salah (ikutin path di server)
		// 	err = os.WriteFile(fileName, fileByte, 0777)
		// 	if err != nil {
		// 		response := helper.APIResponse("internal server error", http.StatusInternalServerError, "error", nil)
		// 		return c.JSON(http.StatusBadRequest, response)
		// 	}

		// 	response := helper.APIResponse("succes to create file", http.StatusCreated, "succes", nil)
		// 	return c.JSON(http.StatusCreated, response)
		// }
	}
}
