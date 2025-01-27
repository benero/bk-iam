// Code generated by MockGen. DO NOT EDIT.
// Source: department.go

// Package mock is a generated GoMock package.
package mock

import (
	types "github.com/TencentBlueKing/bk-iam/pkg/service/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockDepartmentService is a mock of DepartmentService interface.
type MockDepartmentService struct {
	ctrl     *gomock.Controller
	recorder *MockDepartmentServiceMockRecorder
}

// MockDepartmentServiceMockRecorder is the mock recorder for MockDepartmentService.
type MockDepartmentServiceMockRecorder struct {
	mock *MockDepartmentService
}

// NewMockDepartmentService creates a new mock instance.
func NewMockDepartmentService(ctrl *gomock.Controller) *MockDepartmentService {
	mock := &MockDepartmentService{ctrl: ctrl}
	mock.recorder = &MockDepartmentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDepartmentService) EXPECT() *MockDepartmentServiceMockRecorder {
	return m.recorder
}

// BulkCreate mocks base method.
func (m *MockDepartmentService) BulkCreate(subjectDepartments []types.SubjectDepartment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreate", subjectDepartments)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreate indicates an expected call of BulkCreate.
func (mr *MockDepartmentServiceMockRecorder) BulkCreate(subjectDepartments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreate", reflect.TypeOf((*MockDepartmentService)(nil).BulkCreate), subjectDepartments)
}

// BulkDelete mocks base method.
func (m *MockDepartmentService) BulkDelete(subjectPKs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDelete", subjectPKs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDelete indicates an expected call of BulkDelete.
func (mr *MockDepartmentServiceMockRecorder) BulkDelete(subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDelete", reflect.TypeOf((*MockDepartmentService)(nil).BulkDelete), subjectPKs)
}

// BulkDeleteBySubjectPKsWithTx mocks base method.
func (m *MockDepartmentService) BulkDeleteBySubjectPKsWithTx(tx *sqlx.Tx, pks []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteBySubjectPKsWithTx", tx, pks)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteBySubjectPKsWithTx indicates an expected call of BulkDeleteBySubjectPKsWithTx.
func (mr *MockDepartmentServiceMockRecorder) BulkDeleteBySubjectPKsWithTx(tx, pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteBySubjectPKsWithTx", reflect.TypeOf((*MockDepartmentService)(nil).BulkDeleteBySubjectPKsWithTx), tx, pks)
}

// BulkUpdate mocks base method.
func (m *MockDepartmentService) BulkUpdate(subjectDepartments []types.SubjectDepartment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkUpdate", subjectDepartments)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkUpdate indicates an expected call of BulkUpdate.
func (mr *MockDepartmentServiceMockRecorder) BulkUpdate(subjectDepartments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUpdate", reflect.TypeOf((*MockDepartmentService)(nil).BulkUpdate), subjectDepartments)
}

// GetCount mocks base method.
func (m *MockDepartmentService) GetCount() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockDepartmentServiceMockRecorder) GetCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockDepartmentService)(nil).GetCount))
}

// GetSubjectDepartmentPKs mocks base method.
func (m *MockDepartmentService) GetSubjectDepartmentPKs(subjectPK int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubjectDepartmentPKs", subjectPK)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubjectDepartmentPKs indicates an expected call of GetSubjectDepartmentPKs.
func (mr *MockDepartmentServiceMockRecorder) GetSubjectDepartmentPKs(subjectPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubjectDepartmentPKs", reflect.TypeOf((*MockDepartmentService)(nil).GetSubjectDepartmentPKs), subjectPK)
}

// ListPaging mocks base method.
func (m *MockDepartmentService) ListPaging(limit, offset int64) ([]types.SubjectDepartment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPaging", limit, offset)
	ret0, _ := ret[0].([]types.SubjectDepartment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPaging indicates an expected call of ListPaging.
func (mr *MockDepartmentServiceMockRecorder) ListPaging(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPaging", reflect.TypeOf((*MockDepartmentService)(nil).ListPaging), limit, offset)
}
