package services

import "github.com/judoassistant/judoassistant-meta-service/repositories"

type TournamentService struct {
	tournamentRepository *repositories.TournamentRepository
}

func NewTournamentService(tournamentRepository *repositories.TournamentRepository) *TournamentService {
	return &TournamentService{tournamentRepository}
}
