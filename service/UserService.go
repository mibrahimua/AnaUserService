package service

import (
	"AnaUserService/model"
	"AnaUserService/repository"
	"errors"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (us *UserService) GetUserByID(userID int) (*model.User, error) {
	return us.userRepository.GetUserByID(userID)
}

func (us *UserService) Login(username, password string) (*model.User, error) {
	user, err := us.userRepository.GetUserByEmailOrPhone(username)

	if err != nil {
		return nil, errors.New("Invalid username or password")
	}

	if user.Password != password {
		return nil, errors.New("Invalid username or password")
	}

	return user, nil
}
