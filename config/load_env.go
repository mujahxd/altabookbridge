package config

import (
	"os"
)

type Config struct {
	DBPassword  string
	TokenSecret string
}

func LoadConfig() (config Config, err error) {

	password := os.Getenv("PASSWORD")
	tokenSecret := os.Getenv("JWT")

	config.DBPassword = password
	config.TokenSecret = tokenSecret

	return config, nil
}
