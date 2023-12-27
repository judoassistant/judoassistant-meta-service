package dto

import (
	"time"
)

type TournamentResponseDTO struct {
	ShortName string    `json:"shortName"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Date      time.Time `json:"date"`
	Owner     int64     `json:"owner"`
}
