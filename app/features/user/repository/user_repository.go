package repository

import (
	"mime/multipart"

	"github.com/mujahxd/altabookbridge/app/features/user"
)

type Repository interface {
	Save(user user.User) (user.User, error)
	FindByUsername(username string) (user.User, error)
	FindByID(ID int) (user.User, error)
	Delete(username string) error
	Update(username string, name string, password string, avatar *multipart.FileHeader) error
}
