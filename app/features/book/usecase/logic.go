package usecase

import (
	"errors"
	"log"
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
		return errors.New("Internal server error")
	}

	return nil
}
