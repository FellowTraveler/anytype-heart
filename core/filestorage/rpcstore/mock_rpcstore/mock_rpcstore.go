// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/anytype-heart/core/filestorage/rpcstore (interfaces: Service,RpcStore)
//
// Generated by this command:
//
//	mockgen -destination mock_rpcstore/mock_rpcstore.go github.com/anyproto/anytype-heart/core/filestorage/rpcstore Service,RpcStore
//

// Package mock_rpcstore is a generated GoMock package.
package mock_rpcstore

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	fileproto "github.com/anyproto/any-sync/commonfile/fileproto"
	domain "github.com/anyproto/anytype-heart/core/domain"
	rpcstore "github.com/anyproto/anytype-heart/core/filestorage/rpcstore"
	blocks "github.com/ipfs/go-block-format"
	cid "github.com/ipfs/go-cid"
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

// Init mocks base method.
func (m *MockService) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockServiceMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockService)(nil).Init), arg0)
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

// NewStore mocks base method.
func (m *MockService) NewStore() rpcstore.RpcStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStore")
	ret0, _ := ret[0].(rpcstore.RpcStore)
	return ret0
}

// NewStore indicates an expected call of NewStore.
func (mr *MockServiceMockRecorder) NewStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStore", reflect.TypeOf((*MockService)(nil).NewStore))
}

// MockRpcStore is a mock of RpcStore interface.
type MockRpcStore struct {
	ctrl     *gomock.Controller
	recorder *MockRpcStoreMockRecorder
}

// MockRpcStoreMockRecorder is the mock recorder for MockRpcStore.
type MockRpcStoreMockRecorder struct {
	mock *MockRpcStore
}

// NewMockRpcStore creates a new mock instance.
func NewMockRpcStore(ctrl *gomock.Controller) *MockRpcStore {
	mock := &MockRpcStore{ctrl: ctrl}
	mock.recorder = &MockRpcStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRpcStore) EXPECT() *MockRpcStoreMockRecorder {
	return m.recorder
}

// AccountInfo mocks base method.
func (m *MockRpcStore) AccountInfo(arg0 context.Context) (*fileproto.AccountInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountInfo", arg0)
	ret0, _ := ret[0].(*fileproto.AccountInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountInfo indicates an expected call of AccountInfo.
func (mr *MockRpcStoreMockRecorder) AccountInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountInfo", reflect.TypeOf((*MockRpcStore)(nil).AccountInfo), arg0)
}

// Add mocks base method.
func (m *MockRpcStore) Add(arg0 context.Context, arg1 []blocks.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockRpcStoreMockRecorder) Add(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRpcStore)(nil).Add), arg0, arg1)
}

// AddToFile mocks base method.
func (m *MockRpcStore) AddToFile(arg0 context.Context, arg1 string, arg2 domain.FileId, arg3 []blocks.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToFile", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToFile indicates an expected call of AddToFile.
func (mr *MockRpcStoreMockRecorder) AddToFile(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToFile", reflect.TypeOf((*MockRpcStore)(nil).AddToFile), arg0, arg1, arg2, arg3)
}

// BindCids mocks base method.
func (m *MockRpcStore) BindCids(arg0 context.Context, arg1 string, arg2 domain.FileId, arg3 []cid.Cid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindCids", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindCids indicates an expected call of BindCids.
func (mr *MockRpcStoreMockRecorder) BindCids(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindCids", reflect.TypeOf((*MockRpcStore)(nil).BindCids), arg0, arg1, arg2, arg3)
}

// CheckAvailability mocks base method.
func (m *MockRpcStore) CheckAvailability(arg0 context.Context, arg1 string, arg2 []cid.Cid) ([]*fileproto.BlockAvailability, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAvailability", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*fileproto.BlockAvailability)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAvailability indicates an expected call of CheckAvailability.
func (mr *MockRpcStoreMockRecorder) CheckAvailability(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAvailability", reflect.TypeOf((*MockRpcStore)(nil).CheckAvailability), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockRpcStore) Delete(arg0 context.Context, arg1 cid.Cid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRpcStoreMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRpcStore)(nil).Delete), arg0, arg1)
}

// DeleteFiles mocks base method.
func (m *MockRpcStore) DeleteFiles(arg0 context.Context, arg1 string, arg2 ...domain.FileId) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFiles", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFiles indicates an expected call of DeleteFiles.
func (mr *MockRpcStoreMockRecorder) DeleteFiles(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFiles", reflect.TypeOf((*MockRpcStore)(nil).DeleteFiles), varargs...)
}

// FilesInfo mocks base method.
func (m *MockRpcStore) FilesInfo(arg0 context.Context, arg1 string, arg2 ...domain.FileId) ([]*fileproto.FileInfo, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FilesInfo", varargs...)
	ret0, _ := ret[0].([]*fileproto.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilesInfo indicates an expected call of FilesInfo.
func (mr *MockRpcStoreMockRecorder) FilesInfo(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesInfo", reflect.TypeOf((*MockRpcStore)(nil).FilesInfo), varargs...)
}

// Get mocks base method.
func (m *MockRpcStore) Get(arg0 context.Context, arg1 cid.Cid) (blocks.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(blocks.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRpcStoreMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRpcStore)(nil).Get), arg0, arg1)
}

// GetMany mocks base method.
func (m *MockRpcStore) GetMany(arg0 context.Context, arg1 []cid.Cid) <-chan blocks.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", arg0, arg1)
	ret0, _ := ret[0].(<-chan blocks.Block)
	return ret0
}

// GetMany indicates an expected call of GetMany.
func (mr *MockRpcStoreMockRecorder) GetMany(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockRpcStore)(nil).GetMany), arg0, arg1)
}

// IterateFiles mocks base method.
func (m *MockRpcStore) IterateFiles(arg0 context.Context, arg1 func(domain.FullFileId)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IterateFiles", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IterateFiles indicates an expected call of IterateFiles.
func (mr *MockRpcStoreMockRecorder) IterateFiles(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateFiles", reflect.TypeOf((*MockRpcStore)(nil).IterateFiles), arg0, arg1)
}

// SpaceInfo mocks base method.
func (m *MockRpcStore) SpaceInfo(arg0 context.Context, arg1 string) (*fileproto.SpaceInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceInfo", arg0, arg1)
	ret0, _ := ret[0].(*fileproto.SpaceInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpaceInfo indicates an expected call of SpaceInfo.
func (mr *MockRpcStoreMockRecorder) SpaceInfo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceInfo", reflect.TypeOf((*MockRpcStore)(nil).SpaceInfo), arg0, arg1)
}
