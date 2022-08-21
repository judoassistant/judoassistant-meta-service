package dto

import (
	"time"

	"github.com/judoassistant/judoassistant-meta-service/entities"
)

type TournamentResponseDTO struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
}

func MapTournamentResponseDTO(tournament *entities.TournamentEntity) TournamentResponseDTO {
	return TournamentResponseDTO{
		Name:     tournament.Name,
		Location: tournament.Location,
		Date:     tournament.Date,
	}
}

func MapTournamentResponseDTOs(tournaments []entities.TournamentEntity) []TournamentResponseDTO {
	result := make([]TournamentResponseDTO, len(tournaments))
	for key, value := range tournaments {
		result[key] = MapTournamentResponseDTO(&value)
	}

	return result
}
