package postgres

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"

	"github.com/tagirmukail/forum/internal/repository"
	"github.com/tagirmukail/forum/internal/repository/postgres/migrations"
)

var _ repository.Repository = &Postgres{}

const defaultLimitPerPage = 20

type Postgres struct {
	db *sqlx.DB
	*UserPostgres
	*TopicPostgres
	*CommentPostgres
}

func NewPostgres(db *sqlx.DB) *Postgres {
	pg := &Postgres{
		db: db,
		UserPostgres: &UserPostgres{
			db: db,
		},
		TopicPostgres: &TopicPostgres{
			db: db,
		},
		CommentPostgres: &CommentPostgres{
			db: db,
		},
	}

	return pg
}

func (p *Postgres) MigrateUp() error {
	runner, err := migrations.NewRunner(p.db.DB)
	if err != nil {
		return err
	}

	err = runner.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	err = p.prepareStmt()

	return err
}

func (p *Postgres) prepareStmt() (err error) {
	err = p.UserPostgres.prepareStmt()
	if err != nil {
		return err
	}

	err = p.TopicPostgres.prepareStmt()
	if err != nil {
		return err
	}

	err = p.CommentPostgres.prepareStmt()
	if err != nil {
		return err
	}

	return nil
}
