package rent

import (
	"time"

	"github.com/mujahxd/altabookbridge/app/features/book/repository"
	"gorm.io/gorm"
)

type CoreRent struct {
	ID        uint64             `json:"id" gorm:"primaryKey"`
	UserName  string             `json:"username" gorm:"not null"`
	RentedAt  time.Time          `json:"rented_at" gorm:"not null"`
	CreatedAt time.Time          `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt     `json:"deleted_at" gorm:"index"`
	UpdatedAt time.Time          `json:"updated_at" gorm:"autoUpdateTime"`
	Books     []*repository.Book `gorm:"many2many:books_has_rent"`
}

type CoreBooksHasRent struct {
	BookID     uint64     `gorm:"primaryKey"`
	RentID     uint64     `gorm:"primaryKey"`
	ReturnedAt *time.Time `gorm:"default:null"`
}

type Repository interface {
	//username, 
	CheckOutRent() error
}

type UseCase interface {
	CheckOutRentLogic() error
}

type Handler interface {
	CheckOutRentHandler() error
}
