package auth

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/mujahxd/altabookbridge/app/features/user/usecase"
	"github.com/mujahxd/altabookbridge/helper"
)

func AuthMiddleware(authService Service, userService usecase.UseCase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
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

			token, err := authService.ValidateToken(tokenString)
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
			user, err := userService.GetUserByUsername(username)
			if err != nil {
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				return c.JSON(http.StatusUnauthorized, response)
			}

			c.Set("currentUser", user)
			return next(c)
		}
	}
}
