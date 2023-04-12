package repository

import (
	"errors"
	"log"

	"github.com/mujahxd/altabookbridge/app/features/book"
	"gorm.io/gorm"
)

type bookModel struct {
	db *gorm.DB
}

func New(d *gorm.DB) book.Repository {
	return &bookModel{
		db: d,
	}
}

func (bm *bookModel) GetAllBook(offset int, limit int) ([]book.Core, error) {
	res := []book.Core{}

	if err := bm.db.Table("books").Select("books.title as title, books.bookimage as bookimage, books.status as status, books.username as username").Scan(&res).Error; err != nil {
		log.Println("terjadi error saat select book", err.Error())
		return nil, err
	}

	return res, nil
}

func (bm *bookModel) DeleteBook(username string, bookID uint) error {
	b := &Book{}
	if err := bm.db.First(b, bookID, username).Error; err != nil {
		log.Println("error in finding book for delete")
		return errors.New("book not found")
	}

	if err := bm.db.Delete(b).Error; err != nil {
		log.Println("error in delete book")
		return err
	}
	return nil
}
