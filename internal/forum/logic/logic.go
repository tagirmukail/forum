package logic

import (
	"github.com/tagirmukail/forum/internal/config"
	"github.com/tagirmukail/forum/internal/repository"
)

type Dependencies struct {
	Repo repository.Repository
	Conf *config.Config
}

type Logic struct {
	d *Dependencies
}

func NewLogic(d *Dependencies) *Logic {
	return &Logic{
		d: d,
	}
}
