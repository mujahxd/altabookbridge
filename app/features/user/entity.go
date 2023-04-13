package user

import (
	bRepo "github.com/mujahxd/altabookbridge/app/features/book/repository"
	rRepo "github.com/mujahxd/altabookbridge/app/features/rent/repository"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string `json:"name" gorm:"type:varchar(45);not null"`
	Username       string `json:"username" gorm:"type:varchar(12);unique;primaryKey"`
	Password       string `json:"password" gorm:"type:varchar(255);not null"`
	AvatarFileName string
	Book           []bRepo.Book `gorm:"foreignKey:UserName;references:Username"`
	Loans          []rRepo.Rent `gorm:"foreignKey:UserName;references:Username"`
}
