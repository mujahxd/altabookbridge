package config

import (
	"os"
)

type Config struct {
	DBUser        string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
	URLCLOUDINARY string
	// CLOUDINARY_URL=cloudinary://958436536152978:DgCOQcZLn4gIJAzvOoOtL4l_ub4@dc0wgpho2
}

func InitConfig() *Config {
	var cnf = readConfig()
	if cnf == nil {
		return nil
	}
	return cnf
}

var TokenSecret string

func readConfig() *Config {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Cannot read config variable")
	// 	return nil
	// }
	var result = new(Config)
	result.DBUser = os.Getenv("DBUser")
	result.DBPassword = os.Getenv("DBPassword")
	result.DBHost = os.Getenv("DBHost")
	result.DBPort = os.Getenv("DBPort")
	result.DBName = os.Getenv("DBName")
	result.URLCLOUDINARY = os.Getenv("URLCLOUDINARY")
	TokenSecret = os.Getenv("JWT_SECRET")

	return result
}
