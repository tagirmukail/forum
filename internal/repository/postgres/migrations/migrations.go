package migrations

import (
	"database/sql"
	"embed"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
)

//go:embed *.sql
var fs embed.FS

// NewRunner builds migrations executor for PostgreSQL.
func NewRunner(db *sql.DB) (*migrate.Migrate, error) {
	src, err := sourceDriver()
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	return migrate.NewWithInstance("httpfs", src, "postgres", driver)
}

func sourceDriver() (source.Driver, error) {
	return httpfs.New(http.FS(fs), ".")
}
