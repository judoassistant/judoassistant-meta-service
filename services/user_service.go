package services

import "github.com/judoassistant/judoassistant-meta-service/repositories"

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}
