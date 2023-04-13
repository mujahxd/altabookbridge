package user

import (
	"github.com/mujahxd/altabookbridge/app/features/book/repository"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string            `json:"name" gorm:"type:varchar(45);not null"`
	Username string            `json:"username" gorm:"type:varchar(12);unique;primaryKey"`
	Password string            `json:"password" gorm:"type:varchar(255);not null"`
	Book     []repository.Book `gorm:"foreignKey:UserName"`
}
