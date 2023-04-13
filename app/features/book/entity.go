package book

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	Title       string `json:"title"`
	BookImage   string `json:"book_image"`
	Status      bool   `json:"status"`
	Username    string `json:"username"`
	Description string `json:"description"`
}

type Handler interface {
	GetAllBookHandler() echo.HandlerFunc
	DeLeteBookHandler() echo.HandlerFunc
	AddBookHandler() echo.HandlerFunc
	GetBookByIDHandler() echo.HandlerFunc
}

type UseCase interface {
	GetAllBookLogic(offset int, limit int) ([]Core, error)
	DeleteBookLogic(username string, bookID uint) error
	AddBookLogic(username string, description string, title string, bookFile *multipart.FileHeader) error
	GetBookByIdLogic(bookID uint) (Core, error)
}

type Repository interface {
	GetAllBook(offset int, limit int) ([]Core, error)
	DeleteBook(username string, bookID uint) error
	AddBook(username string, description string, title string, bookFile *multipart.FileHeader) error
	GetBookByID(bookID uint) (Core, error)
}
