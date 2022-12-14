package dto

import (
	"time"

	"github.com/judoassistant/judoassistant-meta-service/entity"
)

type TournamentResponseDTO struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
	Owner    int64     `json:"owner"`
}

func MapTournamentResponseDTO(tournament *entity.TournamentEntity) TournamentResponseDTO {
	return TournamentResponseDTO{
		Name:     tournament.Name,
		Location: tournament.Location,
		Date:     tournament.Date,
		Owner:    tournament.Owner,
	}
}

func MapTournamentResponseDTOs(tournaments []entity.TournamentEntity) []TournamentResponseDTO {
	result := make([]TournamentResponseDTO, len(tournaments))
	for key, value := range tournaments {
		result[key] = MapTournamentResponseDTO(&value)
	}

	return result
}
