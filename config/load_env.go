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

func InitConfig() *Config {
	var cnf = readConfig()
	if cnf == nil {
		return nil
	}
	return cnf
}

func readConfig() *Config {
	var result = new(Config)
	result.DBUser = os.Getenv("DBUser")
	result.DBPassword = os.Getenv("DBPassword")
	result.DBHost = os.Getenv("DBHost")
	result.DBPort = os.Getenv("DBPort")
	result.DBName = os.Getenv("DBName")
	result.TokenSecret = os.Getenv("JWT_SECRET")

	return result
}
