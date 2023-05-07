package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/judoassistant/judoassistant-meta-service/entity"
	"github.com/pkg/errors"
)

type TournamentRepository interface {
	Create(entity *entity.TournamentEntity) error
	DeleteByID(id int64) error
	GetByID(id int64) (*entity.TournamentEntity, error)
	ListByDateGreaterThanEqual(minimumDate time.Time, limit int) ([]entity.TournamentEntity, error)
	ListByDateLessThan(maximumDate time.Time, limit int) ([]entity.TournamentEntity, error)
	ListByIDGreaterThan(after int64, count int) ([]entity.TournamentEntity, error)
	ListByOwner(ownerID int64) ([]entity.TournamentEntity, error)
	Update(entity *entity.TournamentEntity) error
}

type tournamentRepository struct {
	db *sqlx.DB
}

func NewTournamentRepository(db *sqlx.DB) TournamentRepository {
	return &tournamentRepository{db}
}

func (repository *tournamentRepository) Create(entity *entity.TournamentEntity) error {
	err := repository.db.Get(&entity.ID, "INSERT INTO tournaments (name, location, date, is_deleted, owner) VALUES ($1, $2, $3, $4, $5) RETURNING id", entity.Name, entity.Location, entity.Date, entity.IsDeleted, entity.Owner)
	if err != nil {
		return errors.Wrap(err, "unable to create tournament")
	}
	return nil
}

func (repository *tournamentRepository) Update(entity *entity.TournamentEntity) error {
	result, err := repository.db.Exec("UPDATE tournaments SET name = $2, location = $3, date = $4, is_deleted = $5, owner = $6 WHERE id = $1", entity.ID, entity.Name, entity.Location, entity.Date, entity.IsDeleted, entity.Owner)
	if err != nil {
		return errors.Wrap(err, "unable to update tournament")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "unable to get rowsAffected")
	}
	if rowsAffected == 0 {
		return errors.New("tournament does not exist")
	}
	return nil
}

func (repository *tournamentRepository) GetByID(id int64) (*entity.TournamentEntity, error) {
	tournament := entity.TournamentEntity{}
	err := repository.db.Get(&tournament, "SELECT * FROM tournaments WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tournament")
	}
	return &tournament, nil
}

func (repository *tournamentRepository) ListByOwner(ownerID int64) ([]entity.TournamentEntity, error) {
	tournaments := []entity.TournamentEntity{}
	err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE owner = $1", ownerID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list tournanents")
	}
	return tournaments, nil
}

func (repository *tournamentRepository) ListByIDGreaterThan(after int64, count int) ([]entity.TournamentEntity, error) {
	tournaments := []entity.TournamentEntity{}
	err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE id >= $1 AND is_deleted = 0 ORDER BY id LIMIT $2", after, count)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list tournanents")
	}
	return tournaments, nil
}

func (repository *tournamentRepository) ListByDateGreaterThanEqual(minimumDate time.Time, limit int) ([]entity.TournamentEntity, error) {
	tournaments := []entity.TournamentEntity{}
	err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE date >= $1 AND is_deleted = 0 ORDER BY date LIMIT $2", minimumDate, limit)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list tournanets")
	}
	return tournaments, nil
}

func (repository *tournamentRepository) ListByDateLessThan(maximumDate time.Time, limit int) ([]entity.TournamentEntity, error) {
	tournaments := []entity.TournamentEntity{}
	err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE date < $1 AND is_deleted = 0 ORDER BY date LIMIT $2", maximumDate, limit)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list tournaments")
	}
	return tournaments, nil
}

func (repository *tournamentRepository) DeleteByID(id int64) error {
	result, err := repository.db.Exec("UPDATE tournaments SET is_deleted = 1 WHERE id = $1 AND is_deleted = 0", id)
	if err != nil {
		return errors.Wrap(err, "unable to delete tournament")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "unable to get rowsAffected")
	}
	if rowsAffected == 0 {
		return errors.New("tournament does not exist")
	}
	return nil
}
