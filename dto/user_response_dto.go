package dto

import "github.com/judoassistant/judoassistant-meta-service/entity"

type UserResponseDTO struct {
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func MapUserResponseDTO(user *entity.UserEntity) UserResponseDTO {
	return UserResponseDTO{
		ID:      user.ID,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
	}
}

func MapUserResponseDTOs(users []entity.UserEntity) []UserResponseDTO {
	result := make([]UserResponseDTO, len(users))

	for key, value := range users {
		result[key] = MapUserResponseDTO(&value)
	}

	return result
}
