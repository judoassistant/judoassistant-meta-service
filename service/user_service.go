package service

import (
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/entity"
	"github.com/judoassistant/judoassistant-meta-service/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Authenticate(request *dto.UserAuthenticationRequestDTO) (*dto.UserResponseDTO, error)
	Register(request *dto.UserRegistrationRequestDTO) (*dto.UserResponseDTO, error)
	Update(id int64, request *dto.UserUpdateRequestDTO) (*dto.UserResponseDTO, error)
	UpdatePassword(id int64, password string) (*dto.UserResponseDTO, error)
	ExistsByEmail(email string) (bool, error)
	GetById(id int64) (*dto.UserResponseDTO, error)
	GetAll() ([]dto.UserResponseDTO, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Authenticate(request *dto.UserAuthenticationRequestDTO) (*dto.UserResponseDTO, error) {
	userEntity, err := s.userRepository.GetByEmail(request.Email)

	if err != nil {
		return nil, err
	}

	if err := checkPasswordHash(request.Password, userEntity.PasswordHash); err != nil {
		return nil, err
	}

	user := dto.MapUserResponseDTO(userEntity)
	return &user, nil
}

func (s *userService) Register(request *dto.UserRegistrationRequestDTO) (*dto.UserResponseDTO, error) {
	passwordHash, err := hashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	userEntity := entity.UserEntity{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		Email:        request.Email,
		PasswordHash: passwordHash,
		IsAdmin:      false,
	}

	if err := s.userRepository.Create(&userEntity); err != nil {
		return nil, err
	}

	response := dto.MapUserResponseDTO(&userEntity)
	return &response, nil
}

func (s *userService) Update(id int64, request *dto.UserUpdateRequestDTO) (*dto.UserResponseDTO, error) {
	userEntity, err := s.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	dto.MapToUserEntity(request, userEntity)

	if err := s.userRepository.Update(userEntity); err != nil {
		return nil, err
	}

	response := dto.MapUserResponseDTO(userEntity)
	return &response, nil
}

func (s *userService) UpdatePassword(id int64, password string) (*dto.UserResponseDTO, error) {
	userEntity, err := s.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	if userEntity.PasswordHash, err = hashPassword(password); err != nil {
		return nil, err
	}

	if err := s.userRepository.Update(userEntity); err != nil {
		return nil, err
	}

	response := dto.MapUserResponseDTO(userEntity)
	return &response, nil
}

func (s *userService) ExistsByEmail(email string) (bool, error) {
	return s.userRepository.ExistsByEmail(email)
}

func (s *userService) GetById(id int64) (*dto.UserResponseDTO, error) {
	userEntity, err := s.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	response := dto.MapUserResponseDTO(userEntity)
	return &response, nil
}

func (s *userService) GetAll() ([]dto.UserResponseDTO, error) {
	users, err := s.userRepository.GetAll()

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
