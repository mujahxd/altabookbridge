package repository

import (
	"errors"
	"log"
	"mime/multipart"

	"github.com/mujahxd/altabookbridge/app/features/book"
	"github.com/mujahxd/altabookbridge/helper"
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

	if err := bm.db.Table("books").
		Select("books.title as title, books.bookimage as book_image, books.status as status, books.username as username").
		Offset(offset).
		Limit(limit).
		Scan(&res).Error; err != nil {
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

func (bm *bookModel) AddBook(username string, description string, title string, bookFile *multipart.FileHeader) error {
	bookurl, err := helper.Upload(bookFile)
	if err != nil {
		log.Println("errors from calling uploader", err.Error())
		return errors.New("cannot upload image to server")
	}

	var addNewBook = &Book{
		Title:       title,
		Description: description,
		BookImage:   bookurl,
		UserName:    username,
	}

	if err := bm.db.Create(addNewBook).Error; err != nil {
		log.Println("error in creating book for add book")
		return errors.New("cannot create book")
	}

	return nil
}

func (bm *bookModel) GetBookByID(bookID uint) (book.Core, error) {
	var res book.Core
	var find Book

	result := bm.db.First(&find, bookID)
	if result.Error != nil {
		log.Printf("error finding books with ID %d: %v", bookID, result.Error)
		return book.Core{}, errors.New("error finding books")
	}

	res.Username = find.UserName
	res.Title = find.Title
	res.Description = find.Description
	res.BookImage = find.BookImage
	return res, nil
}

func (bm *bookModel) UpdateBook(bookID uint, title string, description string, bookImage *multipart.FileHeader) error {
	var UpdateBook Book

	bookurl, err := helper.Upload(bookImage)
	if err != nil {
		log.Println("errors from calling uploader", err.Error())
		return errors.New("cannot upload image to server")
	}

	UpdateBook.Title = title
	UpdateBook.Description = description
	UpdateBook.BookImage = bookurl

	// if err := bm.db.Save(&UpdateBook).Where("id = ?", bookID).Error; err != nil {
	// 	log.Println("error from query save book in update book")
	// 	return err
	// }
	if err := bm.db.Model(&Book{}).Where("id = ?", bookID).Updates(UpdateBook).Error; err != nil {
		log.Println("error from query save book in update book")
		return err
	}

	return nil
}
