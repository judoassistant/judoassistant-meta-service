package services

import (
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (service *UserService) Authenticate(email string, password string) (dto.UserDTO, error) {
	user := dto.UserDTO{
		ID:      1,
		Email:   "svendcs@svendcs.com",
		IsAdmin: true,
	}

	return user, nil
}
