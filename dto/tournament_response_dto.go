package dto

import (
	"time"
)

type TournamentResponseDTO struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
	Owner    int64     `json:"owner"`
	URLSlug  string    `json:"urlSlug"`
}
