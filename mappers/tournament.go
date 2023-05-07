package mappers

import (
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/entity"
)

func TournamentToResponseDTO(tournament *entity.TournamentEntity) dto.TournamentResponseDTO {
	return dto.TournamentResponseDTO{
		Name:     tournament.Name,
		Location: tournament.Location,
		Date:     tournament.Date,
		Owner:    tournament.Owner,
	}
}

func TournamentToResponseDTOs(tournaments []entity.TournamentEntity) []dto.TournamentResponseDTO {
	result := make([]dto.TournamentResponseDTO, len(tournaments))
	for key, value := range tournaments {
		result[key] = TournamentToResponseDTO(&value)
	}

	return result
}

func TournamentFromUpdateRequestDTO(dto *dto.TournamentUpdateRequestDTO, entity *entity.TournamentEntity) {
	entity.Name = dto.Name
	entity.Location = dto.Location
	entity.Date = dto.Date
}
