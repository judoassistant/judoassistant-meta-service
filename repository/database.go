package repository

import (
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	migratedriver "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/judoassistant/judoassistant-meta-service/config"
	judoerrors "github.com/judoassistant/judoassistant-meta-service/errors"
	_ "github.com/mattn/go-sqlite3"
)

const _migrationsPath string = "file://repository/migrations"

func NewDatabase(config *config.Config) (*sqlx.DB, error) {
	// Connect to database
	db, err := sqlx.Connect("sqlite3", config.DatabasePath)
	if err != nil {
		return nil, judoerrors.WrapWithCode(err, "unable to initialize database", errorCodeFromDatabaseError(err))
	}

	// Perform migrations
	driver, err := migratedriver.WithInstance(db.DB, &migratedriver.Config{})
	if err != nil {
		return nil, judoerrors.WrapWithCode(err, "unable to create migration driver", errorCodeFromDatabaseError(err))
	}

	migrations, err := migrate.NewWithDatabaseInstance(_migrationsPath, "sqlite3", driver)
	if err != nil {
		return nil, judoerrors.WrapWithCode(err, "unable to load database migrations", errorCodeFromDatabaseError(err))
	}

	if err := migrations.Up(); !errors.Is(err, migrate.ErrNoChange) {
		return nil, judoerrors.WrapWithCode(err, "unable to execute database migrations", errorCodeFromDatabaseError(err))
	}

	return db, nil
}

func errorCodeFromDatabaseError(err error) int {
	if errors.Is(err, sql.ErrNoRows) {
		return judoerrors.CodeNotFound
	}

	return judoerrors.CodeUnavailable
}
