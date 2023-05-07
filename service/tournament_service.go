package service

import (
	"github.com/benbjohnson/clock"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/entity"
	"github.com/judoassistant/judoassistant-meta-service/mappers"
	"github.com/judoassistant/judoassistant-meta-service/repository"
	"github.com/pkg/errors"
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
	tournamentRepository repository.TournamentRepository
	clock                clock.Clock
}

func NewTournamentService(tournamentRepository repository.TournamentRepository, clock clock.Clock) TournamentService {
	return &tournamentService{
		tournamentRepository: tournamentRepository,
		clock:                clock,
	}
}

func (s *tournamentService) GetPast(count int) ([]dto.TournamentResponseDTO, error) {
	today := s.clock.Now()
	tournaments, err := s.tournamentRepository.GetByDateLessThanAndNotDeleted(today, 10) // TODO: find nice place to put constants
	if err != nil {
		return nil, errors.Wrap(err, "unable to list tournaments")
	}

	return mappers.TournamentToResponseDTOs(tournaments), nil
}

func (s *tournamentService) GetUpcoming(count int) ([]dto.TournamentResponseDTO, error) {
	today := s.clock.Now()
	tournaments, err := s.tournamentRepository.GetByDateGreaterThanEqualAndNotDeleted(today, 10) // TODO: find nice place to put constants
	if err != nil {
		return nil, errors.Wrap(err, "unable to list tournaments")
	}

	return mappers.TournamentToResponseDTOs(tournaments), nil
}

func (s *tournamentService) Get(after int64, count int) ([]dto.TournamentResponseDTO, error) {
	tournaments, err := s.tournamentRepository.GetByIdGreaterThanAndNotDeleted(after, count)
	if err != nil {
		return nil, err
	}

	return mappers.TournamentToResponseDTOs(tournaments), nil
}

func (s *tournamentService) GetById(id int64) (*dto.TournamentResponseDTO, error) {
	tournament, err := s.tournamentRepository.GetById(id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tournament")
	}

	response := mappers.TournamentToResponseDTO(tournament)
	return &response, nil
}

func (s *tournamentService) GetByOwner(ownerID int64) ([]dto.TournamentResponseDTO, error) {
	tournaments, err := s.tournamentRepository.GetByOwner(ownerID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tournament")
	}

	return mappers.TournamentToResponseDTOs(tournaments), nil
}

func (s *tournamentService) Update(id int64, request *dto.TournamentUpdateRequestDTO) (*dto.TournamentResponseDTO, error) {
	tournament, err := s.tournamentRepository.GetById(id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tournament")
	}

	mappers.TournamentFromUpdateRequestDTO(request, tournament)

	if err := s.tournamentRepository.Update(tournament); err != nil {
		return nil, errors.Wrap(err, "unable to update tournament")
	}

	response := mappers.TournamentToResponseDTO(tournament)
	return &response, nil
}

func (s *tournamentService) Create(user *dto.UserResponseDTO, request *dto.TournamentCreationRequestDTO) (*dto.TournamentResponseDTO, error) {
	tournament := &entity.TournamentEntity{
		Name:      request.Name,
		Location:  request.Location,
		Date:      request.Date,
		Owner:     user.ID,
		IsDeleted: false,
	}

	if err := s.tournamentRepository.Create(tournament); err != nil {
		return nil, errors.Wrap(err, "unable to create tournament")
	}

	response := mappers.TournamentToResponseDTO(tournament)
	return &response, nil
}
