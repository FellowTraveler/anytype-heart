// Code generated by mockery v2.26.1. DO NOT EDIT.

package mock_fileservice

import (
	context "context"
	io "io"

	app "github.com/anyproto/any-sync/app"
	unixfsio "github.com/ipfs/boxo/ipld/unixfs/io"
	cid "github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	mock "github.com/stretchr/testify/mock"
)

// MockFileService is an autogenerated mock type for the FileService type
type MockFileService struct {
	mock.Mock
}

type MockFileService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFileService) EXPECT() *MockFileService_Expecter {
	return &MockFileService_Expecter{mock: &_m.Mock}
}

// AddFile provides a mock function with given fields: ctx, r
func (_m *MockFileService) AddFile(ctx context.Context, r io.Reader) (format.Node, error) {
	ret := _m.Called(ctx, r)

	var r0 format.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader) (format.Node, error)); ok {
		return rf(ctx, r)
	}
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader) format.Node); ok {
		r0 = rf(ctx, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(format.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, io.Reader) error); ok {
		r1 = rf(ctx, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileService_AddFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFile'
type MockFileService_AddFile_Call struct {
	*mock.Call
}

// AddFile is a helper method to define mock.On call
//   - ctx context.Context
//   - r io.Reader
func (_e *MockFileService_Expecter) AddFile(ctx interface{}, r interface{}) *MockFileService_AddFile_Call {
	return &MockFileService_AddFile_Call{Call: _e.mock.On("AddFile", ctx, r)}
}

func (_c *MockFileService_AddFile_Call) Run(run func(ctx context.Context, r io.Reader)) *MockFileService_AddFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(io.Reader))
	})
	return _c
}

func (_c *MockFileService_AddFile_Call) Return(_a0 format.Node, _a1 error) *MockFileService_AddFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFileService_AddFile_Call) RunAndReturn(run func(context.Context, io.Reader) (format.Node, error)) *MockFileService_AddFile_Call {
	_c.Call.Return(run)
	return _c
}

// DAGService provides a mock function with given fields:
func (_m *MockFileService) DAGService() format.DAGService {
	ret := _m.Called()

	var r0 format.DAGService
	if rf, ok := ret.Get(0).(func() format.DAGService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(format.DAGService)
		}
	}

	return r0
}

// MockFileService_DAGService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DAGService'
type MockFileService_DAGService_Call struct {
	*mock.Call
}

// DAGService is a helper method to define mock.On call
func (_e *MockFileService_Expecter) DAGService() *MockFileService_DAGService_Call {
	return &MockFileService_DAGService_Call{Call: _e.mock.On("DAGService")}
}

func (_c *MockFileService_DAGService_Call) Run(run func()) *MockFileService_DAGService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileService_DAGService_Call) Return(_a0 format.DAGService) *MockFileService_DAGService_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileService_DAGService_Call) RunAndReturn(run func() format.DAGService) *MockFileService_DAGService_Call {
	_c.Call.Return(run)
	return _c
}

// GetFile provides a mock function with given fields: ctx, c
func (_m *MockFileService) GetFile(ctx context.Context, c cid.Cid) (unixfsio.ReadSeekCloser, error) {
	ret := _m.Called(ctx, c)

	var r0 unixfsio.ReadSeekCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cid.Cid) (unixfsio.ReadSeekCloser, error)); ok {
		return rf(ctx, c)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cid.Cid) unixfsio.ReadSeekCloser); ok {
		r0 = rf(ctx, c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(unixfsio.ReadSeekCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cid.Cid) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileService_GetFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFile'
type MockFileService_GetFile_Call struct {
	*mock.Call
}

// GetFile is a helper method to define mock.On call
//   - ctx context.Context
//   - c cid.Cid
func (_e *MockFileService_Expecter) GetFile(ctx interface{}, c interface{}) *MockFileService_GetFile_Call {
	return &MockFileService_GetFile_Call{Call: _e.mock.On("GetFile", ctx, c)}
}

func (_c *MockFileService_GetFile_Call) Run(run func(ctx context.Context, c cid.Cid)) *MockFileService_GetFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cid.Cid))
	})
	return _c
}

func (_c *MockFileService_GetFile_Call) Return(_a0 unixfsio.ReadSeekCloser, _a1 error) *MockFileService_GetFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFileService_GetFile_Call) RunAndReturn(run func(context.Context, cid.Cid) (unixfsio.ReadSeekCloser, error)) *MockFileService_GetFile_Call {
	_c.Call.Return(run)
	return _c
}

// HasCid provides a mock function with given fields: ctx, c
func (_m *MockFileService) HasCid(ctx context.Context, c cid.Cid) (bool, error) {
	ret := _m.Called(ctx, c)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cid.Cid) (bool, error)); ok {
		return rf(ctx, c)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cid.Cid) bool); ok {
		r0 = rf(ctx, c)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, cid.Cid) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileService_HasCid_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasCid'
type MockFileService_HasCid_Call struct {
	*mock.Call
}

// HasCid is a helper method to define mock.On call
//   - ctx context.Context
//   - c cid.Cid
func (_e *MockFileService_Expecter) HasCid(ctx interface{}, c interface{}) *MockFileService_HasCid_Call {
	return &MockFileService_HasCid_Call{Call: _e.mock.On("HasCid", ctx, c)}
}

func (_c *MockFileService_HasCid_Call) Run(run func(ctx context.Context, c cid.Cid)) *MockFileService_HasCid_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cid.Cid))
	})
	return _c
}

func (_c *MockFileService_HasCid_Call) Return(exists bool, err error) *MockFileService_HasCid_Call {
	_c.Call.Return(exists, err)
	return _c
}

func (_c *MockFileService_HasCid_Call) RunAndReturn(run func(context.Context, cid.Cid) (bool, error)) *MockFileService_HasCid_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockFileService) Init(a *app.App) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.App) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFileService_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockFileService_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockFileService_Expecter) Init(a interface{}) *MockFileService_Init_Call {
	return &MockFileService_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockFileService_Init_Call) Run(run func(a *app.App)) *MockFileService_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockFileService_Init_Call) Return(err error) *MockFileService_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockFileService_Init_Call) RunAndReturn(run func(*app.App) error) *MockFileService_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockFileService) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockFileService_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockFileService_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockFileService_Expecter) Name() *MockFileService_Name_Call {
	return &MockFileService_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockFileService_Name_Call) Run(run func()) *MockFileService_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileService_Name_Call) Return(name string) *MockFileService_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockFileService_Name_Call) RunAndReturn(run func() string) *MockFileService_Name_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockFileService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockFileService creates a new instance of MockFileService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockFileService(t mockConstructorTestingTNewMockFileService) *MockFileService {
	mock := &MockFileService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
