package dto

import "time"

type TournamentCreationRequestDTO struct {
	ShortName string    `json:"shortName"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Date      time.Time `json:"date"`
}

type TournamentResponseDTO struct {
	ShortName string    `json:"shortName"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Date      time.Time `json:"date"`
	Owner     int64     `json:"owner"`
}

type TournamentQueryDTO struct {
	ShortName string `uri:"shortName"`
}
type TournamentIndexQueryDTO struct {
	After string `form:"after"`
}
type TournamentUpdateRequestDTO struct {
	ShortName string    `json:"shortName"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Date      time.Time `json:"date"`
}
