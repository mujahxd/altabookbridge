package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser         string
	DBPassword     string
	DBHost         string
	DBPort         string
	DBName         string
	TokenSecret    string
	URLCLOURDINARY string
	// CLOUDINARY_URL=cloudinary://958436536152978:DgCOQcZLn4gIJAzvOoOtL4l_ub4@dc0wgpho2
}

func LoadConfig() (config Config, err error) {

	err = godotenv.Load(".env")
	if err != nil {
		log.Println("cannot load env")
	}

	config.DBUser = os.Getenv("DBUser")
	config.DBPassword = os.Getenv("DBPassword")
	config.DBHost = os.Getenv("DBHost")
	config.DBPort = os.Getenv("DBPort")
	config.DBName = os.Getenv("DBName")
	config.TokenSecret = os.Getenv("JWT_SECRET")
	config.URLCLOURDINARY = os.Getenv("URLCLOUDRINARY")

	return config, nil
}
