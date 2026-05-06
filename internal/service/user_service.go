package service

import (
	"errors"

	"github.com/xmujin/myblog-backend/internal/model"
	"github.com/xmujin/myblog-backend/internal/repository"
	"github.com/xmujin/myblog-backend/pkg/auth"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Login(user *model.UserLoginDto) (string, error) {
	record, _ := u.userRepository.GetUserByName(user.Name)
	if record == nil {
		return "", errors.New("用户名或密码错误")
	}
	if record.Name != user.Name || record.Password != user.Password {
		return "", errors.New("用户名或密码错误")
	}
	tokens, err := auth.GenerateJWT(record.Id, record.Name, "admin")
	if err != nil {
		return "", err
	}
	return tokens, nil
}

func (u *UserService) Register(user *model.UserRegisterDto) error {
	record, _ := u.userRepository.GetUserByName(user.Name)
	if record != nil {
		return errors.New("用户已经存在")
	}
	newUser := &model.User{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}
	err := u.userRepository.CreateUser(newUser)
	if err != nil {
		return err
	}
	return nil
}
