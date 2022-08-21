package services

import (
	"time"

	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/entities"
	"github.com/judoassistant/judoassistant-meta-service/repositories"
)

type TournamentService struct {
	tournamentRepository *repositories.TournamentRepository
}

func NewTournamentService(tournamentRepository *repositories.TournamentRepository) *TournamentService {
	return &TournamentService{tournamentRepository}
}

func (service *TournamentService) GetPast(count int) ([]dto.TournamentResponseDTO, error) {
	today := time.Now()
	tournaments, err := service.tournamentRepository.GetByDateLessThan(today, 10) // TODO: find nice place to put constants
	if err != nil {
		return nil, err
	}

	return dto.MapTournamentResponseDTOs(tournaments), nil
}

func (service *TournamentService) GetUpcoming(count int) ([]dto.TournamentResponseDTO, error) {
	today := time.Now()
	tournaments, err := service.tournamentRepository.GetByDateGreaterThanEqual(today, 10) // TODO: find nice place to put constants
	if err != nil {
		return nil, err
	}

	return dto.MapTournamentResponseDTOs(tournaments), nil
}

func (service *TournamentService) Get(after int64, count int) ([]dto.TournamentResponseDTO, error) {
	tournamentEntities, err := service.tournamentRepository.GetByIdGreaterThan(after, count)
	if err != nil {
		return nil, err
	}

	tournamentDTOs := dto.MapTournamentResponseDTOs(tournamentEntities)
	return tournamentDTOs, nil
}

func (service *TournamentService) GetById(id int64) (*dto.TournamentResponseDTO, error) {
	tournamentEntity, err := service.tournamentRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	tournamentDTO := dto.MapTournamentResponseDTO(tournamentEntity)
	return &tournamentDTO, nil
}

func (service *TournamentService) Create(tournament *dto.TournamentCreationRequestDTO) (*dto.TournamentResponseDTO, error) {
	entity := entities.TournamentEntity{
		Name:     tournament.Name,
		Location: tournament.Location,
		Date:     tournament.Date,
	}

	if err := service.tournamentRepository.Create(&entity); err != nil {
		return nil, err
	}

	response := dto.MapTournamentResponseDTO(&entity)
	return &response, nil
}
