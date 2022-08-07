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

func (service *UserService) Authenticate(request *dto.UserAuthenticationRequestDTO) (*dto.UserDTO, error) {
	userEntity, err := service.userRepository.GetByEmail(request.Email)

	if err != nil {
		return nil, err
	}

	if err := checkPasswordHash(request.Password, userEntity.PasswordHash); err != nil {
		return nil, err
	}

	user := dto.MapUserDTO(userEntity)
	return &user, nil
}

func (service *UserService) Register(request *dto.UserRegistrationRequestDTO) (*dto.UserDTO, error) {
	passwordHash, err := hashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	userEntity := repositories.UserEntity{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		Email:        request.Email,
		PasswordHash: passwordHash,
		IsAdmin:      false,
	}

	if err := service.userRepository.Create(&userEntity); err != nil {
		return nil, err
	}

	response := dto.MapUserDTO(&userEntity)
	return &response, nil
}

func (service *UserService) GetAll() ([]dto.UserDTO, error) {
	users, err := service.userRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return dto.MapUserDTOs(users), nil
}

func hashPassword(password string) (string, error) {
	const cost = 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(bytes), err
}

func checkPasswordHash(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
