package db

import "github.com/jmoiron/sqlx"

func Init() (*sqlx.DB, error) {
	return sqlx.Connect("sqllite3", ":memory")
}
