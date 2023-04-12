package config

import (
	"os"
)

type Config struct {
	DBUser      string
	DBPassword  string
	DBHost      string
	DBPort      int
	DBName      string
	TokenSecret string
}

func LoadConfig() (config Config, err error) {

	config.DBUser = os.Getenv("DBUSER")
	config.DBPassword = os.Getenv("DBPASSWORD")
	config.DBHost = os.Getenv("DBHOST")
	config.DBHost = os.Getenv("DBPORT")
	config.DBName = os.Getenv("DBNAME")
	config.TokenSecret = os.Getenv("JWT")

	return config, nil
}
