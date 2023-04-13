package usecase

import (
	"errors"
	"log"
	"mime/multipart"
	"strings"

	"github.com/mujahxd/altabookbridge/app/features/book"
)

type bookModel struct {
	repo book.Repository
}

func New(br book.Repository) book.UseCase {
	return &bookModel{
		repo: br,
	}
}

func (bm *bookModel) GetAllBookLogic(offset int, limit int) ([]book.Core, error) {
	result, err := bm.repo.GetAllBook(offset, limit)
	if err != nil {
		if strings.Contains(err.Error(), "too much") {
			log.Println("error occurs in GetAllBookLogic (bad users)", err.Error())
			return nil, errors.New("bad request from user")
		}
		log.Println("error occurs in GetAllBookLogic", err.Error())
		return nil, errors.New("internal server error")
	}

	return result, nil
}

func (bm *bookModel) DeleteBookLogic(username string, bookID uint) error {
	if err := bm.repo.DeleteBook(username, bookID); err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Println("error occurs on deletebooklogic, in finding book(bad user)", err.Error())
			return errors.New("bad request from user")
		}
		log.Println("error occurs in in deletebooklogic for delete method", err.Error())
		return errors.New("internal server cloudinary error")
	}

	return nil
}

func (bm *bookModel) AddBookLogic(username string, description string, title string, bookfile *multipart.FileHeader) error {
	if err := bm.repo.AddBook(username, description, title, bookfile); err != nil {
		log.Println("errors occusr on add books")
		if strings.Contains(err.Error(), "upload image") {
			log.Println("error occurs in uploading image", err.Error())
			return errors.New("internal server error (cloudinary)")
		}

		return errors.New("bad request")
	}

	return nil
}
