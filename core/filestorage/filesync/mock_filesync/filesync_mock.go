// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/anytype-heart/core/filestorage/filesync (interfaces: FileSync)

// Package mock_filesync is a generated GoMock package.
package mock_filesync

import (
	context "context"
	http "net/http"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	filesync "github.com/anyproto/anytype-heart/core/filestorage/filesync"
	format "github.com/ipfs/go-ipld-format"
	gomock "go.uber.org/mock/gomock"
)

// MockFileSync is a mock of FileSync interface.
type MockFileSync struct {
	ctrl     *gomock.Controller
	recorder *MockFileSyncMockRecorder
}

// MockFileSyncMockRecorder is the mock recorder for MockFileSync.
type MockFileSyncMockRecorder struct {
	mock *MockFileSync
}

// NewMockFileSync creates a new mock instance.
func NewMockFileSync(ctrl *gomock.Controller) *MockFileSync {
	mock := &MockFileSync{ctrl: ctrl}
	mock.recorder = &MockFileSyncMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileSync) EXPECT() *MockFileSyncMockRecorder {
	return m.recorder
}

// AddFile mocks base method.
func (m *MockFileSync) AddFile(arg0, arg1 string, arg2, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFile", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFile indicates an expected call of AddFile.
func (mr *MockFileSyncMockRecorder) AddFile(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFile", reflect.TypeOf((*MockFileSync)(nil).AddFile), arg0, arg1, arg2, arg3)
}

// ClearImportEvents mocks base method.
func (m *MockFileSync) ClearImportEvents() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearImportEvents")
}

// ClearImportEvents indicates an expected call of ClearImportEvents.
func (mr *MockFileSyncMockRecorder) ClearImportEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearImportEvents", reflect.TypeOf((*MockFileSync)(nil).ClearImportEvents))
}

// Close mocks base method.
func (m *MockFileSync) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockFileSyncMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFileSync)(nil).Close), arg0)
}

// DebugQueue mocks base method.
func (m *MockFileSync) DebugQueue(arg0 *http.Request) (*filesync.QueueInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DebugQueue", arg0)
	ret0, _ := ret[0].(*filesync.QueueInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DebugQueue indicates an expected call of DebugQueue.
func (mr *MockFileSyncMockRecorder) DebugQueue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DebugQueue", reflect.TypeOf((*MockFileSync)(nil).DebugQueue), arg0)
}

// FetchChunksCount mocks base method.
func (m *MockFileSync) FetchChunksCount(arg0 context.Context, arg1 format.Node) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchChunksCount", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchChunksCount indicates an expected call of FetchChunksCount.
func (mr *MockFileSyncMockRecorder) FetchChunksCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchChunksCount", reflect.TypeOf((*MockFileSync)(nil).FetchChunksCount), arg0, arg1)
}

// FileListStats mocks base method.
func (m *MockFileSync) FileListStats(arg0 context.Context, arg1 string, arg2 []string) ([]filesync.FileStat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileListStats", arg0, arg1, arg2)
	ret0, _ := ret[0].([]filesync.FileStat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FileListStats indicates an expected call of FileListStats.
func (mr *MockFileSyncMockRecorder) FileListStats(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileListStats", reflect.TypeOf((*MockFileSync)(nil).FileListStats), arg0, arg1, arg2)
}

// FileStat mocks base method.
func (m *MockFileSync) FileStat(arg0 context.Context, arg1, arg2 string) (filesync.FileStat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileStat", arg0, arg1, arg2)
	ret0, _ := ret[0].(filesync.FileStat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FileStat indicates an expected call of FileStat.
func (mr *MockFileSyncMockRecorder) FileStat(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileStat", reflect.TypeOf((*MockFileSync)(nil).FileStat), arg0, arg1, arg2)
}

// HasUpload mocks base method.
func (m *MockFileSync) HasUpload(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUpload", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasUpload indicates an expected call of HasUpload.
func (mr *MockFileSyncMockRecorder) HasUpload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUpload", reflect.TypeOf((*MockFileSync)(nil).HasUpload), arg0, arg1)
}

// Init mocks base method.
func (m *MockFileSync) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockFileSyncMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockFileSync)(nil).Init), arg0)
}

// IsFileUploadLimited mocks base method.
func (m *MockFileSync) IsFileUploadLimited(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFileUploadLimited", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFileUploadLimited indicates an expected call of IsFileUploadLimited.
func (mr *MockFileSyncMockRecorder) IsFileUploadLimited(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFileUploadLimited", reflect.TypeOf((*MockFileSync)(nil).IsFileUploadLimited), arg0, arg1)
}

// Name mocks base method.
func (m *MockFileSync) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockFileSyncMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockFileSync)(nil).Name))
}

// OnUpload mocks base method.
func (m *MockFileSync) OnUpload(arg0 func(string, string) error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnUpload", arg0)
}

// OnUpload indicates an expected call of OnUpload.
func (mr *MockFileSyncMockRecorder) OnUpload(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnUpload", reflect.TypeOf((*MockFileSync)(nil).OnUpload), arg0)
}

// RemoveFile mocks base method.
func (m *MockFileSync) RemoveFile(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFile indicates an expected call of RemoveFile.
func (mr *MockFileSyncMockRecorder) RemoveFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFile", reflect.TypeOf((*MockFileSync)(nil).RemoveFile), arg0, arg1)
}

// Run mocks base method.
func (m *MockFileSync) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockFileSyncMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockFileSync)(nil).Run), arg0)
}

// SendImportEvents mocks base method.
func (m *MockFileSync) SendImportEvents() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendImportEvents")
}

// SendImportEvents indicates an expected call of SendImportEvents.
func (mr *MockFileSyncMockRecorder) SendImportEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendImportEvents", reflect.TypeOf((*MockFileSync)(nil).SendImportEvents))
}

// SpaceStat mocks base method.
func (m *MockFileSync) SpaceStat(arg0 context.Context, arg1 string) (filesync.SpaceStat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceStat", arg0, arg1)
	ret0, _ := ret[0].(filesync.SpaceStat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpaceStat indicates an expected call of SpaceStat.
func (mr *MockFileSyncMockRecorder) SpaceStat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceStat", reflect.TypeOf((*MockFileSync)(nil).SpaceStat), arg0, arg1)
}

// SyncStatus mocks base method.
func (m *MockFileSync) SyncStatus() (filesync.SyncStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatus")
	ret0, _ := ret[0].(filesync.SyncStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncStatus indicates an expected call of SyncStatus.
func (mr *MockFileSyncMockRecorder) SyncStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatus", reflect.TypeOf((*MockFileSync)(nil).SyncStatus))
}
