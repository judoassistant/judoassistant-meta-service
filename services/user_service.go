package services

import (
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/entities"
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

	userEntity := entities.UserEntity{
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

func (service *UserService) Update(request *dto.UserUpdateRequestDTO) (*dto.UserDTO, error) {
	return nil, nil // TODO: Implement
}

func (service *UserService) UpdatePassword(request *dto.UserPasswordUpdateRequestDTO) (*dto.UserDTO, error) {
	userEntity, err := service.userRepository.GetById(request.ID)
	if err != nil {
		return nil, err
	}

	if userEntity.PasswordHash, err = hashPassword(request.Password); err != nil {
		return nil, err
	}

	if err := service.userRepository.Update(userEntity); err != nil {
		return nil, err
	}

	response := dto.MapUserDTO(userEntity)
	return &response, nil
}

func (service *UserService) ExistsByEmail(email string) (bool, error) {
	return service.userRepository.ExistsByEmail(email)
}

func (service *UserService) GetById(id int64) (*dto.UserDTO, error) {
	userEntity, err := service.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	response := dto.MapUserDTO(userEntity)
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
