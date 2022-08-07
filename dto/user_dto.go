package dto

import "github.com/judoassistant/judoassistant-meta-service/repositories"

type UserDTO struct {
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func MapUserDTO(user *repositories.UserEntity) UserDTO {
	return UserDTO{
		ID:      user.ID,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
	}
}

func MapUserDTOs(users []repositories.UserEntity) []UserDTO {
	result := make([]UserDTO, len(users))

	for key, value := range users {
		result[key] = MapUserDTO(&value)
	}

	return result
}
