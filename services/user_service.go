package services

import (
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (service *UserService) Authenticate(email string, password string) (*dto.UserDTO, error) {
	userEntity, err := service.userRepository.GetByEmail(email)

	if err != nil {
		return nil, err
	}

	if err := checkPasswordHash(password, userEntity.PasswordHash); err != nil {
		return nil, err
	}

	user := dto.MapUserDTO(userEntity)
	return &user, nil
}

func hashPassword(password string) (string, error) {
	const cost = 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(bytes), err
}

func checkPasswordHash(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
