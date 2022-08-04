package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type TournamentEntity struct {
	Id int64
}

type TournamentRepository struct {
	db *sqlx.DB
}

func NewTournamentRepository(db *sqlx.DB) *TournamentRepository {
	return &TournamentRepository{db}
}

func (repository *TournamentRepository) GetById(id int64) (*TournamentEntity, error) {
    tournament := TournamentEntity{}
    err := repository.db.Get(&tournament, "SELECT * FROM tournaments WHERE id = $1 LIMIT 1", id)
	return &tournament, err
}

func (repository *TournamentRepository) GetByDateGreaterThanEqual(minimumDate time.Time, limit uint) (*[]TournamentEntity, error) {
    tournaments := []TournamentEntity{}
    err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE date >= $1 LIMIT $2", minimumDate, limit)
	return &tournaments, err
}

func (repository *TournamentRepository) GetByDateLessThan(maximumDate time.Time, limit uint) (*[]TournamentEntity, error) {
    tournaments := []TournamentEntity{}
    err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE date < $1 LIMIT $2", maximumDate, limit)
	return &tournaments, err
}
