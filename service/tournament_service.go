package service

import (
	"time"

	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/entity"
	"github.com/judoassistant/judoassistant-meta-service/repository"
)

type TournamentService struct {
	tournamentRepository *repository.TournamentRepository
}

func NewTournamentService(tournamentRepository *repository.TournamentRepository) *TournamentService {
	return &TournamentService{tournamentRepository}
}

func (service *TournamentService) GetPast(count int) ([]dto.TournamentResponseDTO, error) {
	today := time.Now()
	tournaments, err := service.tournamentRepository.GetByDateLessThanAndNotDeleted(today, 10) // TODO: find nice place to put constants
	if err != nil {
		return nil, err
	}

	return dto.MapTournamentResponseDTOs(tournaments), nil
}

func (service *TournamentService) GetUpcoming(count int) ([]dto.TournamentResponseDTO, error) {
	today := time.Now()
	tournaments, err := service.tournamentRepository.GetByDateGreaterThanEqualAndNotDeleted(today, 10) // TODO: find nice place to put constants
	if err != nil {
		return nil, err
	}

	return dto.MapTournamentResponseDTOs(tournaments), nil
}

func (service *TournamentService) Get(after int64, count int) ([]dto.TournamentResponseDTO, error) {
	tournamentEntities, err := service.tournamentRepository.GetByIdGreaterThanAndNotDeleted(after, count)
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

func (service *TournamentService) GetByOwner(ownerID int64) ([]dto.TournamentResponseDTO, error) {
	entities, err := service.tournamentRepository.GetByOwner(ownerID)
	if err != nil {
		return nil, err
	}

	tournamentDTOs := dto.MapTournamentResponseDTOs(entities)
	return tournamentDTOs, err
}

func (service *TournamentService) Update(id int64, request *dto.TournamentUpdateRequestDTO) (*dto.TournamentResponseDTO, error) {
	entity, err := service.tournamentRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	dto.MapToTournamentEntity(request, entity)

	if err := service.tournamentRepository.Update(entity); err != nil {
		return nil, err
	}

	response := dto.MapTournamentResponseDTO(entity)
	return &response, nil
}

func (service *TournamentService) Create(user *dto.UserResponseDTO, tournament *dto.TournamentCreationRequestDTO) (*dto.TournamentResponseDTO, error) {
	entity := entity.TournamentEntity{
		Name:      tournament.Name,
		Location:  tournament.Location,
		Date:      tournament.Date,
		Owner:     user.ID,
		IsDeleted: false,
	}

	if err := service.tournamentRepository.Create(&entity); err != nil {
		return nil, err
	}

	response := dto.MapTournamentResponseDTO(&entity)
	return &response, nil
}
