// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/contract.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	services "github.com/labbsr0x/githunter-api/services"
)

// MockContract is a mock of Contract interface.
type MockContract struct {
	ctrl     *gomock.Controller
	recorder *MockContractMockRecorder
}

// MockContractMockRecorder is the mock recorder for MockContract.
type MockContractMockRecorder struct {
	mock *MockContract
}

// NewMockContract creates a new mock instance.
func NewMockContract(ctrl *gomock.Controller) *MockContract {
	mock := &MockContract{ctrl: ctrl}
	mock.recorder = &MockContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContract) EXPECT() *MockContractMockRecorder {
	return m.recorder
}

// GetComments mocks base method.
func (m *MockContract) GetComments(arg0 []string, arg1, arg2 string) (*services.CommentsResponseContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComments", arg0, arg1, arg2)
	ret0, _ := ret[0].(*services.CommentsResponseContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComments indicates an expected call of GetComments.
func (mr *MockContractMockRecorder) GetComments(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComments", reflect.TypeOf((*MockContract)(nil).GetComments), arg0, arg1, arg2)
}

// GetCommitsRepo mocks base method.
func (m *MockContract) GetCommitsRepo(arg0, arg1, arg2, arg3 string) (*services.CommitsResponseContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommitsRepo", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*services.CommitsResponseContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommitsRepo indicates an expected call of GetCommitsRepo.
func (mr *MockContractMockRecorder) GetCommitsRepo(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommitsRepo", reflect.TypeOf((*MockContract)(nil).GetCommitsRepo), arg0, arg1, arg2, arg3)
}

// GetInfoCodePage mocks base method.
func (m *MockContract) GetInfoCodePage(arg0, arg1, arg2, arg3 string) (*services.CodeResponseContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfoCodePage", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*services.CodeResponseContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfoCodePage indicates an expected call of GetInfoCodePage.
func (mr *MockContractMockRecorder) GetInfoCodePage(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfoCodePage", reflect.TypeOf((*MockContract)(nil).GetInfoCodePage), arg0, arg1, arg2, arg3)
}

// GetIssues mocks base method.
func (m *MockContract) GetIssues(arg0, arg1, arg2, arg3 string) (*services.IssuesResponseContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssues", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*services.IssuesResponseContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssues indicates an expected call of GetIssues.
func (mr *MockContractMockRecorder) GetIssues(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssues", reflect.TypeOf((*MockContract)(nil).GetIssues), arg0, arg1, arg2, arg3)
}

// GetLastRepos mocks base method.
func (m *MockContract) GetLastRepos(arg0, arg1 string) (*services.ReposResponseContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastRepos", arg0, arg1)
	ret0, _ := ret[0].(*services.ReposResponseContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastRepos indicates an expected call of GetLastRepos.
func (mr *MockContractMockRecorder) GetLastRepos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastRepos", reflect.TypeOf((*MockContract)(nil).GetLastRepos), arg0, arg1)
}

// GetMembers mocks base method.
func (m *MockContract) GetMembers(arg0, arg1, arg2 string) (*services.OrganizationResponseContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMembers", arg0, arg1, arg2)
	ret0, _ := ret[0].(*services.OrganizationResponseContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMembers indicates an expected call of GetMembers.
func (mr *MockContractMockRecorder) GetMembers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembers", reflect.TypeOf((*MockContract)(nil).GetMembers), arg0, arg1, arg2)
}

// GetPulls mocks base method.
func (m *MockContract) GetPulls(arg0, arg1, arg2, arg3 string) (*services.PullsResponseContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPulls", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*services.PullsResponseContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPulls indicates an expected call of GetPulls.
func (mr *MockContractMockRecorder) GetPulls(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPulls", reflect.TypeOf((*MockContract)(nil).GetPulls), arg0, arg1, arg2, arg3)
}

// GetUserScore mocks base method.
func (m *MockContract) GetUserScore(arg0, arg1, arg2 string) (*services.UserScoreResponseContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserScore", arg0, arg1, arg2)
	ret0, _ := ret[0].(*services.UserScoreResponseContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserScore indicates an expected call of GetUserScore.
func (mr *MockContractMockRecorder) GetUserScore(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserScore", reflect.TypeOf((*MockContract)(nil).GetUserScore), arg0, arg1, arg2)
}

// GetUserStats mocks base method.
func (m *MockContract) GetUserStats(arg0, arg1, arg2 string) (*services.UserResponseContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserStats", arg0, arg1, arg2)
	ret0, _ := ret[0].(*services.UserResponseContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserStats indicates an expected call of GetUserStats.
func (mr *MockContractMockRecorder) GetUserStats(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserStats", reflect.TypeOf((*MockContract)(nil).GetUserStats), arg0, arg1, arg2)
}
