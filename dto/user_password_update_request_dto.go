package dto

type UserPasswordUpdateRequestDTO struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
}
