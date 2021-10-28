package logic

import (
	"context"

	"github.com/tagirmukail/forum/internal/repository/model"

	"github.com/tagirmukail/forum/internal/dto"
)

func (l *Logic) CreateUser(ctx context.Context, user dto.UserRequest) (result dto.User, err error) {
	var userM model.User

	userM, err = l.d.Repo.NewUser(ctx, model.User{
		Username: user.Username,
	})
	if err != nil {
		return result, err
	}

	result = userM.ToDTO()

	return result, nil
}
