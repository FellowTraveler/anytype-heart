// Code generated by mockery. DO NOT EDIT.

package mock_clientserver

import (
	app "github.com/anyproto/any-sync/app"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockClientServer is an autogenerated mock type for the ClientServer type
type MockClientServer struct {
	mock.Mock
}

type MockClientServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClientServer) EXPECT() *MockClientServer_Expecter {
	return &MockClientServer_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: ctx
func (_m *MockClientServer) Close(ctx context.Context) error {
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

// MockClientServer_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockClientServer_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockClientServer_Expecter) Close(ctx interface{}) *MockClientServer_Close_Call {
	return &MockClientServer_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *MockClientServer_Close_Call) Run(run func(ctx context.Context)) *MockClientServer_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockClientServer_Close_Call) Return(err error) *MockClientServer_Close_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockClientServer_Close_Call) RunAndReturn(run func(context.Context) error) *MockClientServer_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockClientServer) Init(a *app.App) error {
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

// MockClientServer_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockClientServer_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockClientServer_Expecter) Init(a interface{}) *MockClientServer_Init_Call {
	return &MockClientServer_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockClientServer_Init_Call) Run(run func(a *app.App)) *MockClientServer_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockClientServer_Init_Call) Return(err error) *MockClientServer_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockClientServer_Init_Call) RunAndReturn(run func(*app.App) error) *MockClientServer_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockClientServer) Name() string {
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

// MockClientServer_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockClientServer_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockClientServer_Expecter) Name() *MockClientServer_Name_Call {
	return &MockClientServer_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockClientServer_Name_Call) Run(run func()) *MockClientServer_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClientServer_Name_Call) Return(name string) *MockClientServer_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockClientServer_Name_Call) RunAndReturn(run func() string) *MockClientServer_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Port provides a mock function with given fields:
func (_m *MockClientServer) Port() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Port")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockClientServer_Port_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Port'
type MockClientServer_Port_Call struct {
	*mock.Call
}

// Port is a helper method to define mock.On call
func (_e *MockClientServer_Expecter) Port() *MockClientServer_Port_Call {
	return &MockClientServer_Port_Call{Call: _e.mock.On("Port")}
}

func (_c *MockClientServer_Port_Call) Run(run func()) *MockClientServer_Port_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClientServer_Port_Call) Return(_a0 int) *MockClientServer_Port_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClientServer_Port_Call) RunAndReturn(run func() int) *MockClientServer_Port_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: ctx
func (_m *MockClientServer) Run(ctx context.Context) error {
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

// MockClientServer_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockClientServer_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockClientServer_Expecter) Run(ctx interface{}) *MockClientServer_Run_Call {
	return &MockClientServer_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *MockClientServer_Run_Call) Run(run func(ctx context.Context)) *MockClientServer_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockClientServer_Run_Call) Return(err error) *MockClientServer_Run_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockClientServer_Run_Call) RunAndReturn(run func(context.Context) error) *MockClientServer_Run_Call {
	_c.Call.Return(run)
	return _c
}

// ServerStarted provides a mock function with given fields:
func (_m *MockClientServer) ServerStarted() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ServerStarted")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockClientServer_ServerStarted_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServerStarted'
type MockClientServer_ServerStarted_Call struct {
	*mock.Call
}

// ServerStarted is a helper method to define mock.On call
func (_e *MockClientServer_Expecter) ServerStarted() *MockClientServer_ServerStarted_Call {
	return &MockClientServer_ServerStarted_Call{Call: _e.mock.On("ServerStarted")}
}

func (_c *MockClientServer_ServerStarted_Call) Run(run func()) *MockClientServer_ServerStarted_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClientServer_ServerStarted_Call) Return(_a0 bool) *MockClientServer_ServerStarted_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClientServer_ServerStarted_Call) RunAndReturn(run func() bool) *MockClientServer_ServerStarted_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClientServer creates a new instance of MockClientServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClientServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClientServer {
	mock := &MockClientServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
