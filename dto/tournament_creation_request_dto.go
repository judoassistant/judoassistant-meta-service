package dto

import "time"

type TournamentCreationRequestDTO struct {
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
}
