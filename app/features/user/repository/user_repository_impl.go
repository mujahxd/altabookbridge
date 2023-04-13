package repository

import (
	"github.com/mujahxd/altabookbridge/app/features/user"
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

func (m *model) Update(user user.User) (user.User, error) {
	err := m.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
