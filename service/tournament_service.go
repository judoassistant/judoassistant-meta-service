package service

import (
	"github.com/benbjohnson/clock"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/entity"
	"github.com/judoassistant/judoassistant-meta-service/errors"
	"github.com/judoassistant/judoassistant-meta-service/mappers"
	"github.com/judoassistant/judoassistant-meta-service/repository"
)

type TournamentService interface {
	Create(tournament *dto.TournamentCreationRequestDTO, user *dto.UserResponseDTO) (*dto.TournamentResponseDTO, error)
	Delete(id int64, user *dto.UserResponseDTO) error
	GetByID(id int64) (*dto.TournamentResponseDTO, error)
	List(after int64, count int) ([]dto.TournamentResponseDTO, error)
	ListByOwner(ownerID int64) ([]dto.TournamentResponseDTO, error)
	ListPast(count int) ([]dto.TournamentResponseDTO, error)
	ListUpcoming(count int) ([]dto.TournamentResponseDTO, error)
	Update(id int64, request *dto.TournamentUpdateRequestDTO, user *dto.UserResponseDTO) (*dto.TournamentResponseDTO, error)
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

func (s *tournamentService) ListPast(count int) ([]dto.TournamentResponseDTO, error) {
	today := s.clock.Now()
	tournaments, err := s.tournamentRepository.ListByDateLessThan(today, 10)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list tournaments")
	}

	return mappers.TournamentToResponseDTOs(tournaments), nil
}

func (s *tournamentService) ListUpcoming(count int) ([]dto.TournamentResponseDTO, error) {
	today := s.clock.Now()
	tournaments, err := s.tournamentRepository.ListByDateGreaterThanEqual(today, 10)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list tournaments")
	}

	return mappers.TournamentToResponseDTOs(tournaments), nil
}

func (s *tournamentService) List(after int64, count int) ([]dto.TournamentResponseDTO, error) {
	tournaments, err := s.tournamentRepository.ListByIDGreaterThan(after, count)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list tournaments")
	}

	return mappers.TournamentToResponseDTOs(tournaments), nil
}

func (s *tournamentService) GetByID(id int64) (*dto.TournamentResponseDTO, error) {
	tournament, err := s.tournamentRepository.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tournament")
	}

	response := mappers.TournamentToResponseDTO(tournament)
	return &response, nil
}

func (s *tournamentService) ListByOwner(ownerID int64) ([]dto.TournamentResponseDTO, error) {
	tournaments, err := s.tournamentRepository.ListByOwner(ownerID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list tournament")
	}

	return mappers.TournamentToResponseDTOs(tournaments), nil
}

func (s *tournamentService) Update(id int64, request *dto.TournamentUpdateRequestDTO, user *dto.UserResponseDTO) (*dto.TournamentResponseDTO, error) {
	tournament, err := s.tournamentRepository.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get tournament")
	}
	if tournament.Owner != user.ID {
		return nil, errors.New("tournament is owned by someone else", errors.CodeForbidden)
	}

	mappers.TournamentFromUpdateRequestDTO(request, tournament)

	if err := s.tournamentRepository.Update(tournament); err != nil {
		return nil, errors.Wrap(err, "unable to update tournament")
	}

	response := mappers.TournamentToResponseDTO(tournament)
	return &response, nil
}

func (s *tournamentService) Create(request *dto.TournamentCreationRequestDTO, user *dto.UserResponseDTO) (*dto.TournamentResponseDTO, error) {
	tournament := &entity.TournamentEntity{
		Name:     request.Name,
		Location: request.Location,
		Date:     request.Date,
		Owner:    user.ID,
	}

	if err := s.tournamentRepository.Create(tournament); err != nil {
		return nil, errors.Wrap(err, "unable to create tournament")
	}

	response := mappers.TournamentToResponseDTO(tournament)
	return &response, nil
}

func (s *tournamentService) Delete(id int64, user *dto.UserResponseDTO) error {
	tournament, err := s.tournamentRepository.GetByID(id)
	if err != nil {
		return errors.Wrap(err, "unable to get tournament")
	}
	if tournament.Owner != user.ID {
		return errors.New("tournament is owned by someone else", errors.CodeForbidden)
	}

	if err := s.tournamentRepository.DeleteByID(id); err != nil {
		return errors.Wrap(err, "unable to delete tournament")
	}

	return nil
}
