package tests

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"github.com/tagirmukail/forum/internal/config"
	"github.com/tagirmukail/forum/internal/forum/logic"

	"github.com/tagirmukail/forum/internal/repository"
	"github.com/tagirmukail/forum/internal/repository/mock"
)

type logicSuite struct {
	suite.Suite

	conf *config.Config

	repo repository.Repository

	ctrl     *gomock.Controller
	repoMock *mock.MockRepository

	testCtx    context.Context
	testCancel context.CancelFunc

	l *logic.Logic
}

func (s *logicSuite) SetupSuite() {

}

func (s *logicSuite) TearDownSuite() {

}

func (s *logicSuite) SetupTest() {
	s.testCtx, s.testCancel = context.WithTimeout(context.Background(), 5*time.Second)

	s.conf = &config.Config{
		LogLevel: "debug",
	}

	s.ctrl = gomock.NewController(s.T())
	s.repoMock = mock.NewMockRepository(s.ctrl)

	s.repo = s.repoMock

	s.l = logic.NewLogic(&logic.Dependencies{
		Repo: s.repo,
		Conf: s.conf,
	})
}

func (s *logicSuite) TearDownTest() {
	s.testCancel()
	s.ctrl.Finish()
}

func TestLogic(t *testing.T) {
	suite.Run(t, &logicSuite{})
}
