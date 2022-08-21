package dto

import "github.com/judoassistant/judoassistant-meta-service/entities"

type UserResponseDTO struct {
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func MapUserDTO(user *entities.UserEntity) UserResponseDTO {
	return UserResponseDTO{
		ID:      user.ID,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
	}
}

func MapUserDTOs(users []entities.UserEntity) []UserResponseDTO {
	result := make([]UserResponseDTO, len(users))

	for key, value := range users {
		result[key] = MapUserDTO(&value)
	}

	return result
}
