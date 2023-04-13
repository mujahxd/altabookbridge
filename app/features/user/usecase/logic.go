package usecase

import (
	"github.com/mujahxd/altabookbridge/app/features/user"
	"github.com/mujahxd/altabookbridge/app/features/user/data"
)

type UseCase interface {
	RegisterUser(input data.RegisterUserInput) (user.User, error)
	Login(input data.LoginInput) (user.User, error)
	GetUserByUsername(username string) (user.User, error)
	DeleteUser(username string) error
}
