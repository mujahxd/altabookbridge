package database

import (
	bookrepo "github.com/mujahxd/altabookbridge/app/features/book/repository"
	rentrepo "github.com/mujahxd/altabookbridge/app/features/rent/repository"
	"github.com/mujahxd/altabookbridge/app/features/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	users := user.User{}
	books := bookrepo.Book{}
	rent := rentrepo.Rent{}

	db.AutoMigrate(users)
	db.AutoMigrate(books)
	db.AutoMigrate(rent)
}
