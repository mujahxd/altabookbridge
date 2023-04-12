package helper

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(ID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	DecodeToken(token *jwt.Token) string
}

type jwtService struct {
}

var SECRET_KEY = []byte(os.Getenv("JWT"))

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(ID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["id"] = ID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
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

func (s *jwtService) DecodeToken(token *jwt.Token) string {
	if token.Valid {
		data := token.Claims.(jwt.MapClaims)
		user_id := data["id"].(string)

		return user_id
	}

	return ""
}
