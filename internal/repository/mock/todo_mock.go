// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/todo.go

// Package repository_mock is a generated GoMock package.
package repository_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	todo "gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/todo"
)

// MockTodoRepository is a mock of TodoRepository interface.
type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

// MockTodoRepositoryMockRecorder is the mock recorder for MockTodoRepository.
type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

// NewMockTodoRepository creates a new mock instance.
func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

// AutoMigrate mocks base method.
func (m *MockTodoRepository) AutoMigrate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoMigrate")
	ret0, _ := ret[0].(error)
	return ret0
}

// AutoMigrate indicates an expected call of AutoMigrate.
func (mr *MockTodoRepositoryMockRecorder) AutoMigrate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoMigrate", reflect.TypeOf((*MockTodoRepository)(nil).AutoMigrate))
}

// Close mocks base method.
func (m *MockTodoRepository) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockTodoRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTodoRepository)(nil).Close))
}

// Connect mocks base method.
func (m *MockTodoRepository) Connect() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Connect")
}

// Connect indicates an expected call of Connect.
func (mr *MockTodoRepositoryMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockTodoRepository)(nil).Connect))
}

// Create mocks base method.
func (m *MockTodoRepository) Create(toDo *todo.Todo) *todo.Todo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", toDo)
	ret0, _ := ret[0].(*todo.Todo)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTodoRepositoryMockRecorder) Create(toDo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodoRepository)(nil).Create), toDo)
}

// Delete mocks base method.
func (m *MockTodoRepository) Delete(id uint) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(int64)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTodoRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoRepository)(nil).Delete), id)
}

// Find mocks base method.
func (m *MockTodoRepository) Find(id uint) *todo.Todo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*todo.Todo)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockTodoRepositoryMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockTodoRepository)(nil).Find), id)
}

// FindAll mocks base method.
func (m *MockTodoRepository) FindAll() []*todo.Todo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*todo.Todo)
	return ret0
}

// FindAll indicates an expected call of FindAll.
func (mr *MockTodoRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockTodoRepository)(nil).FindAll))
}

// Update mocks base method.
func (m *MockTodoRepository) Update(toDo *todo.Todo) *todo.Todo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", toDo)
	ret0, _ := ret[0].(*todo.Todo)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTodoRepositoryMockRecorder) Update(toDo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoRepository)(nil).Update), toDo)
}
