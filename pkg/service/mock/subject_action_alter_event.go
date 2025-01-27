// Code generated by MockGen. DO NOT EDIT.
// Source: subject_action_alter_event.go

// Package mock is a generated GoMock package.
package mock

import (
	types "github.com/TencentBlueKing/bk-iam/pkg/service/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockSubjectActionAlterEventService is a mock of SubjectActionAlterEventService interface.
type MockSubjectActionAlterEventService struct {
	ctrl     *gomock.Controller
	recorder *MockSubjectActionAlterEventServiceMockRecorder
}

// MockSubjectActionAlterEventServiceMockRecorder is the mock recorder for MockSubjectActionAlterEventService.
type MockSubjectActionAlterEventServiceMockRecorder struct {
	mock *MockSubjectActionAlterEventService
}

// NewMockSubjectActionAlterEventService creates a new mock instance.
func NewMockSubjectActionAlterEventService(ctrl *gomock.Controller) *MockSubjectActionAlterEventService {
	mock := &MockSubjectActionAlterEventService{ctrl: ctrl}
	mock.recorder = &MockSubjectActionAlterEventServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubjectActionAlterEventService) EXPECT() *MockSubjectActionAlterEventServiceMockRecorder {
	return m.recorder
}

// BulkCreateWithTx mocks base method.
func (m *MockSubjectActionAlterEventService) BulkCreateWithTx(tx *sqlx.Tx, events []types.SubjectActionAlterEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateWithTx", tx, events)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreateWithTx indicates an expected call of BulkCreateWithTx.
func (mr *MockSubjectActionAlterEventServiceMockRecorder) BulkCreateWithTx(tx, events interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateWithTx", reflect.TypeOf((*MockSubjectActionAlterEventService)(nil).BulkCreateWithTx), tx, events)
}

// BulkIncrCheckCount mocks base method.
func (m *MockSubjectActionAlterEventService) BulkIncrCheckCount(uuids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkIncrCheckCount", uuids)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkIncrCheckCount indicates an expected call of BulkIncrCheckCount.
func (mr *MockSubjectActionAlterEventServiceMockRecorder) BulkIncrCheckCount(uuids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkIncrCheckCount", reflect.TypeOf((*MockSubjectActionAlterEventService)(nil).BulkIncrCheckCount), uuids)
}

// BulkUpdateStatus mocks base method.
func (m *MockSubjectActionAlterEventService) BulkUpdateStatus(uuids []string, status int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkUpdateStatus", uuids, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkUpdateStatus indicates an expected call of BulkUpdateStatus.
func (mr *MockSubjectActionAlterEventServiceMockRecorder) BulkUpdateStatus(uuids, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUpdateStatus", reflect.TypeOf((*MockSubjectActionAlterEventService)(nil).BulkUpdateStatus), uuids, status)
}

// Delete mocks base method.
func (m *MockSubjectActionAlterEventService) Delete(uuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", uuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSubjectActionAlterEventServiceMockRecorder) Delete(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSubjectActionAlterEventService)(nil).Delete), uuid)
}

// Get mocks base method.
func (m *MockSubjectActionAlterEventService) Get(uuid string) (types.SubjectActionAlterEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", uuid)
	ret0, _ := ret[0].(types.SubjectActionAlterEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSubjectActionAlterEventServiceMockRecorder) Get(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSubjectActionAlterEventService)(nil).Get), uuid)
}

// ListUUIDByStatusBeforeUpdatedAt mocks base method.
func (m *MockSubjectActionAlterEventService) ListUUIDByStatusBeforeUpdatedAt(status, updateAt int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUUIDByStatusBeforeUpdatedAt", status, updateAt)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUUIDByStatusBeforeUpdatedAt indicates an expected call of ListUUIDByStatusBeforeUpdatedAt.
func (mr *MockSubjectActionAlterEventServiceMockRecorder) ListUUIDByStatusBeforeUpdatedAt(status, updateAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUUIDByStatusBeforeUpdatedAt", reflect.TypeOf((*MockSubjectActionAlterEventService)(nil).ListUUIDByStatusBeforeUpdatedAt), status, updateAt)
}

// ListUUIDGreaterThanStatusLessThanCheckCountBeforeUpdatedAt mocks base method.
func (m *MockSubjectActionAlterEventService) ListUUIDGreaterThanStatusLessThanCheckCountBeforeUpdatedAt(status, checkCount, updateAt int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUUIDGreaterThanStatusLessThanCheckCountBeforeUpdatedAt", status, checkCount, updateAt)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUUIDGreaterThanStatusLessThanCheckCountBeforeUpdatedAt indicates an expected call of ListUUIDGreaterThanStatusLessThanCheckCountBeforeUpdatedAt.
func (mr *MockSubjectActionAlterEventServiceMockRecorder) ListUUIDGreaterThanStatusLessThanCheckCountBeforeUpdatedAt(status, checkCount, updateAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUUIDGreaterThanStatusLessThanCheckCountBeforeUpdatedAt", reflect.TypeOf((*MockSubjectActionAlterEventService)(nil).ListUUIDGreaterThanStatusLessThanCheckCountBeforeUpdatedAt), status, checkCount, updateAt)
}
