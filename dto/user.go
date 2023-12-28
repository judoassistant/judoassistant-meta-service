package dto
type UserAuthenticationRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserUpdateRequestDTO struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type UserRegistrationRequestDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsAdmin   bool   `json:"is_admin"`
}
type UserQueryDTO struct {
	ID int64 `uri:"id"`
}
type UserPasswordUpdateRequestDTO struct {
	Password string `json:"password"`
}
type UserResponseDTO struct {
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
}
