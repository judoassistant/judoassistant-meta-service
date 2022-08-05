package dto

import "github.com/judoassistant/judoassistant-meta-service/repositories"

type UserDTO struct {
	ID      int64
	Email   string
	IsAdmin bool
}

func MapUserDTO(user *repositories.UserEntity) UserDTO {
	return UserDTO{
		ID:      user.ID,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
	}
}
