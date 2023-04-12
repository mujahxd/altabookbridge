package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBPassword  string
	TokenSecret string
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
		return
	}

	password := os.Getenv("PASSWORD")
	tokenSecret := os.Getenv("TOKEN_SECRET")

	config.DBPassword = password
	config.TokenSecret = tokenSecret

	return config, nil
}
