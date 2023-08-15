// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/anytype-heart/pkg/lib/core (interfaces: Service)

// Package testMock is a generated GoMock package.
package testMock

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	uniquekey "github.com/anyproto/anytype-heart/core/block/uniquekey"
	core "github.com/anyproto/anytype-heart/pkg/lib/core"
	threads "github.com/anyproto/anytype-heart/pkg/lib/threads"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AccountObjects mocks base method.
func (m *MockService) AccountObjects() threads.DerivedSmartblockIds {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountObjects")
	ret0, _ := ret[0].(threads.DerivedSmartblockIds)
	return ret0
}

// AccountObjects indicates an expected call of AccountObjects.
func (mr *MockServiceMockRecorder) AccountObjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountObjects", reflect.TypeOf((*MockService)(nil).AccountObjects))
}

// Close mocks base method.
func (m *MockService) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockServiceMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockService)(nil).Close), arg0)
}

// DeriveObjectId mocks base method.
func (m *MockService) DeriveObjectId(arg0 context.Context, arg1 string, arg2 uniquekey.UniqueKey) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeriveObjectId", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeriveObjectId indicates an expected call of DeriveObjectId.
func (mr *MockServiceMockRecorder) DeriveObjectId(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeriveObjectId", reflect.TypeOf((*MockService)(nil).DeriveObjectId), arg0, arg1, arg2)
}

// DerivePredefinedObjects mocks base method.
func (m *MockService) DerivePredefinedObjects(arg0 context.Context, arg1 string, arg2 bool) (threads.DerivedSmartblockIds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DerivePredefinedObjects", arg0, arg1, arg2)
	ret0, _ := ret[0].(threads.DerivedSmartblockIds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DerivePredefinedObjects indicates an expected call of DerivePredefinedObjects.
func (mr *MockServiceMockRecorder) DerivePredefinedObjects(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DerivePredefinedObjects", reflect.TypeOf((*MockService)(nil).DerivePredefinedObjects), arg0, arg1, arg2)
}

// EnsurePredefinedBlocks mocks base method.
func (m *MockService) EnsurePredefinedBlocks(arg0 context.Context, arg1 string) (threads.DerivedSmartblockIds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsurePredefinedBlocks", arg0, arg1)
	ret0, _ := ret[0].(threads.DerivedSmartblockIds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsurePredefinedBlocks indicates an expected call of EnsurePredefinedBlocks.
func (mr *MockServiceMockRecorder) EnsurePredefinedBlocks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsurePredefinedBlocks", reflect.TypeOf((*MockService)(nil).EnsurePredefinedBlocks), arg0, arg1)
}

// GetAllWorkspaces mocks base method.
func (m *MockService) GetAllWorkspaces() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllWorkspaces")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllWorkspaces indicates an expected call of GetAllWorkspaces.
func (mr *MockServiceMockRecorder) GetAllWorkspaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllWorkspaces", reflect.TypeOf((*MockService)(nil).GetAllWorkspaces))
}

// GetWorkspaceIdForObject mocks base method.
func (m *MockService) GetWorkspaceIdForObject(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaceIdForObject", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspaceIdForObject indicates an expected call of GetWorkspaceIdForObject.
func (mr *MockServiceMockRecorder) GetWorkspaceIdForObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaceIdForObject", reflect.TypeOf((*MockService)(nil).GetWorkspaceIdForObject), arg0, arg1)
}

// Init mocks base method.
func (m *MockService) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockServiceMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockService)(nil).Init), arg0)
}

// IsStarted mocks base method.
func (m *MockService) IsStarted() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsStarted")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsStarted indicates an expected call of IsStarted.
func (mr *MockServiceMockRecorder) IsStarted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStarted", reflect.TypeOf((*MockService)(nil).IsStarted))
}

// LocalProfile mocks base method.
func (m *MockService) LocalProfile(arg0 string) (core.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalProfile", arg0)
	ret0, _ := ret[0].(core.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocalProfile indicates an expected call of LocalProfile.
func (mr *MockServiceMockRecorder) LocalProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalProfile", reflect.TypeOf((*MockService)(nil).LocalProfile), arg0)
}

// Name mocks base method.
func (m *MockService) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockServiceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockService)(nil).Name))
}

// PredefinedObjects mocks base method.
func (m *MockService) PredefinedObjects(arg0 string) threads.DerivedSmartblockIds {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PredefinedObjects", arg0)
	ret0, _ := ret[0].(threads.DerivedSmartblockIds)
	return ret0
}

// PredefinedObjects indicates an expected call of PredefinedObjects.
func (mr *MockServiceMockRecorder) PredefinedObjects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PredefinedObjects", reflect.TypeOf((*MockService)(nil).PredefinedObjects), arg0)
}

// ProfileID mocks base method.
func (m *MockService) ProfileID(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProfileID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// ProfileID indicates an expected call of ProfileID.
func (mr *MockServiceMockRecorder) ProfileID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProfileID", reflect.TypeOf((*MockService)(nil).ProfileID), arg0)
}

// Run mocks base method.
func (m *MockService) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockServiceMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockService)(nil).Run), arg0)
}

// Stop mocks base method.
func (m *MockService) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockServiceMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockService)(nil).Stop))
}
