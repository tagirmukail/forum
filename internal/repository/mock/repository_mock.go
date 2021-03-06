// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tagirmukail/forum/internal/repository (interfaces: Repository)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/tagirmukail/forum/internal/repository/model"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetTopic mocks base method.
func (m *MockRepository) GetTopic(arg0 context.Context, arg1 string) (model.TopicDetailed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopic", arg0, arg1)
	ret0, _ := ret[0].(model.TopicDetailed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopic indicates an expected call of GetTopic.
func (mr *MockRepositoryMockRecorder) GetTopic(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopic", reflect.TypeOf((*MockRepository)(nil).GetTopic), arg0, arg1)
}

// ListComments mocks base method.
func (m *MockRepository) ListComments(arg0 context.Context, arg1 string, arg2, arg3 int) ([]model.Comment, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComments", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]model.Comment)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListComments indicates an expected call of ListComments.
func (mr *MockRepositoryMockRecorder) ListComments(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComments", reflect.TypeOf((*MockRepository)(nil).ListComments), arg0, arg1, arg2, arg3)
}

// ListTopics mocks base method.
func (m *MockRepository) ListTopics(arg0 context.Context, arg1, arg2 int) ([]model.Topic, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTopics", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.Topic)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTopics indicates an expected call of ListTopics.
func (mr *MockRepositoryMockRecorder) ListTopics(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTopics", reflect.TypeOf((*MockRepository)(nil).ListTopics), arg0, arg1, arg2)
}

// NewComment mocks base method.
func (m *MockRepository) NewComment(arg0 context.Context, arg1 model.Comment) (model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewComment", arg0, arg1)
	ret0, _ := ret[0].(model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewComment indicates an expected call of NewComment.
func (mr *MockRepositoryMockRecorder) NewComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewComment", reflect.TypeOf((*MockRepository)(nil).NewComment), arg0, arg1)
}

// NewTopic mocks base method.
func (m *MockRepository) NewTopic(arg0 context.Context, arg1 model.Topic) (model.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTopic", arg0, arg1)
	ret0, _ := ret[0].(model.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewTopic indicates an expected call of NewTopic.
func (mr *MockRepositoryMockRecorder) NewTopic(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTopic", reflect.TypeOf((*MockRepository)(nil).NewTopic), arg0, arg1)
}

// NewUser mocks base method.
func (m *MockRepository) NewUser(arg0 context.Context, arg1 model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUser", arg0, arg1)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewUser indicates an expected call of NewUser.
func (mr *MockRepositoryMockRecorder) NewUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUser", reflect.TypeOf((*MockRepository)(nil).NewUser), arg0, arg1)
}
