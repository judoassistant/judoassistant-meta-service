package mappers

import (
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/entity"
)

func UserToResponseDTO(user *entity.UserEntity) dto.UserResponseDTO {
	return dto.UserResponseDTO{
		ID:      user.ID,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
	}
}

func UserToResponseDTOs(users []entity.UserEntity) []dto.UserResponseDTO {
	result := make([]dto.UserResponseDTO, len(users))

	for key, value := range users {
		result[key] = UserToResponseDTO(&value)
	}

	return result
}

func UserFromUpdateRequestDTO(dto *dto.UserUpdateRequestDTO, entity *entity.UserEntity) {
	entity.Email = dto.Email
	entity.FirstName = dto.FirstName
	entity.LastName = dto.LastName
}
