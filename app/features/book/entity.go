package book

import "github.com/labstack/echo"

type Core struct {
	Title     string
	BookImage string
	Status    bool
	Username  string
}

type Handler interface {
	GetAllBookHandler() echo.HandlerFunc
	DeLeteBookHandler() echo.HandlerFunc
}

type UseCase interface {
	GetAllBookLogic(offset int, limit int) ([]Core, error)
	DeleteBookLogic(username string, bookID uint) error
}

type Repository interface {
	GetAllBook(offset int, limit int) ([]Core, error)
	DeleteBook(username string, bookID uint) error
}
