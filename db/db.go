package db

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	migratedriver "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Init() (*sqlx.DB, error) {
	return sqlx.Connect("sqlite3", "test.db")
}

func Migrate(db *sqlx.DB) error {
	config := migratedriver.Config{}
	driver, err := migratedriver.WithInstance(db.DB, &config)

	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "sqlite3", driver)

	if err := m.Up(); errors.Is(err, migrate.ErrNoChange) {
		return nil
	} else {
		return err
	}
}
