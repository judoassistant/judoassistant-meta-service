package service

import (
	"time"

	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/entity"
	"github.com/judoassistant/judoassistant-meta-service/repository"
)

type TournamentService interface {
	GetPast(count int) ([]dto.TournamentResponseDTO, error)
	GetUpcoming(count int) ([]dto.TournamentResponseDTO, error)
	Get(after int64, count int) ([]dto.TournamentResponseDTO, error)
	GetById(id int64) (*dto.TournamentResponseDTO, error)
	GetByOwner(ownerID int64) ([]dto.TournamentResponseDTO, error)
	Update(id int64, request *dto.TournamentUpdateRequestDTO) (*dto.TournamentResponseDTO, error)
	Create(user *dto.UserResponseDTO, tournament *dto.TournamentCreationRequestDTO) (*dto.TournamentResponseDTO, error)
}

type tournamentService struct {
	tournamentRepository *repository.TournamentRepository
}

func NewTournamentService(tournamentRepository *repository.TournamentRepository) TournamentService {
	return &tournamentService{tournamentRepository}
}

func (s *tournamentService) GetPast(count int) ([]dto.TournamentResponseDTO, error) {
	today := time.Now()
	tournaments, err := s.tournamentRepository.GetByDateLessThanAndNotDeleted(today, 10) // TODO: find nice place to put constants
	if err != nil {
		return nil, err
	}

	return dto.MapTournamentResponseDTOs(tournaments), nil
}

func (s *tournamentService) GetUpcoming(count int) ([]dto.TournamentResponseDTO, error) {
	today := time.Now()
	tournaments, err := s.tournamentRepository.GetByDateGreaterThanEqualAndNotDeleted(today, 10) // TODO: find nice place to put constants
	if err != nil {
		return nil, err
	}

	return dto.MapTournamentResponseDTOs(tournaments), nil
}

func (s *tournamentService) Get(after int64, count int) ([]dto.TournamentResponseDTO, error) {
	tournamentEntities, err := s.tournamentRepository.GetByIdGreaterThanAndNotDeleted(after, count)
	if err != nil {
		return nil, err
	}

	tournamentDTOs := dto.MapTournamentResponseDTOs(tournamentEntities)
	return tournamentDTOs, nil
}

func (s *tournamentService) GetById(id int64) (*dto.TournamentResponseDTO, error) {
	tournamentEntity, err := s.tournamentRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	tournamentDTO := dto.MapTournamentResponseDTO(tournamentEntity)
	return &tournamentDTO, nil
}

func (s *tournamentService) GetByOwner(ownerID int64) ([]dto.TournamentResponseDTO, error) {
	entities, err := s.tournamentRepository.GetByOwner(ownerID)
	if err != nil {
		return nil, err
	}

	tournamentDTOs := dto.MapTournamentResponseDTOs(entities)
	return tournamentDTOs, err
}

func (s *tournamentService) Update(id int64, request *dto.TournamentUpdateRequestDTO) (*dto.TournamentResponseDTO, error) {
	entity, err := s.tournamentRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	dto.MapToTournamentEntity(request, entity)

	if err := s.tournamentRepository.Update(entity); err != nil {
		return nil, err
	}

	response := dto.MapTournamentResponseDTO(entity)
	return &response, nil
}

func (s *tournamentService) Create(user *dto.UserResponseDTO, tournament *dto.TournamentCreationRequestDTO) (*dto.TournamentResponseDTO, error) {
	entity := entity.TournamentEntity{
		Name:      tournament.Name,
		Location:  tournament.Location,
		Date:      tournament.Date,
		Owner:     user.ID,
		IsDeleted: false,
	}

	if err := s.tournamentRepository.Create(&entity); err != nil {
		return nil, err
	}

	response := dto.MapTournamentResponseDTO(&entity)
	return &response, nil
}
