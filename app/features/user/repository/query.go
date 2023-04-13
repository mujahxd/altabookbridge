package repository

import (
	"github.com/mujahxd/altabookbridge/app/features/user"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func NewModel(db *gorm.DB) user.Repository {
	return &model{db}
}

func (m *model) Save(user user.Core) (user.Core, error) {
	err := m.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (m *model) FindByUsername(username string) (user.Core, error) {
	var user user.Core
	err := m.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
