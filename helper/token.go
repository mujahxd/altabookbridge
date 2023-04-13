package helper

import (
	"os"

	"github.com/golang-jwt/jwt"
)

var SECRET_KEY = []byte(os.Getenv("JWT"))

func GenerateToken(ID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["id"] = ID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func DecodeToken(token *jwt.Token) string {
	if token.Valid {
		data := token.Claims.(jwt.MapClaims)
		username := data["username"].(string)

		return username
	}

	return ""
}
