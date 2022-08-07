package dto

type UserAuthenticationRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
