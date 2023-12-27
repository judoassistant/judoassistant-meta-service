package entity

import "time"

type TournamentEntity struct {
	ID        int64     `db:"id"`
	ShortName string    `db:"short_name"`
	Name      string    `db:"name"`
	Location  string    `db:"location"`
	Date      time.Time `db:"date"`
	IsDeleted bool      `db:"is_deleted"`
	Owner     int64     `db:"owner"`
}
