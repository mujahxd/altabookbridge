package database

import (
	bookrepo "github.com/mujahxd/altabookbridge/app/features/book/repository"
	userrepo "github.com/mujahxd/altabookbridge/app/features/user/repository"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	users := userrepo.User{}
	books := bookrepo.Book{}

	db.AutoMigrate(users)
	db.AutoMigrate(books)
}
