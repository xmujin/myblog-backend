package repository

import (
	"github.com/xmujin/myblog-backend/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(user *model.User) error {
	return u.db.Create(user).Error
}

func (u *UserRepository) GetUserById(id uint) (*model.User, error) {
	var user model.User
	err := u.db.First(user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) GetUserByName(name string) (*model.User, error) {
	var user model.User
	err := u.db.Where("name = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
