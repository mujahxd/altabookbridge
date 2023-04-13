package usecase

import (
	"github.com/mujahxd/altabookbridge/app/features/user"
	"github.com/mujahxd/altabookbridge/app/features/user/data"
	"github.com/mujahxd/altabookbridge/app/features/user/repository"
	"github.com/mujahxd/altabookbridge/helper"
)

type logic struct {
	repo repository.Repository
}

func NewLogic(repo repository.Repository) *logic {
	return &logic{repo}
}

func (l *logic) RegisterUser(input data.RegisterUserInput) (user.User, error) {
	user := user.User{}
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
