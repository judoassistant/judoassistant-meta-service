package repositories

import (
	"github.com/jmoiron/sqlx"
)

type UserEntity struct {
    Email string `db:"email"`
	PasswordHash string `db:"password_hash"`
    IsAdmin      bool `db:"is_admin"`
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (repository *UserRepository) GetByEmail(email string) (*UserEntity, error) {
    user := UserEntity{}
    err := repository.db.Get(&user, "SELECT * FROM users WHERE email = $1 LIMIT 1", email)
	return &user, err
}

func (repository *UserRepository) GetAll() (*[]UserEntity, error) {
    users := []UserEntity{}
    err := repository.db.Select(&users, "SELECT * FROM users")
	return &users, err
}
