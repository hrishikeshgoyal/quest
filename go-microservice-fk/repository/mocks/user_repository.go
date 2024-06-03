// Code generated by MockGen. DO NOT EDIT.
// Source: user_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/hrishikeshgoyal/quest/go-microservice-fk/entity"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(user *entity.User) *entity.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(*entity.User)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), user)
}

// GetAll mocks base method.
func (m *MockUserRepository) GetAll() []entity.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entity.User)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockUserRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUserRepository)(nil).GetAll))
}

// GetUser mocks base method.
func (m *MockUserRepository) GetUser(id string) *entity.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(*entity.User)
	return ret0
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepositoryMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepository)(nil).GetUser), id)
}

// GetUserByName mocks base method.
func (m *MockUserRepository) GetUserByName(name string) *entity.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", name)
	ret0, _ := ret[0].(*entity.User)
	return ret0
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockUserRepositoryMockRecorder) GetUserByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockUserRepository)(nil).GetUserByName), name)
}