package database

import (
	bookrepo "github.com/mujahxd/altabookbridge/app/features/book/repository"
	"github.com/mujahxd/altabookbridge/app/features/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	users := user.User{}
	books := bookrepo.Book{}

	db.AutoMigrate(users)
	db.AutoMigrate(books)
}
