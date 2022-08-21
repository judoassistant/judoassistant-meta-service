package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/judoassistant/judoassistant-meta-service/entities"
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

func (repository *TournamentRepository) Create(entity *entities.TournamentEntity) error {
	return repository.db.Get(&entity.ID, "INSERT INTO tournaments (name, location, date) VALUES ($1, $2, $3) RETURNING id", entity.Name, entity.Location, entity.Date)
}

func (repository *TournamentRepository) GetById(id int64) (*entities.TournamentEntity, error) {
	tournament := entities.TournamentEntity{}
	err := repository.db.Get(&tournament, "SELECT * FROM tournaments WHERE id = $1 LIMIT 1", id)
	return &tournament, err
}

func (repository *TournamentRepository) GetByIdGreaterThan(after int64, count int) ([]entities.TournamentEntity, error) {
	tournaments := []entities.TournamentEntity{}
	err := repository.db.Get(&tournaments, "SELECT * FROM tournaments WHERE id >= $1 LIMIT $2 ORDER BY id", after, count)
	return tournaments, err
}

func (repository *TournamentRepository) GetByDateGreaterThanEqual(minimumDate time.Time, limit int) ([]entities.TournamentEntity, error) {
	tournaments := []entities.TournamentEntity{}
	err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE date >= $1 LIMIT $2 ORDER BY date", minimumDate, limit)
	return tournaments, err
}

func (repository *TournamentRepository) GetByDateLessThan(maximumDate time.Time, limit int) ([]entities.TournamentEntity, error) {
	tournaments := []entities.TournamentEntity{}
	err := repository.db.Select(&tournaments, "SELECT * FROM tournaments WHERE date < $1 LIMIT $2 ORDER BY date", maximumDate, limit)
	return tournaments, err
}
