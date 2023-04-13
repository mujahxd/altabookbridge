package usecase

import (
	"github.com/mujahxd/altabookbridge/app/features/user/data"
)

type UseCase interface {
	RegisterUser(input data.RegisterUserInput) error
}
