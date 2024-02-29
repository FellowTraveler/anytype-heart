// Code generated by mockery. DO NOT EDIT.

package mock_spacecore

import (
	context "context"

	app "github.com/anyproto/any-sync/app"

	mock "github.com/stretchr/testify/mock"

	spacecore "github.com/anyproto/anytype-heart/space/spacecore"

	streampool "github.com/anyproto/any-sync/net/streampool"
)

// MockSpaceCoreService is an autogenerated mock type for the SpaceCoreService type
type MockSpaceCoreService struct {
	mock.Mock
}

type MockSpaceCoreService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSpaceCoreService) EXPECT() *MockSpaceCoreService_Expecter {
	return &MockSpaceCoreService_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: ctx
func (_m *MockSpaceCoreService) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSpaceCoreService_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockSpaceCoreService_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSpaceCoreService_Expecter) Close(ctx interface{}) *MockSpaceCoreService_Close_Call {
	return &MockSpaceCoreService_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *MockSpaceCoreService_Close_Call) Run(run func(ctx context.Context)) *MockSpaceCoreService_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSpaceCoreService_Close_Call) Return(err error) *MockSpaceCoreService_Close_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockSpaceCoreService_Close_Call) RunAndReturn(run func(context.Context) error) *MockSpaceCoreService_Close_Call {
	_c.Call.Return(run)
	return _c
}

// CloseSpace provides a mock function with given fields: ctx, id
func (_m *MockSpaceCoreService) CloseSpace(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for CloseSpace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSpaceCoreService_CloseSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSpace'
type MockSpaceCoreService_CloseSpace_Call struct {
	*mock.Call
}

// CloseSpace is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockSpaceCoreService_Expecter) CloseSpace(ctx interface{}, id interface{}) *MockSpaceCoreService_CloseSpace_Call {
	return &MockSpaceCoreService_CloseSpace_Call{Call: _e.mock.On("CloseSpace", ctx, id)}
}

func (_c *MockSpaceCoreService_CloseSpace_Call) Run(run func(ctx context.Context, id string)) *MockSpaceCoreService_CloseSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSpaceCoreService_CloseSpace_Call) Return(_a0 error) *MockSpaceCoreService_CloseSpace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpaceCoreService_CloseSpace_Call) RunAndReturn(run func(context.Context, string) error) *MockSpaceCoreService_CloseSpace_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, replicationKey, metadataPayload
func (_m *MockSpaceCoreService) Create(ctx context.Context, replicationKey uint64, metadataPayload []byte) (*spacecore.AnySpace, error) {
	ret := _m.Called(ctx, replicationKey, metadataPayload)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *spacecore.AnySpace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, []byte) (*spacecore.AnySpace, error)); ok {
		return rf(ctx, replicationKey, metadataPayload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, []byte) *spacecore.AnySpace); ok {
		r0 = rf(ctx, replicationKey, metadataPayload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*spacecore.AnySpace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, []byte) error); ok {
		r1 = rf(ctx, replicationKey, metadataPayload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceCoreService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockSpaceCoreService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - replicationKey uint64
//   - metadataPayload []byte
func (_e *MockSpaceCoreService_Expecter) Create(ctx interface{}, replicationKey interface{}, metadataPayload interface{}) *MockSpaceCoreService_Create_Call {
	return &MockSpaceCoreService_Create_Call{Call: _e.mock.On("Create", ctx, replicationKey, metadataPayload)}
}

func (_c *MockSpaceCoreService_Create_Call) Run(run func(ctx context.Context, replicationKey uint64, metadataPayload []byte)) *MockSpaceCoreService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].([]byte))
	})
	return _c
}

func (_c *MockSpaceCoreService_Create_Call) Return(_a0 *spacecore.AnySpace, _a1 error) *MockSpaceCoreService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSpaceCoreService_Create_Call) RunAndReturn(run func(context.Context, uint64, []byte) (*spacecore.AnySpace, error)) *MockSpaceCoreService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, spaceID
func (_m *MockSpaceCoreService) Delete(ctx context.Context, spaceID string) error {
	ret := _m.Called(ctx, spaceID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, spaceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSpaceCoreService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockSpaceCoreService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceID string
func (_e *MockSpaceCoreService_Expecter) Delete(ctx interface{}, spaceID interface{}) *MockSpaceCoreService_Delete_Call {
	return &MockSpaceCoreService_Delete_Call{Call: _e.mock.On("Delete", ctx, spaceID)}
}

func (_c *MockSpaceCoreService_Delete_Call) Run(run func(ctx context.Context, spaceID string)) *MockSpaceCoreService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSpaceCoreService_Delete_Call) Return(err error) *MockSpaceCoreService_Delete_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockSpaceCoreService_Delete_Call) RunAndReturn(run func(context.Context, string) error) *MockSpaceCoreService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Derive provides a mock function with given fields: ctx, spaceType
func (_m *MockSpaceCoreService) Derive(ctx context.Context, spaceType string) (*spacecore.AnySpace, error) {
	ret := _m.Called(ctx, spaceType)

	if len(ret) == 0 {
		panic("no return value specified for Derive")
	}

	var r0 *spacecore.AnySpace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*spacecore.AnySpace, error)); ok {
		return rf(ctx, spaceType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *spacecore.AnySpace); ok {
		r0 = rf(ctx, spaceType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*spacecore.AnySpace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, spaceType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceCoreService_Derive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Derive'
type MockSpaceCoreService_Derive_Call struct {
	*mock.Call
}

// Derive is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceType string
func (_e *MockSpaceCoreService_Expecter) Derive(ctx interface{}, spaceType interface{}) *MockSpaceCoreService_Derive_Call {
	return &MockSpaceCoreService_Derive_Call{Call: _e.mock.On("Derive", ctx, spaceType)}
}

func (_c *MockSpaceCoreService_Derive_Call) Run(run func(ctx context.Context, spaceType string)) *MockSpaceCoreService_Derive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSpaceCoreService_Derive_Call) Return(space *spacecore.AnySpace, err error) *MockSpaceCoreService_Derive_Call {
	_c.Call.Return(space, err)
	return _c
}

func (_c *MockSpaceCoreService_Derive_Call) RunAndReturn(run func(context.Context, string) (*spacecore.AnySpace, error)) *MockSpaceCoreService_Derive_Call {
	_c.Call.Return(run)
	return _c
}

// DeriveID provides a mock function with given fields: ctx, spaceType
func (_m *MockSpaceCoreService) DeriveID(ctx context.Context, spaceType string) (string, error) {
	ret := _m.Called(ctx, spaceType)

	if len(ret) == 0 {
		panic("no return value specified for DeriveID")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, spaceType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, spaceType)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, spaceType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceCoreService_DeriveID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeriveID'
type MockSpaceCoreService_DeriveID_Call struct {
	*mock.Call
}

// DeriveID is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceType string
func (_e *MockSpaceCoreService_Expecter) DeriveID(ctx interface{}, spaceType interface{}) *MockSpaceCoreService_DeriveID_Call {
	return &MockSpaceCoreService_DeriveID_Call{Call: _e.mock.On("DeriveID", ctx, spaceType)}
}

func (_c *MockSpaceCoreService_DeriveID_Call) Run(run func(ctx context.Context, spaceType string)) *MockSpaceCoreService_DeriveID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSpaceCoreService_DeriveID_Call) Return(id string, err error) *MockSpaceCoreService_DeriveID_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockSpaceCoreService_DeriveID_Call) RunAndReturn(run func(context.Context, string) (string, error)) *MockSpaceCoreService_DeriveID_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *MockSpaceCoreService) Get(ctx context.Context, id string) (*spacecore.AnySpace, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *spacecore.AnySpace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*spacecore.AnySpace, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *spacecore.AnySpace); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*spacecore.AnySpace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceCoreService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockSpaceCoreService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockSpaceCoreService_Expecter) Get(ctx interface{}, id interface{}) *MockSpaceCoreService_Get_Call {
	return &MockSpaceCoreService_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *MockSpaceCoreService_Get_Call) Run(run func(ctx context.Context, id string)) *MockSpaceCoreService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSpaceCoreService_Get_Call) Return(_a0 *spacecore.AnySpace, _a1 error) *MockSpaceCoreService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSpaceCoreService_Get_Call) RunAndReturn(run func(context.Context, string) (*spacecore.AnySpace, error)) *MockSpaceCoreService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockSpaceCoreService) Init(a *app.App) error {
	ret := _m.Called(a)

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.App) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSpaceCoreService_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockSpaceCoreService_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockSpaceCoreService_Expecter) Init(a interface{}) *MockSpaceCoreService_Init_Call {
	return &MockSpaceCoreService_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockSpaceCoreService_Init_Call) Run(run func(a *app.App)) *MockSpaceCoreService_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockSpaceCoreService_Init_Call) Return(err error) *MockSpaceCoreService_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockSpaceCoreService_Init_Call) RunAndReturn(run func(*app.App) error) *MockSpaceCoreService_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockSpaceCoreService) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockSpaceCoreService_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockSpaceCoreService_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockSpaceCoreService_Expecter) Name() *MockSpaceCoreService_Name_Call {
	return &MockSpaceCoreService_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockSpaceCoreService_Name_Call) Run(run func()) *MockSpaceCoreService_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSpaceCoreService_Name_Call) Return(name string) *MockSpaceCoreService_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockSpaceCoreService_Name_Call) RunAndReturn(run func() string) *MockSpaceCoreService_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Pick provides a mock function with given fields: ctx, id
func (_m *MockSpaceCoreService) Pick(ctx context.Context, id string) (*spacecore.AnySpace, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Pick")
	}

	var r0 *spacecore.AnySpace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*spacecore.AnySpace, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *spacecore.AnySpace); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*spacecore.AnySpace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceCoreService_Pick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Pick'
type MockSpaceCoreService_Pick_Call struct {
	*mock.Call
}

// Pick is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockSpaceCoreService_Expecter) Pick(ctx interface{}, id interface{}) *MockSpaceCoreService_Pick_Call {
	return &MockSpaceCoreService_Pick_Call{Call: _e.mock.On("Pick", ctx, id)}
}

func (_c *MockSpaceCoreService_Pick_Call) Run(run func(ctx context.Context, id string)) *MockSpaceCoreService_Pick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSpaceCoreService_Pick_Call) Return(_a0 *spacecore.AnySpace, _a1 error) *MockSpaceCoreService_Pick_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSpaceCoreService_Pick_Call) RunAndReturn(run func(context.Context, string) (*spacecore.AnySpace, error)) *MockSpaceCoreService_Pick_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: ctx
func (_m *MockSpaceCoreService) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSpaceCoreService_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockSpaceCoreService_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSpaceCoreService_Expecter) Run(ctx interface{}) *MockSpaceCoreService_Run_Call {
	return &MockSpaceCoreService_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *MockSpaceCoreService_Run_Call) Run(run func(ctx context.Context)) *MockSpaceCoreService_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSpaceCoreService_Run_Call) Return(err error) *MockSpaceCoreService_Run_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockSpaceCoreService_Run_Call) RunAndReturn(run func(context.Context) error) *MockSpaceCoreService_Run_Call {
	_c.Call.Return(run)
	return _c
}

// StreamPool provides a mock function with given fields:
func (_m *MockSpaceCoreService) StreamPool() streampool.StreamPool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StreamPool")
	}

	var r0 streampool.StreamPool
	if rf, ok := ret.Get(0).(func() streampool.StreamPool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(streampool.StreamPool)
		}
	}

	return r0
}

// MockSpaceCoreService_StreamPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StreamPool'
type MockSpaceCoreService_StreamPool_Call struct {
	*mock.Call
}

// StreamPool is a helper method to define mock.On call
func (_e *MockSpaceCoreService_Expecter) StreamPool() *MockSpaceCoreService_StreamPool_Call {
	return &MockSpaceCoreService_StreamPool_Call{Call: _e.mock.On("StreamPool")}
}

func (_c *MockSpaceCoreService_StreamPool_Call) Run(run func()) *MockSpaceCoreService_StreamPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSpaceCoreService_StreamPool_Call) Return(_a0 streampool.StreamPool) *MockSpaceCoreService_StreamPool_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpaceCoreService_StreamPool_Call) RunAndReturn(run func() streampool.StreamPool) *MockSpaceCoreService_StreamPool_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSpaceCoreService creates a new instance of MockSpaceCoreService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSpaceCoreService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSpaceCoreService {
	mock := &MockSpaceCoreService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
