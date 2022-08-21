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

func (service *UserService) Authenticate(request *dto.UserAuthenticationRequestDTO) (*dto.UserResponseDTO, error) {
	userEntity, err := service.userRepository.GetByEmail(request.Email)

	if err != nil {
		return nil, err
	}

	if err := checkPasswordHash(request.Password, userEntity.PasswordHash); err != nil {
		return nil, err
	}

	user := dto.MapUserResponseDTO(userEntity)
	return &user, nil
}

func (service *UserService) Register(request *dto.UserRegistrationRequestDTO) (*dto.UserResponseDTO, error) {
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

	response := dto.MapUserResponseDTO(&userEntity)
	return &response, nil
}

func (service *UserService) Update(id int64, request *dto.UserUpdateRequestDTO) (*dto.UserResponseDTO, error) {
	userEntity, err := service.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	dto.MapToUserEntity(request, userEntity)

	if err := service.userRepository.Update(userEntity); err != nil {
		return nil, err
	}

	response := dto.MapUserResponseDTO(userEntity)
	return &response, nil
}

func (service *UserService) UpdatePassword(id int64, password string) (*dto.UserResponseDTO, error) {
	userEntity, err := service.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	if userEntity.PasswordHash, err = hashPassword(password); err != nil {
		return nil, err
	}

	if err := service.userRepository.Update(userEntity); err != nil {
		return nil, err
	}

	response := dto.MapUserResponseDTO(userEntity)
	return &response, nil
}

func (service *UserService) ExistsByEmail(email string) (bool, error) {
	return service.userRepository.ExistsByEmail(email)
}

func (service *UserService) GetById(id int64) (*dto.UserResponseDTO, error) {
	userEntity, err := service.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	response := dto.MapUserResponseDTO(userEntity)
	return &response, nil
}

func (service *UserService) GetAll() ([]dto.UserResponseDTO, error) {
	users, err := service.userRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return dto.MapUserResponseDTOs(users), nil
}

func hashPassword(password string) (string, error) {
	const cost = 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(bytes), err
}

func checkPasswordHash(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
