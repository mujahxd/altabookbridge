package helper

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/mujahxd/altabookbridge/config"
)

func GenerateToken(username string) (string, error) {
	claim := jwt.MapClaims{}
	claim["username"] = username

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(config.Secret), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func DecodeToken(token *jwt.Token) string {
	if token.Valid {
		data := token.Claims.(jwt.MapClaims)
		username := data["username"].(string)

		return username
	}

	return ""
}
