package dto

import "github.com/judoassistant/judoassistant-meta-service/entity"

type UserUpdateRequestDTO struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func MapToUserEntity(dto *UserUpdateRequestDTO, entity *entity.UserEntity) {
	entity.Email = dto.Email
	entity.FirstName = dto.FirstName
	entity.LastName = dto.LastName
}
