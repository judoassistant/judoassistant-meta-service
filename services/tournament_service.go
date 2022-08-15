package services

import (
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/repositories"
)

type TournamentService struct {
	tournamentRepository *repositories.TournamentRepository
}

func NewTournamentService(tournamentRepository *repositories.TournamentRepository) *TournamentService {
	return &TournamentService{tournamentRepository}
}

func (ts *TournamentService) GetPast(count int) ([]dto.TournamentIndexResponseDTO, error) {
	return nil, nil
}

func (ts *TournamentService) GetUpcoming(count int) ([]dto.TournamentIndexResponseDTO, error) {
	return nil, nil
}

func (ts *TournamentService) Get(after int64, count int) ([]dto.TournamentIndexResponseDTO, error) {
	return nil, nil
}

func (ts *TournamentService) Create(tournament *dto.TournamentCreationRequestDTO) (*dto.TournamentCreationResponseDTO, error) {
	return nil, nil
}
