// Code generated by MockGen. DO NOT EDIT.
// Source: subject_department.go

// Package mock is a generated GoMock package.
package mock

import (
	dao "github.com/TencentBlueKing/bk-iam/pkg/database/dao"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockSubjectDepartmentManager is a mock of SubjectDepartmentManager interface.
type MockSubjectDepartmentManager struct {
	ctrl     *gomock.Controller
	recorder *MockSubjectDepartmentManagerMockRecorder
}

// MockSubjectDepartmentManagerMockRecorder is the mock recorder for MockSubjectDepartmentManager.
type MockSubjectDepartmentManagerMockRecorder struct {
	mock *MockSubjectDepartmentManager
}

// NewMockSubjectDepartmentManager creates a new mock instance.
func NewMockSubjectDepartmentManager(ctrl *gomock.Controller) *MockSubjectDepartmentManager {
	mock := &MockSubjectDepartmentManager{ctrl: ctrl}
	mock.recorder = &MockSubjectDepartmentManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubjectDepartmentManager) EXPECT() *MockSubjectDepartmentManagerMockRecorder {
	return m.recorder
}

// BulkCreate mocks base method.
func (m *MockSubjectDepartmentManager) BulkCreate(subjectDepartments []dao.SubjectDepartment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreate", subjectDepartments)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreate indicates an expected call of BulkCreate.
func (mr *MockSubjectDepartmentManagerMockRecorder) BulkCreate(subjectDepartments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreate", reflect.TypeOf((*MockSubjectDepartmentManager)(nil).BulkCreate), subjectDepartments)
}

// BulkDelete mocks base method.
func (m *MockSubjectDepartmentManager) BulkDelete(subjectPKs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDelete", subjectPKs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDelete indicates an expected call of BulkDelete.
func (mr *MockSubjectDepartmentManagerMockRecorder) BulkDelete(subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDelete", reflect.TypeOf((*MockSubjectDepartmentManager)(nil).BulkDelete), subjectPKs)
}

// BulkDeleteWithTx mocks base method.
func (m *MockSubjectDepartmentManager) BulkDeleteWithTx(tx *sqlx.Tx, subjectPKs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteWithTx", tx, subjectPKs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteWithTx indicates an expected call of BulkDeleteWithTx.
func (mr *MockSubjectDepartmentManagerMockRecorder) BulkDeleteWithTx(tx, subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteWithTx", reflect.TypeOf((*MockSubjectDepartmentManager)(nil).BulkDeleteWithTx), tx, subjectPKs)
}

// BulkUpdate mocks base method.
func (m *MockSubjectDepartmentManager) BulkUpdate(subjectDepartments []dao.SubjectDepartment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkUpdate", subjectDepartments)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkUpdate indicates an expected call of BulkUpdate.
func (mr *MockSubjectDepartmentManagerMockRecorder) BulkUpdate(subjectDepartments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUpdate", reflect.TypeOf((*MockSubjectDepartmentManager)(nil).BulkUpdate), subjectDepartments)
}

// Get mocks base method.
func (m *MockSubjectDepartmentManager) Get(subjectPK int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", subjectPK)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSubjectDepartmentManagerMockRecorder) Get(subjectPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSubjectDepartmentManager)(nil).Get), subjectPK)
}

// GetCount mocks base method.
func (m *MockSubjectDepartmentManager) GetCount() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockSubjectDepartmentManagerMockRecorder) GetCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockSubjectDepartmentManager)(nil).GetCount))
}

// ListPaging mocks base method.
func (m *MockSubjectDepartmentManager) ListPaging(limit, offset int64) ([]dao.SubjectDepartment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPaging", limit, offset)
	ret0, _ := ret[0].([]dao.SubjectDepartment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPaging indicates an expected call of ListPaging.
func (mr *MockSubjectDepartmentManagerMockRecorder) ListPaging(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPaging", reflect.TypeOf((*MockSubjectDepartmentManager)(nil).ListPaging), limit, offset)
}
