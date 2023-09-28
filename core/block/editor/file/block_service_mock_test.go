// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/anytype-heart/core/block/editor/file (interfaces: BlockService)
//
// Generated by this command:
//
//	mockgen -package file_test -destination block_service_mock_test.go github.com/anyproto/anytype-heart/core/block/editor/file BlockService
//
// Package file_test is a generated GoMock package.
package file_test

import (
	reflect "reflect"

	file "github.com/anyproto/anytype-heart/core/block/editor/file"
	smartblock "github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	process "github.com/anyproto/anytype-heart/core/block/process"
	session "github.com/anyproto/anytype-heart/core/session"
	pb "github.com/anyproto/anytype-heart/pb"
	gomock "go.uber.org/mock/gomock"
)

// MockBlockService is a mock of BlockService interface.
type MockBlockService struct {
	ctrl     *gomock.Controller
	recorder *MockBlockServiceMockRecorder
}

// MockBlockServiceMockRecorder is the mock recorder for MockBlockService.
type MockBlockServiceMockRecorder struct {
	mock *MockBlockService
}

// NewMockBlockService creates a new mock instance.
func NewMockBlockService(ctrl *gomock.Controller) *MockBlockService {
	mock := &MockBlockService{ctrl: ctrl}
	mock.recorder = &MockBlockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockService) EXPECT() *MockBlockServiceMockRecorder {
	return m.recorder
}

// CreateLinkToTheNewObject mocks base method.
func (m *MockBlockService) CreateLinkToTheNewObject(arg0 *session.Context, arg1 *pb.RpcBlockLinkCreateWithObjectRequest) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLinkToTheNewObject", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateLinkToTheNewObject indicates an expected call of CreateLinkToTheNewObject.
func (mr *MockBlockServiceMockRecorder) CreateLinkToTheNewObject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLinkToTheNewObject", reflect.TypeOf((*MockBlockService)(nil).CreateLinkToTheNewObject), arg0, arg1)
}

// Do mocks base method.
func (m *MockBlockService) Do(arg0 string, arg1 func(smartblock.SmartBlock) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do.
func (mr *MockBlockServiceMockRecorder) Do(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockBlockService)(nil).Do), arg0, arg1)
}

// DoFile mocks base method.
func (m *MockBlockService) DoFile(arg0 string, arg1 func(file.File) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoFile indicates an expected call of DoFile.
func (mr *MockBlockServiceMockRecorder) DoFile(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoFile", reflect.TypeOf((*MockBlockService)(nil).DoFile), arg0, arg1)
}

// ProcessAdd mocks base method.
func (m *MockBlockService) ProcessAdd(arg0 process.Process) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessAdd", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessAdd indicates an expected call of ProcessAdd.
func (mr *MockBlockServiceMockRecorder) ProcessAdd(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessAdd", reflect.TypeOf((*MockBlockService)(nil).ProcessAdd), arg0)
}
