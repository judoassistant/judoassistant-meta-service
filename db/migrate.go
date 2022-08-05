package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(db *sqlx.DB) error {
	config := sqlite3.Config{}
	driver, err := sqlite3.WithInstance(db.DB, &config)

	if err != nil {
		return err
	}

	migrate, err := migrate.NewWithDatabaseInstance("file://db/migrations", "sqlite3", driver)
	log.Println("Test")

	return migrate.Up()
}
