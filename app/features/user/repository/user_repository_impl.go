package repository

import (
	"errors"
	"log"
	"mime/multipart"

	"github.com/mujahxd/altabookbridge/app/features/user"
	"github.com/mujahxd/altabookbridge/helper"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func NewModel(db *gorm.DB) *model {
	return &model{db}
}

func (m *model) Save(user user.User) (user.User, error) {
	err := m.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (m *model) FindByUsername(username string) (user.User, error) {
	var user user.User
	err := m.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (m *model) FindByID(ID int) (user.User, error) {
	var user user.User
	err := m.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (m *model) Update(username string, name string, password string, avatar *multipart.FileHeader) error {
	var updatedUser user.User
	avatarurl, err := helper.Upload(avatar)
	if err != nil {
		log.Println("errors from calling uploader", err.Error())
		return errors.New("cannot upload image to server")
	}

	updatedUser.Name = name
	updatedUser.Password = password
	updatedUser.AvatarFileName = avatarurl
	err = m.db.Model(&user.User{}).Where("username = ?", username).Updates(updatedUser).Error
	if err != nil {
		return err
	}
	return nil

}

func (m *model) Delete(username string) error {
	var user user.User
	err := m.db.Where("username = ?", username).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
