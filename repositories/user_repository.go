package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type UserEntity struct {
	ID           int64  `db:"id"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
	IsAdmin      bool   `db:"is_admin"`
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (repository *UserRepository) Create(entity *UserEntity) error {
	log.Println("Hello WOrld!")
	return repository.db.Get(&entity.ID, "INSERT INTO users (first_name, last_name, email, password_hash, is_admin) VALUES ($1, $2, $3, $4, $5) RETURNING id", entity.FirstName, entity.LastName, entity.Email, entity.PasswordHash, entity.IsAdmin)
}

func (repository *UserRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := repository.db.Get(&count, "SELECT count(*) FROM users WHERE email = $1 LIMIT 1", email)
	return count > 0, err
}

func (repository *UserRepository) GetByEmail(email string) (*UserEntity, error) {
	user := UserEntity{}
	err := repository.db.Get(&user, "SELECT * FROM users WHERE email = $1 LIMIT 1", email)
	return &user, err
}

func (repository *UserRepository) GetAll() ([]UserEntity, error) {
	users := []UserEntity{}
	err := repository.db.Select(&users, "SELECT * FROM users")
	return users, err
}
