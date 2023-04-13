package repository

import (
	"time"

	"github.com/mujahxd/altabookbridge/app/features/book/repository"
	"gorm.io/gorm"
)

type Rent struct {
	ID        uint64             `gorm:"primaryKey"`
	UserName  string             `gorm:"not null"`
	RentedAt  time.Time          `gorm:"not null"`
	CreatedAt time.Time          `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt     `gorm:"index"`
	UpdatedAt time.Time          `gorm:"autoUpdateTime"`
	Books     []*repository.Book `gorm:"many2many:books_has_rent"`
}

type BooksHasRent struct {
	BookID     uint64     `gorm:"primaryKey"`
	RentID     uint64     `gorm:"primaryKey"`
	ReturnedAt *time.Time `gorm:"default:null"`
}
