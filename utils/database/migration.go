package database

import (
	bookrepo "github.com/mujahxd/altabookbridge/app/features/book/repository"
	userrepo "github.com/mujahxd/altabookbridge/app/features/user/repository"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	books := bookrepo.Book{}
	users := userrepo.User{}

	db.AutoMigrate(books)
	db.AutoMigrate(users)
}
