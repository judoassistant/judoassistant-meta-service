package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/judoassistant/judoassistant-meta-service/entity"
	"github.com/pkg/errors"
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

func (repository *TournamentRepository) Create(entity *entity.TournamentEntity) error {
	err := repository.db.Get(&entity.ID, "INSERT INTO tournaments (name, location, date, is_deleted, owner) VALUES ($1, $2, $3, $4, $5) RETURNING id", entity.Name, entity.Location, entity.Date, entity.IsDeleted, entity.Owner)
	if err != nil {
		return errors.Wrap(err, "unable to create tournament")
	}
	return nil
}

func (repository *TournamentRepository) Update(entity *entity.TournamentEntity) error {
	_, err := repository.db.Exec("UPDATE tournaments SET name = $2, location = $3, date = $4, is_deleted = $5, owner = $6 WHERE id = $1", entity.ID, entity.Name, entity.Location, entity.Date, entity.IsDeleted, entity.Owner)
	if err != nil {
		return errors.Wrap(err, "unable to update tournament")
	}
	return nil
}

func (repository *TournamentRepository) GetById(id int64) (*entity.TournamentEntity, error) {
	tournament := entity.TournamentEntity{}
	err := repository.db.Get(&tournament, "SELECT * FROM tournaments WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tournament")
	}
	return &tournament, nil
}

func (repository *TournamentRepository) GetByOwner(ownerID int64) ([]entity.TournamentEntity, error) {
	tournaments := []entity.TournamentEntity{}
	err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE owner = $1", ownerID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tourmanent")
	}
	return tournaments, nil
}

func (repository *TournamentRepository) GetByIdGreaterThanAndNotDeleted(after int64, count int) ([]entity.TournamentEntity, error) {
	tournaments := []entity.TournamentEntity{}
	err := repository.db.Get(&tournaments, "SELECT * FROM tournaments WHERE id >= $1 AND is_deleted = 0 LIMIT $2 ORDER BY id", after, count)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tourmanents")
	}
	return tournaments, nil
}

func (repository *TournamentRepository) GetByDateGreaterThanEqualAndNotDeleted(minimumDate time.Time, limit int) ([]entity.TournamentEntity, error) {
	tournaments := []entity.TournamentEntity{}
	err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE date >= $1 AND is_deleted = 0 LIMIT $2 ORDER BY date", minimumDate, limit)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tourmanets")
	}
	return tournaments, nil
}

func (repository *TournamentRepository) GetByDateLessThanAndNotDeleted(maximumDate time.Time, limit int) ([]entity.TournamentEntity, error) {
	tournaments := []entity.TournamentEntity{}
	err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE date < $1 AND is_deleted = 0 LIMIT $2 ORDER BY date", maximumDate, limit)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tournaments")
	}
	return tournaments, nil
}
