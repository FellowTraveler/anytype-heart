// Code generated by mockery v2.39.1. DO NOT EDIT.

package mock_spacecontroller

import (
	context "context"

	mode "github.com/anyproto/anytype-heart/space/internal/spaceprocess/mode"
	mock "github.com/stretchr/testify/mock"

	spaceinfo "github.com/anyproto/anytype-heart/space/spaceinfo"
)

// MockSpaceController is an autogenerated mock type for the SpaceController type
type MockSpaceController struct {
	mock.Mock
}

type MockSpaceController_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSpaceController) EXPECT() *MockSpaceController_Expecter {
	return &MockSpaceController_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: ctx
func (_m *MockSpaceController) Close(ctx context.Context) error {
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

// MockSpaceController_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockSpaceController_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSpaceController_Expecter) Close(ctx interface{}) *MockSpaceController_Close_Call {
	return &MockSpaceController_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *MockSpaceController_Close_Call) Run(run func(ctx context.Context)) *MockSpaceController_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSpaceController_Close_Call) Return(_a0 error) *MockSpaceController_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpaceController_Close_Call) RunAndReturn(run func(context.Context) error) *MockSpaceController_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Current provides a mock function with given fields:
func (_m *MockSpaceController) Current() interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Current")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// MockSpaceController_Current_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Current'
type MockSpaceController_Current_Call struct {
	*mock.Call
}

// Current is a helper method to define mock.On call
func (_e *MockSpaceController_Expecter) Current() *MockSpaceController_Current_Call {
	return &MockSpaceController_Current_Call{Call: _e.mock.On("Current")}
}

func (_c *MockSpaceController_Current_Call) Run(run func()) *MockSpaceController_Current_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSpaceController_Current_Call) Return(_a0 interface{}) *MockSpaceController_Current_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpaceController_Current_Call) RunAndReturn(run func() interface{}) *MockSpaceController_Current_Call {
	_c.Call.Return(run)
	return _c
}

// Mode provides a mock function with given fields:
func (_m *MockSpaceController) Mode() mode.Mode {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Mode")
	}

	var r0 mode.Mode
	if rf, ok := ret.Get(0).(func() mode.Mode); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(mode.Mode)
	}

	return r0
}

// MockSpaceController_Mode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Mode'
type MockSpaceController_Mode_Call struct {
	*mock.Call
}

// Mode is a helper method to define mock.On call
func (_e *MockSpaceController_Expecter) Mode() *MockSpaceController_Mode_Call {
	return &MockSpaceController_Mode_Call{Call: _e.mock.On("Mode")}
}

func (_c *MockSpaceController_Mode_Call) Run(run func()) *MockSpaceController_Mode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSpaceController_Mode_Call) Return(_a0 mode.Mode) *MockSpaceController_Mode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpaceController_Mode_Call) RunAndReturn(run func() mode.Mode) *MockSpaceController_Mode_Call {
	_c.Call.Return(run)
	return _c
}

// SpaceId provides a mock function with given fields:
func (_m *MockSpaceController) SpaceId() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SpaceId")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockSpaceController_SpaceId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SpaceId'
type MockSpaceController_SpaceId_Call struct {
	*mock.Call
}

// SpaceId is a helper method to define mock.On call
func (_e *MockSpaceController_Expecter) SpaceId() *MockSpaceController_SpaceId_Call {
	return &MockSpaceController_SpaceId_Call{Call: _e.mock.On("SpaceId")}
}

func (_c *MockSpaceController_SpaceId_Call) Run(run func()) *MockSpaceController_SpaceId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSpaceController_SpaceId_Call) Return(_a0 string) *MockSpaceController_SpaceId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpaceController_SpaceId_Call) RunAndReturn(run func() string) *MockSpaceController_SpaceId_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx
func (_m *MockSpaceController) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSpaceController_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockSpaceController_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSpaceController_Expecter) Start(ctx interface{}) *MockSpaceController_Start_Call {
	return &MockSpaceController_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *MockSpaceController_Start_Call) Run(run func(ctx context.Context)) *MockSpaceController_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSpaceController_Start_Call) Return(_a0 error) *MockSpaceController_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpaceController_Start_Call) RunAndReturn(run func(context.Context) error) *MockSpaceController_Start_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateRemoteStatus provides a mock function with given fields: ctx, status
func (_m *MockSpaceController) UpdateRemoteStatus(ctx context.Context, status spaceinfo.RemoteStatus) error {
	ret := _m.Called(ctx, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRemoteStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, spaceinfo.RemoteStatus) error); ok {
		r0 = rf(ctx, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSpaceController_UpdateRemoteStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRemoteStatus'
type MockSpaceController_UpdateRemoteStatus_Call struct {
	*mock.Call
}

// UpdateRemoteStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - status spaceinfo.RemoteStatus
func (_e *MockSpaceController_Expecter) UpdateRemoteStatus(ctx interface{}, status interface{}) *MockSpaceController_UpdateRemoteStatus_Call {
	return &MockSpaceController_UpdateRemoteStatus_Call{Call: _e.mock.On("UpdateRemoteStatus", ctx, status)}
}

func (_c *MockSpaceController_UpdateRemoteStatus_Call) Run(run func(ctx context.Context, status spaceinfo.RemoteStatus)) *MockSpaceController_UpdateRemoteStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(spaceinfo.RemoteStatus))
	})
	return _c
}

func (_c *MockSpaceController_UpdateRemoteStatus_Call) Return(_a0 error) *MockSpaceController_UpdateRemoteStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpaceController_UpdateRemoteStatus_Call) RunAndReturn(run func(context.Context, spaceinfo.RemoteStatus) error) *MockSpaceController_UpdateRemoteStatus_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: ctx, status
func (_m *MockSpaceController) UpdateStatus(ctx context.Context, status spaceinfo.AccountStatus) error {
	ret := _m.Called(ctx, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, spaceinfo.AccountStatus) error); ok {
		r0 = rf(ctx, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSpaceController_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type MockSpaceController_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - status spaceinfo.AccountStatus
func (_e *MockSpaceController_Expecter) UpdateStatus(ctx interface{}, status interface{}) *MockSpaceController_UpdateStatus_Call {
	return &MockSpaceController_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", ctx, status)}
}

func (_c *MockSpaceController_UpdateStatus_Call) Run(run func(ctx context.Context, status spaceinfo.AccountStatus)) *MockSpaceController_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(spaceinfo.AccountStatus))
	})
	return _c
}

func (_c *MockSpaceController_UpdateStatus_Call) Return(_a0 error) *MockSpaceController_UpdateStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpaceController_UpdateStatus_Call) RunAndReturn(run func(context.Context, spaceinfo.AccountStatus) error) *MockSpaceController_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSpaceController creates a new instance of MockSpaceController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSpaceController(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSpaceController {
	mock := &MockSpaceController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
