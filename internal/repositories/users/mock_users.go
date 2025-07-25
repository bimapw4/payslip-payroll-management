// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repositories/users/contract.go

// Package users is a generated GoMock package.
package users

import (
	context "context"
	presentations "payslips/internal/presentations"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUsers is a mock of Users interface.
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers.
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance.
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUsers) Create(ctx context.Context, input presentations.Users) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUsersMockRecorder) Create(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsers)(nil).Create), ctx, input)
}

// DeleteUser mocks base method.
func (m *MockUsers) DeleteUser(ctx context.Context, userID, updatedBy string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, userID, updatedBy)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUsersMockRecorder) DeleteUser(ctx, userID, updatedBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUsers)(nil).DeleteUser), ctx, userID, updatedBy)
}

// Detail mocks base method.
func (m *MockUsers) Detail(ctx context.Context, id string) (*presentations.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Detail", ctx, id)
	ret0, _ := ret[0].(*presentations.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Detail indicates an expected call of Detail.
func (mr *MockUsersMockRecorder) Detail(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Detail", reflect.TypeOf((*MockUsers)(nil).Detail), ctx, id)
}

// GetAllUsers mocks base method.
func (m *MockUsers) GetAllUsers(ctx context.Context) ([]presentations.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", ctx)
	ret0, _ := ret[0].([]presentations.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUsersMockRecorder) GetAllUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUsers)(nil).GetAllUsers), ctx)
}

// GetUserByUsername mocks base method.
func (m *MockUsers) GetUserByUsername(ctx context.Context, username string) (*presentations.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", ctx, username)
	ret0, _ := ret[0].(*presentations.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockUsersMockRecorder) GetUserByUsername(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockUsers)(nil).GetUserByUsername), ctx, username)
}

// Update mocks base method.
func (m *MockUsers) Update(ctx context.Context, payload presentations.Users) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUsersMockRecorder) Update(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUsers)(nil).Update), ctx, payload)
}

// UpdatePassword mocks base method.
func (m *MockUsers) UpdatePassword(ctx context.Context, userID, password, updatedBy string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", ctx, userID, password, updatedBy)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUsersMockRecorder) UpdatePassword(ctx, userID, password, updatedBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUsers)(nil).UpdatePassword), ctx, userID, password, updatedBy)
}
