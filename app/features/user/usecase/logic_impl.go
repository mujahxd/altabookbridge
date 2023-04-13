package usecase

import (
	"errors"

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
func (l *logic) Login(input data.LoginInput) (user.User, error) {
	username := input.Username
	password := input.Password

	user, err := l.repo.FindByUsername(username)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that username")
	}

	err = helper.VerifyPassword(user.Password, password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (l *logic) SaveAvatar(ID int, fileLocation string) (user.User, error) {
	user, err := l.repo.FindByID(ID)
	if err != nil {
		return user, err
	}
	user.AvatarFileName = fileLocation
	updatedUser, err := l.repo.Update(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}
