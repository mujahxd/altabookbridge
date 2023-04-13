package repository

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description" gorm:"type:text;not null"`
	BookImage   string `json:"book_image" gorm:"type:varchar(255)"`
	Status      bool   `json:"status" gorm:"default:true"`
	UserName    string `json:"username" gorm:"type:varchar(12)"`
}
