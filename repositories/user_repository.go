package repositories

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type UserEntity struct {
	Email        string
	PasswordHash string
	IsAdmin      bool
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (t *UserRepository) GetByEmail(email string) (*UserEntity, error) {
	return nil, errors.New("Blah")
}

func (t *UserRepository) GetByAll() (*[]UserEntity, error) {
	return nil, errors.New("Blah")
}
