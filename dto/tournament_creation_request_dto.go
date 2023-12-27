package dto

import "time"

type TournamentCreationRequestDTO struct {
	ShortName string    `json:"shortName"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Date      time.Time `json:"date"`
}
