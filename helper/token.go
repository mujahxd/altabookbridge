package helper

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

var SECRET_KEY = []byte(os.Getenv("JWT"))

func GenerateToken(username string) (string, error) {
	claim := jwt.MapClaims{}
	claim["username"] = username

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
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

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func DecodeToken(token *jwt.Token) string {
	if token.Valid {
		data := token.Claims.(jwt.MapClaims)
		user_id := data["id"].(string)

		return user_id
	}

	return ""
}
