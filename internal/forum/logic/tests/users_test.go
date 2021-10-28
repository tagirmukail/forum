package tests

import (
	"errors"

	"github.com/tagirmukail/forum/internal/dto"
	"github.com/tagirmukail/forum/internal/repository/model"
)

func (s *logicSuite) TestCreateUserOK() {
	s.repoMock.EXPECT().NewUser(s.testCtx, model.User{Username: "test"}).
		Return(model.User{
			ID:       "17ca79c8-65ff-41b0-9765-7351e69666e0",
			Username: "test",
		}, nil)

	result, err := s.l.CreateUser(s.testCtx, dto.UserRequest{
		Username: "test",
	})
	s.Require().NoError(err)

	s.Require().NotEmpty(result.ID)
}

func (s *logicSuite) TestCreateUserErr() {
	s.repoMock.EXPECT().NewUser(s.testCtx, model.User{Username: "test"}).
		Return(model.User{}, errors.New("new user creation failed"))

	result, err := s.l.CreateUser(s.testCtx, dto.UserRequest{
		Username: "test",
	})
	s.Require().EqualError(err, "new user creation failed")

	s.Require().Empty(result)
}
