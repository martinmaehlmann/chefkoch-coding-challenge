// Code generated by MockGen. DO NOT EDIT.
// Source: internal/config/registry.go

// Package config_mock is a generated GoMock package.
package config_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockPrettyStringer is a mock of PrettyStringer interface.
type MockPrettyStringer struct {
	ctrl     *gomock.Controller
	recorder *MockPrettyStringerMockRecorder
}

// MockPrettyStringerMockRecorder is the mock recorder for MockPrettyStringer.
type MockPrettyStringerMockRecorder struct {
	mock *MockPrettyStringer
}

// NewMockPrettyStringer creates a new mock instance.
func NewMockPrettyStringer(ctrl *gomock.Controller) *MockPrettyStringer {
	mock := &MockPrettyStringer{ctrl: ctrl}
	mock.recorder = &MockPrettyStringerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrettyStringer) EXPECT() *MockPrettyStringerMockRecorder {
	return m.recorder
}

// PrettyString mocks base method.
func (m *MockPrettyStringer) PrettyString(indentation string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrettyString", indentation)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrettyString indicates an expected call of PrettyString.
func (mr *MockPrettyStringerMockRecorder) PrettyString(indentation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrettyString", reflect.TypeOf((*MockPrettyStringer)(nil).PrettyString), indentation)
}

// MockDialector is a mock of Dialector interface.
type MockDialector struct {
	ctrl     *gomock.Controller
	recorder *MockDialectorMockRecorder
}

// MockDialectorMockRecorder is the mock recorder for MockDialector.
type MockDialectorMockRecorder struct {
	mock *MockDialector
}

// NewMockDialector creates a new mock instance.
func NewMockDialector(ctrl *gomock.Controller) *MockDialector {
	mock := &MockDialector{ctrl: ctrl}
	mock.recorder = &MockDialectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDialector) EXPECT() *MockDialectorMockRecorder {
	return m.recorder
}

// DSN mocks base method.
func (m *MockDialector) DSN() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DSN")
	ret0, _ := ret[0].(string)
	return ret0
}

// DSN indicates an expected call of DSN.
func (mr *MockDialectorMockRecorder) DSN() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DSN", reflect.TypeOf((*MockDialector)(nil).DSN))
}

// Dialector mocks base method.
func (m *MockDialector) Dialector() gorm.Dialector {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dialector")
	ret0, _ := ret[0].(gorm.Dialector)
	return ret0
}

// Dialector indicates an expected call of Dialector.
func (mr *MockDialectorMockRecorder) Dialector() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dialector", reflect.TypeOf((*MockDialector)(nil).Dialector))
}
