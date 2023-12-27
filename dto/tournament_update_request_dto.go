package dto

import (
	"time"
)

type TournamentUpdateRequestDTO struct {
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
	URLSlug  string    `json:"urlSlug"`
}
