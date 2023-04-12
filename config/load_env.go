package config

import (
	"os"
)

type Config struct {
	DBUser      string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
	TokenSecret string
}

func LoadConfig() (config Config, err error) {

	config.DBUser = os.Getenv("DBUser")
	config.DBPassword = os.Getenv("DBPassword")
	config.DBHost = os.Getenv("DBHost")
	config.DBPort = os.Getenv("DBPort")
	config.DBName = os.Getenv("DBName")
	config.TokenSecret = os.Getenv("JWT_SECRET")

	return config, nil
}
