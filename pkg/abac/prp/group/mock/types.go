// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package mock is a generated GoMock package.
package mock

import (
	types "github.com/TencentBlueKing/bk-iam/pkg/service/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGroupAuthTypeRetriever is a mock of GroupAuthTypeRetriever interface.
type MockGroupAuthTypeRetriever struct {
	ctrl     *gomock.Controller
	recorder *MockGroupAuthTypeRetrieverMockRecorder
}

// MockGroupAuthTypeRetrieverMockRecorder is the mock recorder for MockGroupAuthTypeRetriever.
type MockGroupAuthTypeRetrieverMockRecorder struct {
	mock *MockGroupAuthTypeRetriever
}

// NewMockGroupAuthTypeRetriever creates a new mock instance.
func NewMockGroupAuthTypeRetriever(ctrl *gomock.Controller) *MockGroupAuthTypeRetriever {
	mock := &MockGroupAuthTypeRetriever{ctrl: ctrl}
	mock.recorder = &MockGroupAuthTypeRetrieverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupAuthTypeRetriever) EXPECT() *MockGroupAuthTypeRetrieverMockRecorder {
	return m.recorder
}

// Retrieve mocks base method.
func (m *MockGroupAuthTypeRetriever) Retrieve(groupPKs []int64) ([]types.GroupAuthType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retrieve", groupPKs)
	ret0, _ := ret[0].([]types.GroupAuthType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Retrieve indicates an expected call of Retrieve.
func (mr *MockGroupAuthTypeRetrieverMockRecorder) Retrieve(groupPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retrieve", reflect.TypeOf((*MockGroupAuthTypeRetriever)(nil).Retrieve), groupPKs)
}
