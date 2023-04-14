package usecase

import (
	"errors"
	"log"
	"mime/multipart"
	"strings"

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

	// Create new user
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
	loginUser, err := l.repo.FindByUsername(input.Username)
	if err != nil {
		log.Println(err)
		return user.User{}, err
	}

	if loginUser.ID == 0 {
		return user.User{}, errors.New("no user found with that username")
	}

	if err := helper.VerifyPassword(loginUser.Password, input.Password); err != nil {
		log.Println(err)
		return user.User{}, err
	}

	return loginUser, nil
}

func (l *logic) GetUserByUsername(username string) (user.User, error) {
	user, err := l.repo.FindByUsername(username)
	if err != nil {
		return user, err
	}

	// cek user
	if user.ID == 0 {
		return user, errors.New("no user found on with that username")
	}

	return user, nil
}

func (l *logic) DeleteUser(username string) error {
	err := l.repo.Delete(username)
	if err != nil {
		return err
	}
	return nil
}

func (l *logic) UpdateUser(username string, name string, password string, avatar *multipart.FileHeader) error {

	err := l.repo.Update(username, name, password, avatar)
	if err != nil {
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
