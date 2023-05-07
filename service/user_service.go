package service

import (
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/entity"
	"github.com/judoassistant/judoassistant-meta-service/mappers"
	"github.com/judoassistant/judoassistant-meta-service/repository"
	"github.com/pkg/errors"
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
		return nil, errors.Wrap(err, "unable to get user")
	}

	if ok := checkPasswordHash(request.Password, userEntity.PasswordHash); !ok {
		return nil, errors.New("incorrect password")
	}

	user := mappers.UserToResponseDTO(userEntity)
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
		return nil, errors.Wrap(err, "unable to create user")
	}

	response := mappers.UserToResponseDTO(&userEntity)
	return &response, nil
}

func (s *userService) Update(id int64, request *dto.UserUpdateRequestDTO) (*dto.UserResponseDTO, error) {
	userEntity, err := s.userRepository.GetById(id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to update user")
	}

	mappers.UserFromUpdateRequestDTO(request, userEntity)

	if err := s.userRepository.Update(userEntity); err != nil {
		return nil, errors.Wrap(err, "unable to update user")
	}

	response := mappers.UserToResponseDTO(userEntity)
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
		return nil, errors.Wrap(err, "unable to update user")
	}

	response := mappers.UserToResponseDTO(userEntity)
	return &response, nil
}

func (s *userService) ExistsByEmail(email string) (bool, error) {
	return s.userRepository.ExistsByEmail(email)
}

func (s *userService) GetById(id int64) (*dto.UserResponseDTO, error) {
	userEntity, err := s.userRepository.GetById(id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get user")
	}

	response := mappers.UserToResponseDTO(userEntity)
	return &response, nil
}

func (s *userService) GetAll() ([]dto.UserResponseDTO, error) {
	users, err := s.userRepository.GetAll()

	if err != nil {
		return nil, errors.Wrap(err, "unable to list users")
	}

	return mappers.UserToResponseDTOs(users), nil
}

func hashPassword(password string) (string, error) {
	const cost = 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return "", errors.Wrap(err, "unable to hash password")
	}

	return string(bytes), nil
}

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}
