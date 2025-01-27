// Code generated by MockGen. DO NOT EDIT.
// Source: subject_system_group.go

// Package mock is a generated GoMock package.
package mock

import (
	dao "github.com/TencentBlueKing/bk-iam/pkg/database/dao"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockSubjectSystemGroupManager is a mock of SubjectSystemGroupManager interface.
type MockSubjectSystemGroupManager struct {
	ctrl     *gomock.Controller
	recorder *MockSubjectSystemGroupManagerMockRecorder
}

// MockSubjectSystemGroupManagerMockRecorder is the mock recorder for MockSubjectSystemGroupManager.
type MockSubjectSystemGroupManagerMockRecorder struct {
	mock *MockSubjectSystemGroupManager
}

// NewMockSubjectSystemGroupManager creates a new mock instance.
func NewMockSubjectSystemGroupManager(ctrl *gomock.Controller) *MockSubjectSystemGroupManager {
	mock := &MockSubjectSystemGroupManager{ctrl: ctrl}
	mock.recorder = &MockSubjectSystemGroupManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubjectSystemGroupManager) EXPECT() *MockSubjectSystemGroupManagerMockRecorder {
	return m.recorder
}

// CreateWithTx mocks base method.
func (m *MockSubjectSystemGroupManager) CreateWithTx(tx *sqlx.Tx, subjectSystemGroup dao.SubjectSystemGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWithTx", tx, subjectSystemGroup)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWithTx indicates an expected call of CreateWithTx.
func (mr *MockSubjectSystemGroupManagerMockRecorder) CreateWithTx(tx, subjectSystemGroup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithTx", reflect.TypeOf((*MockSubjectSystemGroupManager)(nil).CreateWithTx), tx, subjectSystemGroup)
}

// DeleteBySubjectPKsWithTx mocks base method.
func (m *MockSubjectSystemGroupManager) DeleteBySubjectPKsWithTx(tx *sqlx.Tx, subjectPKs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBySubjectPKsWithTx", tx, subjectPKs)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBySubjectPKsWithTx indicates an expected call of DeleteBySubjectPKsWithTx.
func (mr *MockSubjectSystemGroupManagerMockRecorder) DeleteBySubjectPKsWithTx(tx, subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBySubjectPKsWithTx", reflect.TypeOf((*MockSubjectSystemGroupManager)(nil).DeleteBySubjectPKsWithTx), tx, subjectPKs)
}

// GetBySystemSubject mocks base method.
func (m *MockSubjectSystemGroupManager) GetBySystemSubject(systemID string, subjectPK int64) (dao.SubjectSystemGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySystemSubject", systemID, subjectPK)
	ret0, _ := ret[0].(dao.SubjectSystemGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySystemSubject indicates an expected call of GetBySystemSubject.
func (mr *MockSubjectSystemGroupManagerMockRecorder) GetBySystemSubject(systemID, subjectPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySystemSubject", reflect.TypeOf((*MockSubjectSystemGroupManager)(nil).GetBySystemSubject), systemID, subjectPK)
}

// ListSubjectGroups mocks base method.
func (m *MockSubjectSystemGroupManager) ListSubjectGroups(systemID string, subjectPKs []int64) ([]dao.SubjectGroups, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubjectGroups", systemID, subjectPKs)
	ret0, _ := ret[0].([]dao.SubjectGroups)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubjectGroups indicates an expected call of ListSubjectGroups.
func (mr *MockSubjectSystemGroupManagerMockRecorder) ListSubjectGroups(systemID, subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjectGroups", reflect.TypeOf((*MockSubjectSystemGroupManager)(nil).ListSubjectGroups), systemID, subjectPKs)
}

// UpdateWithTx mocks base method.
func (m *MockSubjectSystemGroupManager) UpdateWithTx(tx *sqlx.Tx, subjectSystemGroup dao.SubjectSystemGroup) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWithTx", tx, subjectSystemGroup)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWithTx indicates an expected call of UpdateWithTx.
func (mr *MockSubjectSystemGroupManagerMockRecorder) UpdateWithTx(tx, subjectSystemGroup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWithTx", reflect.TypeOf((*MockSubjectSystemGroupManager)(nil).UpdateWithTx), tx, subjectSystemGroup)
}
