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

func InitConfig() *Config {
	var cnf = readConfig()
	if cnf == nil {
		return nil
	}
	return cnf
}
var Secret string
func readConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot read config variable")
		return nil
	}
	var result = new(Config)
	result.DBUser = os.Getenv("DBUser")
	result.DBPassword = os.Getenv("DBPassword")
	result.DBHost = os.Getenv("DBHost")
	result.DBPort = os.Getenv("DBPort")
	result.DBName = os.Getenv("DBName")
	Secret = os.Getenv("JWT")
	result.URLCLOURDINARY = os.Getenv("URLCLOUDRINARY")

	return result
}
