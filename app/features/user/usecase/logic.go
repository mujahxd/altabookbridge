package usecase

import (
	"github.com/mujahxd/altabookbridge/app/features/user"
	"github.com/mujahxd/altabookbridge/app/features/user/data"
	"github.com/mujahxd/altabookbridge/helper"
)

type logic struct {
	repo user.Repository
}

func NewLogic(repo user.Repository) user.UseCase {
	return &logic{repo}
}

func (l *logic) RegisterUser(input data.RegisterUserInput) (user.Core, error) {
	user := user.Core{}
	user.Name = input.Name
	user.Username = input.Username
	passwordHash, err := helper.HashPassword(input.Password)
	if err != nil {
		return user, err
	}
	user.Password = passwordHash
	newUser, err := l.repo.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil

}
