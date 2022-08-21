package dto

import (
	"time"

	"github.com/judoassistant/judoassistant-meta-service/entities"
)

type TournamentUpdateRequestDTO struct {
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
}

func MapToTournamentEntity(dto *TournamentUpdateRequestDTO, entity *entities.TournamentEntity) {
	entity.Name = dto.Name
	entity.Location = dto.Location
	entity.Date = dto.Date
}
