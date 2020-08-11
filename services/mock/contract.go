// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/contract.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	services "github.com/labbsr0x/githunter-api/services"
	reflect "reflect"
)

// MockContract is a mock of Contract interface
type MockContract struct {
	ctrl     *gomock.Controller
	recorder *MockContractMockRecorder
}

// MockContractMockRecorder is the mock recorder for MockContract
type MockContractMockRecorder struct {
	mock *MockContract
}

// NewMockContract creates a new mock instance
func NewMockContract(ctrl *gomock.Controller) *MockContract {
	mock := &MockContract{ctrl: ctrl}
	mock.recorder = &MockContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContract) EXPECT() *MockContractMockRecorder {
	return m.recorder
}

// GetLastRepos mocks base method
func (m *MockContract) GetLastRepos(arg0 int, arg1, arg2 string) (*services.ReposResponseContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastRepos", arg0, arg1, arg2)
	ret0, _ := ret[0].(*services.ReposResponseContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastRepos indicates an expected call of GetLastRepos
func (mr *MockContractMockRecorder) GetLastRepos(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastRepos", reflect.TypeOf((*MockContract)(nil).GetLastRepos), arg0, arg1, arg2)
}

// GetIssues mocks base method
func (m *MockContract) GetIssues(arg0 int, arg1, arg2, arg3, arg4 string) (*services.IssuesResponseContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssues", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*services.IssuesResponseContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssues indicates an expected call of GetIssues
func (mr *MockContractMockRecorder) GetIssues(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssues", reflect.TypeOf((*MockContract)(nil).GetIssues), arg0, arg1, arg2, arg3, arg4)
}

// GetPulls mocks base method
func (m *MockContract) GetPulls(arg0 int, arg1, arg2, arg3, arg4 string) (*services.PullsResponseContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPulls", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*services.PullsResponseContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPulls indicates an expected call of GetPulls
func (mr *MockContractMockRecorder) GetPulls(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPulls", reflect.TypeOf((*MockContract)(nil).GetPulls), arg0, arg1, arg2, arg3, arg4)
}
