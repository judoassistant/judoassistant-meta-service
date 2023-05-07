package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/judoassistant/judoassistant-meta-service/entity"
	"github.com/pkg/errors"
)

type UserRepository interface {
	Create(entity *entity.UserEntity) error
	Update(entity *entity.UserEntity) error
	ExistsByEmail(email string) (bool, error)
	GetById(id int64) (*entity.UserEntity, error)
	GetByEmail(email string) (*entity.UserEntity, error)
	GetAll() ([]entity.UserEntity, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db}
}

func (repository *userRepository) Create(entity *entity.UserEntity) error {
	err := repository.db.Get(&entity.ID, "INSERT INTO users (first_name, last_name, email, password_hash, is_admin) VALUES ($1, $2, $3, $4, $5) RETURNING id", entity.FirstName, entity.LastName, entity.Email, entity.PasswordHash, entity.IsAdmin)
	if err != nil {
		return errors.Wrap(err, "unable to create user")
	}
	return nil
}

func (repository *userRepository) Update(entity *entity.UserEntity) error {
	_, err := repository.db.Exec("UPDATE users SET first_name = $2, last_name = $3, email = $4, password_hash = $5, is_admin = $6 WHERE id = $1", entity.ID, entity.FirstName, entity.LastName, entity.Email, entity.PasswordHash, entity.IsAdmin)
	if err != nil {
		return errors.Wrap(err, "unable to update user")
	}
	return nil
}

func (repository *userRepository) ExistsByEmail(email string) (bool, error) {
	var count int
	err := repository.db.Get(&count, "SELECT count(*) FROM users WHERE email = $1 LIMIT 1", email)
	if err != nil {
		return false, errors.Wrap(err, "unable to get user")
	}

	return count > 0, nil
}

func (repository *userRepository) GetById(id int64) (*entity.UserEntity, error) {
	user := entity.UserEntity{}
	err := repository.db.Get(&user, "SELECT * FROM users WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get user")
	}

	return &user, nil
}

func (repository *userRepository) GetByEmail(email string) (*entity.UserEntity, error) {
	user := entity.UserEntity{}
	err := repository.db.Get(&user, "SELECT * FROM users WHERE email = $1 LIMIT 1", email)
	if err != nil {
		return nil, errors.Wrap(err, "unable get get user")
	}
	return &user, nil
}

func (repository *userRepository) GetAll() ([]entity.UserEntity, error) {
	users := []entity.UserEntity{}
	err := repository.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, errors.Wrap(err, "unable get get users")
	}
	return users, nil
}
