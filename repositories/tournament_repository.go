package repositories

import (
	"errors"

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

func (t *TournamentRepository) GetById(id int64) (*TournamentEntity, error) {
	return nil, errors.New("Blah")
}
