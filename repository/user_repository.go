package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/judoassistant/judoassistant-meta-service/entity"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (repository *UserRepository) Create(entity *entity.UserEntity) error {
	return repository.db.Get(&entity.ID, "INSERT INTO users (first_name, last_name, email, password_hash, is_admin) VALUES ($1, $2, $3, $4, $5) RETURNING id", entity.FirstName, entity.LastName, entity.Email, entity.PasswordHash, entity.IsAdmin)
}

func (repository *UserRepository) Update(entity *entity.UserEntity) error {
	// TODO: Implement
	return nil
}

func (repository *UserRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := repository.db.Get(&count, "SELECT count(*) FROM users WHERE email = $1 LIMIT 1", email)
	return count > 0, err
}

func (repository *UserRepository) GetById(id int64) (*entity.UserEntity, error) {
	user := entity.UserEntity{}
	err := repository.db.Get(&user, "SELECT * FROM users WHERE id = $1 LIMIT 1", id)
	return &user, err
}

func (repository *UserRepository) GetByEmail(email string) (*entity.UserEntity, error) {
	user := entity.UserEntity{}
	err := repository.db.Get(&user, "SELECT * FROM users WHERE email = $1 LIMIT 1", email)
	return &user, err
}

func (repository *UserRepository) GetAll() ([]entity.UserEntity, error) {
	users := []entity.UserEntity{}
	err := repository.db.Select(&users, "SELECT * FROM users")
	return users, err
}
