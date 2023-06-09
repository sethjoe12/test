// Code generated by MockGen. DO NOT EDIT.
// Source: test/controllers (interfaces: Iuser)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockIuser is a mock of Iuser interface.
type MockIuser struct {
	ctrl     *gomock.Controller
	recorder *MockIuserMockRecorder
}

// MockIuserMockRecorder is the mock recorder for MockIuser.
type MockIuserMockRecorder struct {
	mock *MockIuser
}

// NewMockIuser creates a new mock instance.
func NewMockIuser(ctrl *gomock.Controller) *MockIuser {
	mock := &MockIuser{ctrl: ctrl}
	mock.recorder = &MockIuserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIuser) EXPECT() *MockIuserMockRecorder {
	return m.recorder
}

// AllUsers mocks base method.
func (m *MockIuser) AllUsers(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AllUsers", arg0)
}

// AllUsers indicates an expected call of AllUsers.
func (mr *MockIuserMockRecorder) AllUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllUsers", reflect.TypeOf((*MockIuser)(nil).AllUsers), arg0)
}

// DeleteUserById mocks base method.
func (m *MockIuser) DeleteUserById(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUserById", arg0)
}

// DeleteUserById indicates an expected call of DeleteUserById.
func (mr *MockIuserMockRecorder) DeleteUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserById", reflect.TypeOf((*MockIuser)(nil).DeleteUserById), arg0)
}

// GetUserById mocks base method.
func (m *MockIuser) GetUserById(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUserById", arg0)
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockIuserMockRecorder) GetUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockIuser)(nil).GetUserById), arg0)
}

// UpdateUserById mocks base method.
func (m *MockIuser) UpdateUserById(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateUserById", arg0)
}

// UpdateUserById indicates an expected call of UpdateUserById.
func (mr *MockIuserMockRecorder) UpdateUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserById", reflect.TypeOf((*MockIuser)(nil).UpdateUserById), arg0)
}
