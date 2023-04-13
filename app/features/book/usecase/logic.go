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

func (bm *bookModel) GetBookByIdLogic(bookID uint) (book.Core, error) {
	result, err := bm.repo.GetBookByID(bookID)
	if err != nil {
		log.Println("errors occurs in getting book by id", err.Error())
		if strings.Contains(err.Error(), "finding") {
			return book.Core{}, errors.New("finding book error from gorm")
		}
		log.Println("errors occurs in getting book by id (bad request)", err.Error())
		return book.Core{}, errors.New("books doesnt exist")
	}

	return result, nil
}

func (bm *bookModel) UpdateBookLogic(bookID uint, title string, description string, bookImage *multipart.FileHeader) error {
	if err := bm.repo.UpdateBook(bookID, title, description, bookImage); err != nil {
		if strings.Contains("too much", err.Error()) {
			log.Println("bad request data value from user", err.Error())
			return errors.New("bad request")
		} else if strings.Contains("cannot", err.Error()) {
			log.Println("failed to connect with cloud images", err.Error())
			return errors.New("third party server down")
		}
		return err
	}

	return nil
}
