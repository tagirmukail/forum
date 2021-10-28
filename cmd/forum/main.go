package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/tagirmukail/forum/internal/config"
	"github.com/tagirmukail/forum/internal/forum"
	"github.com/tagirmukail/forum/internal/repository/postgres"
)

var Version string

func main() {
	cfg := config.NewConfig()

	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.SetLevel(level)

	pg, err := setupDB(cfg.Postgres.ConnectionURI())
	if err != nil {
		logrus.Fatal(err)
	}

	service := forum.NewService(&forum.Dependencies{
		Repo: pg,
		Conf: cfg,
	})

	logrus.WithField("version", Version).WithField("addr", cfg.API.Addr).Info("started listen")

	err = service.Serve()
	if err != nil {
		logrus.WithError(err).Fatal("forum serve failed")
	}
}

func setupDB(connectionURL string) (*postgres.Postgres, error) {
	dbx, err := sqlx.Connect("postgres", connectionURL)
	if err != nil {
		return nil, fmt.Errorf("db connect: %v", err)
	}

	d := postgres.NewPostgres(dbx)

	err = d.MigrateUp()
	if err != nil {
		return nil, err
	}

	return d, nil
}
