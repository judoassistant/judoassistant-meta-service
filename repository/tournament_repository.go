package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/judoassistant/judoassistant-meta-service/entity"
	"github.com/judoassistant/judoassistant-meta-service/errors"
)

type TournamentRepository interface {
	Create(entity *entity.TournamentEntity) error
	DeleteByShortName(shortName string) error
	GetByShortName(shortName string) (*entity.TournamentEntity, error)
	ListByDateGreaterThanEqual(minimumDate time.Time, limit int) ([]entity.TournamentEntity, error)
	ListByDateLessThan(maximumDate time.Time, limit int) ([]entity.TournamentEntity, error)
	ListByShortNameGreaterThan(shortName string, limit int) ([]entity.TournamentEntity, error)
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
	err := repository.db.Get(&entity.ID, "INSERT INTO tournaments (name, location, date, owner, is_deleted, short_name) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", entity.Name, entity.Location, entity.Date, entity.Owner, entity.IsDeleted, entity.ShortName)
	if err != nil {
		return errors.WrapWithCode(err, "unable to create tournament", errorCodeFromDatabaseError(err))
	}
	return nil
}

func (repository *tournamentRepository) Update(entity *entity.TournamentEntity) error {
	result, err := repository.db.Exec("UPDATE tournaments SET name = $2, location = $3, date = $4, owner = $5, short_name = $6 WHERE id = $1", entity.ID, entity.Name, entity.Location, entity.Date, entity.Owner, entity.ShortName)
	if err != nil {
		return errors.WrapWithCode(err, "unable to update tournament", errorCodeFromDatabaseError(err))
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.WrapWithCode(err, "unable to get rowsAffected", errors.CodeInternal)
	}
	if rowsAffected == 0 {
		return errors.New("tournament does not exist", errors.CodeNotFound)
	}
	return nil
}

func (repository *tournamentRepository) GetByShortName(shortName string) (*entity.TournamentEntity, error) {
	tournament := entity.TournamentEntity{}
	err := repository.db.Get(&tournament, "SELECT * FROM tournaments WHERE short_name = $1 AND is_deleted = 0 LIMIT 1", shortName)
	if err != nil {
		if errCode := errorCodeFromDatabaseError(err); errCode == errors.CodeNotFound {
			return nil, errors.New("tournament does not exist", errCode)
		} else {
			return nil, errors.WrapWithCode(err, "unable to get tournament", errCode)
		}
	}
	return &tournament, nil
}

func (repository *tournamentRepository) ListByOwner(ownerID int64) ([]entity.TournamentEntity, error) {
	tournaments := []entity.TournamentEntity{}
	err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE owner = $1 AND is_deleted = 0", ownerID)
	if err != nil {
		return nil, errors.WrapWithCode(err, "unable to list tournanents", errorCodeFromDatabaseError(err))
	}
	return tournaments, nil
}

func (repository *tournamentRepository) ListByShortNameGreaterThan(after string, limit int) ([]entity.TournamentEntity, error) {
	tournaments := []entity.TournamentEntity{}
	err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE short_name > $1 AND is_deleted = 0 ORDER BY short_name LIMIT $2", after, limit)
	if err != nil {
		return nil, errors.WrapWithCode(err, "unable to list tournanents", errorCodeFromDatabaseError(err))
	}
	return tournaments, nil
}

func (repository *tournamentRepository) ListByDateGreaterThanEqual(minimumDate time.Time, limit int) ([]entity.TournamentEntity, error) {
	tournaments := []entity.TournamentEntity{}
	err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE date >= $1 AND is_deleted = 0 ORDER BY date LIMIT $2", minimumDate, limit)
	if err != nil {
		return nil, errors.WrapWithCode(err, "unable to list tournanets", errorCodeFromDatabaseError(err))
	}
	return tournaments, nil
}

func (repository *tournamentRepository) ListByDateLessThan(maximumDate time.Time, limit int) ([]entity.TournamentEntity, error) {
	tournaments := []entity.TournamentEntity{}
	err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE date < $1 AND is_deleted = 0 ORDER BY date LIMIT $2", maximumDate, limit)
	if err != nil {
		return nil, errors.WrapWithCode(err, "unable to list tournaments", errorCodeFromDatabaseError(err))
	}
	return tournaments, nil
}

func (repository *tournamentRepository) DeleteByShortName(shortName string) error {
	result, err := repository.db.Exec("UPDATE tournaments SET is_deleted = 1 WHERE short_name = $1 AND is_deleted = 0", shortName)
	if err != nil {
		return errors.WrapWithCode(err, "unable to delete tournament", errorCodeFromDatabaseError(err))
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.WrapWithCode(err, "unable to get rowsAffected", errors.CodeInternal)
	}
	if rowsAffected == 0 {
		return errors.New("tournament does not exist", errors.CodeNotFound)
	}
	return nil
}
