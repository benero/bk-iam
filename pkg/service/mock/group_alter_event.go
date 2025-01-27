// Code generated by MockGen. DO NOT EDIT.
// Source: group_alter_event.go

// Package mock is a generated GoMock package.
package mock

import (
	types "github.com/TencentBlueKing/bk-iam/pkg/service/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockGroupAlterEventService is a mock of GroupAlterEventService interface.
type MockGroupAlterEventService struct {
	ctrl     *gomock.Controller
	recorder *MockGroupAlterEventServiceMockRecorder
}

// MockGroupAlterEventServiceMockRecorder is the mock recorder for MockGroupAlterEventService.
type MockGroupAlterEventServiceMockRecorder struct {
	mock *MockGroupAlterEventService
}

// NewMockGroupAlterEventService creates a new mock instance.
func NewMockGroupAlterEventService(ctrl *gomock.Controller) *MockGroupAlterEventService {
	mock := &MockGroupAlterEventService{ctrl: ctrl}
	mock.recorder = &MockGroupAlterEventServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupAlterEventService) EXPECT() *MockGroupAlterEventServiceMockRecorder {
	return m.recorder
}

// BulkDeleteWithTx mocks base method.
func (m *MockGroupAlterEventService) BulkDeleteWithTx(tx *sqlx.Tx, uuids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteWithTx", tx, uuids)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteWithTx indicates an expected call of BulkDeleteWithTx.
func (mr *MockGroupAlterEventServiceMockRecorder) BulkDeleteWithTx(tx, uuids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteWithTx", reflect.TypeOf((*MockGroupAlterEventService)(nil).BulkDeleteWithTx), tx, uuids)
}

// CreateByGroupAction mocks base method.
func (m *MockGroupAlterEventService) CreateByGroupAction(groupPK int64, actionPKs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateByGroupAction", groupPK, actionPKs)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateByGroupAction indicates an expected call of CreateByGroupAction.
func (mr *MockGroupAlterEventServiceMockRecorder) CreateByGroupAction(groupPK, actionPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateByGroupAction", reflect.TypeOf((*MockGroupAlterEventService)(nil).CreateByGroupAction), groupPK, actionPKs)
}

// CreateByGroupSubject mocks base method.
func (m *MockGroupAlterEventService) CreateByGroupSubject(groupPK int64, subjectPKs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateByGroupSubject", groupPK, subjectPKs)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateByGroupSubject indicates an expected call of CreateByGroupSubject.
func (mr *MockGroupAlterEventServiceMockRecorder) CreateByGroupSubject(groupPK, subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateByGroupSubject", reflect.TypeOf((*MockGroupAlterEventService)(nil).CreateByGroupSubject), groupPK, subjectPKs)
}

// CreateBySubjectActionGroup mocks base method.
func (m *MockGroupAlterEventService) CreateBySubjectActionGroup(subjectPK, actionPK, groupPK int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBySubjectActionGroup", subjectPK, actionPK, groupPK)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBySubjectActionGroup indicates an expected call of CreateBySubjectActionGroup.
func (mr *MockGroupAlterEventServiceMockRecorder) CreateBySubjectActionGroup(subjectPK, actionPK, groupPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBySubjectActionGroup", reflect.TypeOf((*MockGroupAlterEventService)(nil).CreateBySubjectActionGroup), subjectPK, actionPK, groupPK)
}

// ListBeforeCreateAt mocks base method.
func (m *MockGroupAlterEventService) ListBeforeCreateAt(createdAt, limit int64) ([]types.GroupAlterEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBeforeCreateAt", createdAt, limit)
	ret0, _ := ret[0].([]types.GroupAlterEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBeforeCreateAt indicates an expected call of ListBeforeCreateAt.
func (mr *MockGroupAlterEventServiceMockRecorder) ListBeforeCreateAt(createdAt, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBeforeCreateAt", reflect.TypeOf((*MockGroupAlterEventService)(nil).ListBeforeCreateAt), createdAt, limit)
}
