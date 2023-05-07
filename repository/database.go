package repository

import (
	"github.com/golang-migrate/migrate/v4"
	migratedriver "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/judoassistant/judoassistant-meta-service/config"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

const _migrationsPath string = "file://repository/migrations"

func NewDatabase(config *config.Config) (*sqlx.DB, error) {
	// Connect to database
	db, err := sqlx.Connect("sqlite3", config.DatabasePath)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize database")
	}

	// Perform migrations
	driver, err := migratedriver.WithInstance(db.DB, &migratedriver.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "unable to create migration driver")
	}

	migrations, err := migrate.NewWithDatabaseInstance(_migrationsPath, "sqlite3", driver)
	if err != nil {
		return nil, errors.Wrap(err, "unable to load database migrations")
	}

	if err := migrations.Up(); !errors.Is(err, migrate.ErrNoChange) {
		return nil, errors.Wrap(err, "unable to execute database migrations")
	}

	return db, nil
}
