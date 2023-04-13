package handler

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/mujahxd/altabookbridge/app/features/user"
	"github.com/mujahxd/altabookbridge/app/features/user/auth"
	"github.com/mujahxd/altabookbridge/app/features/user/data"
	"github.com/mujahxd/altabookbridge/app/features/user/usecase"
	"github.com/mujahxd/altabookbridge/helper"
)

type handler struct {
	userService usecase.UseCase
	authService auth.Service
}

func NewHandler(userService usecase.UseCase, authService auth.Service) *handler {
	return &handler{userService, authService}
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
		_, err = h.userService.RegisterUser(input)
		if err != nil {
			response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
		}

		response := helper.APIResponse("Account has been registered", http.StatusCreated, "success", nil)
		return c.JSON(http.StatusCreated, response)
	}
}

func (h *handler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input data.LoginInput
		err := c.Bind(&input)
		if err != nil {
			errors := helper.FormatValidationError(err)
			errorMessage := echo.Map{"errors": errors}
			response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
			return c.JSON(http.StatusUnprocessableEntity, response)
		}
		loggedinUser, err := h.userService.Login(input)
		if err != nil {
			errorMessage := echo.Map{"errors": err.Error()}
			response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
			return c.JSON(http.StatusBadRequest, response)
		}

		token, err := helper.GenerateToken(loggedinUser.Username)
		if err != nil {
			// errorMessage := echo.Map{"errors": err.Error()}
			response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
		}

		formatter := user.FormatLoginUser(loggedinUser, token)
		response := helper.APIResponse("Successfully loggedin", http.StatusOK, "success", formatter)
		return c.JSON(http.StatusOK, response)
	}
}

func (h *handler) DeleteActiveUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		// bearer {token}
		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := helper.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		username := claim["username"].(string)

		// hapus pengguna dengan username yang sesuai
		err = h.userService.DeleteUser(username)
		if err != nil {
			c.Logger().Error(err.Error())
			response := helper.APIResponse("failed to delete user", http.StatusInternalServerError, "error", nil)
			return c.JSON(http.StatusInternalServerError, response)
		}

		// kirim respons jika pengguna berhasil dihapus

		response := helper.APIResponse("User has been deleted", http.StatusOK, "success", nil)
		return c.JSON(http.StatusOK, response)
	}

}

func (h *handler) GetProfileUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		// bearer {token}
		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := helper.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		username := claim["username"].(string)
		activeUser, err := h.userService.GetUserByUsername(username)
		if err != nil {
			c.Logger().Error(err.Error())
			response := helper.APIResponse("failed to get profile user", http.StatusInternalServerError, "error", nil)
			return c.JSON(http.StatusInternalServerError, response)
		}

		// kirim respons jika pengguna berhasil dihapus
		formatter := user.FormatGetProfile(activeUser)
		response := helper.APIResponse("User profile successfully displayed", http.StatusOK, "success", formatter)
		return c.JSON(http.StatusOK, response)
	}
}
