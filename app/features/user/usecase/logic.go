package usecase

import (
	"github.com/mujahxd/altabookbridge/app/features/user"
	"github.com/mujahxd/altabookbridge/app/features/user/data"
)

type UseCase interface {
	RegisterUser(input data.RegisterUserInput) (user.User, error)
	Login(input data.LoginInput) (user.User, error)
	SaveAvatar(ID int, fileLocation string) (user.User, error)
}
