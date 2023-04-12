package database

import (
	"fmt"
	"log"

	"github.com/mujahxd/altabookbridge/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
func ConnectionDB(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/abb?charset=utf8mb4&parseTime=True&loc=Local", config.DBPassword)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected Successfully to the database")
	return db
}
