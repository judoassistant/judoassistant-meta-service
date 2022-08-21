package entities

import "time"

type TournamentEntity struct {
	ID       int64     `db:"id"`
	Name     string    `db:"name"`
	Location string    `db:"location"`
	Date     time.Time `db:"date"`
}
