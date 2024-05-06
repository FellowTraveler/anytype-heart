// Code generated by mockery. DO NOT EDIT.

package mock_syncer

import (
	context "context"

	block "github.com/anyproto/anytype-heart/core/block"

	domain "github.com/anyproto/anytype-heart/core/domain"

	mock "github.com/stretchr/testify/mock"

	smartblock "github.com/anyproto/anytype-heart/core/block/editor/smartblock"

	types "github.com/gogo/protobuf/types"
)

// MockBlockService is an autogenerated mock type for the BlockService type
type MockBlockService struct {
	mock.Mock
}

type MockBlockService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBlockService) EXPECT() *MockBlockService_Expecter {
	return &MockBlockService_Expecter{mock: &_m.Mock}
}

// GetObject provides a mock function with given fields: ctx, objectID
func (_m *MockBlockService) GetObject(ctx context.Context, objectID string) (smartblock.SmartBlock, error) {
	ret := _m.Called(ctx, objectID)

	if len(ret) == 0 {
		panic("no return value specified for GetObject")
	}

	var r0 smartblock.SmartBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (smartblock.SmartBlock, error)); ok {
		return rf(ctx, objectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) smartblock.SmartBlock); ok {
		r0 = rf(ctx, objectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(smartblock.SmartBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, objectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBlockService_GetObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObject'
type MockBlockService_GetObject_Call struct {
	*mock.Call
}

// GetObject is a helper method to define mock.On call
//   - ctx context.Context
//   - objectID string
func (_e *MockBlockService_Expecter) GetObject(ctx interface{}, objectID interface{}) *MockBlockService_GetObject_Call {
	return &MockBlockService_GetObject_Call{Call: _e.mock.On("GetObject", ctx, objectID)}
}

func (_c *MockBlockService_GetObject_Call) Run(run func(ctx context.Context, objectID string)) *MockBlockService_GetObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockBlockService_GetObject_Call) Return(sb smartblock.SmartBlock, err error) *MockBlockService_GetObject_Call {
	_c.Call.Return(sb, err)
	return _c
}

func (_c *MockBlockService_GetObject_Call) RunAndReturn(run func(context.Context, string) (smartblock.SmartBlock, error)) *MockBlockService_GetObject_Call {
	_c.Call.Return(run)
	return _c
}

// GetObjectByFullID provides a mock function with given fields: ctx, id
func (_m *MockBlockService) GetObjectByFullID(ctx context.Context, id domain.FullID) (smartblock.SmartBlock, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetObjectByFullID")
	}

	var r0 smartblock.SmartBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.FullID) (smartblock.SmartBlock, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.FullID) smartblock.SmartBlock); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(smartblock.SmartBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.FullID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBlockService_GetObjectByFullID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObjectByFullID'
type MockBlockService_GetObjectByFullID_Call struct {
	*mock.Call
}

// GetObjectByFullID is a helper method to define mock.On call
//   - ctx context.Context
//   - id domain.FullID
func (_e *MockBlockService_Expecter) GetObjectByFullID(ctx interface{}, id interface{}) *MockBlockService_GetObjectByFullID_Call {
	return &MockBlockService_GetObjectByFullID_Call{Call: _e.mock.On("GetObjectByFullID", ctx, id)}
}

func (_c *MockBlockService_GetObjectByFullID_Call) Run(run func(ctx context.Context, id domain.FullID)) *MockBlockService_GetObjectByFullID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.FullID))
	})
	return _c
}

func (_c *MockBlockService_GetObjectByFullID_Call) Return(sb smartblock.SmartBlock, err error) *MockBlockService_GetObjectByFullID_Call {
	_c.Call.Return(sb, err)
	return _c
}

func (_c *MockBlockService_GetObjectByFullID_Call) RunAndReturn(run func(context.Context, domain.FullID) (smartblock.SmartBlock, error)) *MockBlockService_GetObjectByFullID_Call {
	_c.Call.Return(run)
	return _c
}

// UploadFile provides a mock function with given fields: ctx, spaceId, req
func (_m *MockBlockService) UploadFile(ctx context.Context, spaceId string, req block.FileUploadRequest) (string, *types.Struct, error) {
	ret := _m.Called(ctx, spaceId, req)

	if len(ret) == 0 {
		panic("no return value specified for UploadFile")
	}

	var r0 string
	var r1 *types.Struct
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, block.FileUploadRequest) (string, *types.Struct, error)); ok {
		return rf(ctx, spaceId, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, block.FileUploadRequest) string); ok {
		r0 = rf(ctx, spaceId, req)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, block.FileUploadRequest) *types.Struct); ok {
		r1 = rf(ctx, spaceId, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Struct)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, block.FileUploadRequest) error); ok {
		r2 = rf(ctx, spaceId, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockBlockService_UploadFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadFile'
type MockBlockService_UploadFile_Call struct {
	*mock.Call
}

// UploadFile is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
//   - req block.FileUploadRequest
func (_e *MockBlockService_Expecter) UploadFile(ctx interface{}, spaceId interface{}, req interface{}) *MockBlockService_UploadFile_Call {
	return &MockBlockService_UploadFile_Call{Call: _e.mock.On("UploadFile", ctx, spaceId, req)}
}

func (_c *MockBlockService_UploadFile_Call) Run(run func(ctx context.Context, spaceId string, req block.FileUploadRequest)) *MockBlockService_UploadFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(block.FileUploadRequest))
	})
	return _c
}

func (_c *MockBlockService_UploadFile_Call) Return(objectId string, details *types.Struct, err error) *MockBlockService_UploadFile_Call {
	_c.Call.Return(objectId, details, err)
	return _c
}

func (_c *MockBlockService_UploadFile_Call) RunAndReturn(run func(context.Context, string, block.FileUploadRequest) (string, *types.Struct, error)) *MockBlockService_UploadFile_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBlockService creates a new instance of MockBlockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBlockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBlockService {
	mock := &MockBlockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
