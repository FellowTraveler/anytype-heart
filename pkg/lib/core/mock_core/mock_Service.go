// Code generated by mockery v2.26.1. DO NOT EDIT.

package mock_core

import (
	context "context"

	app "github.com/anyproto/any-sync/app"

	mock "github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

type MockService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockService) EXPECT() *MockService_Expecter {
	return &MockService_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: ctx
func (_m *MockService) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockService_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockService_Expecter) Close(ctx interface{}) *MockService_Close_Call {
	return &MockService_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *MockService_Close_Call) Run(run func(ctx context.Context)) *MockService_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockService_Close_Call) Return(err error) *MockService_Close_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockService_Close_Call) RunAndReturn(run func(context.Context) error) *MockService_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockService) Init(a *app.App) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.App) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockService_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockService_Expecter) Init(a interface{}) *MockService_Init_Call {
	return &MockService_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockService_Init_Call) Run(run func(a *app.App)) *MockService_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockService_Init_Call) Return(err error) *MockService_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockService_Init_Call) RunAndReturn(run func(*app.App) error) *MockService_Init_Call {
	_c.Call.Return(run)
	return _c
}

// IsStarted provides a mock function with given fields:
func (_m *MockService) IsStarted() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockService_IsStarted_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsStarted'
type MockService_IsStarted_Call struct {
	*mock.Call
}

// IsStarted is a helper method to define mock.On call
func (_e *MockService_Expecter) IsStarted() *MockService_IsStarted_Call {
	return &MockService_IsStarted_Call{Call: _e.mock.On("IsStarted")}
}

func (_c *MockService_IsStarted_Call) Run(run func()) *MockService_IsStarted_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_IsStarted_Call) Return(_a0 bool) *MockService_IsStarted_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_IsStarted_Call) RunAndReturn(run func() bool) *MockService_IsStarted_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockService) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockService_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockService_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockService_Expecter) Name() *MockService_Name_Call {
	return &MockService_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockService_Name_Call) Run(run func()) *MockService_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_Name_Call) Return(name string) *MockService_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockService_Name_Call) RunAndReturn(run func() string) *MockService_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: ctx
func (_m *MockService) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockService_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockService_Expecter) Run(ctx interface{}) *MockService_Run_Call {
	return &MockService_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *MockService_Run_Call) Run(run func(ctx context.Context)) *MockService_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockService_Run_Call) Return(err error) *MockService_Run_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockService_Run_Call) RunAndReturn(run func(context.Context) error) *MockService_Run_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockService) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockService_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockService_Expecter) Stop() *MockService_Stop_Call {
	return &MockService_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockService_Stop_Call) Run(run func()) *MockService_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_Stop_Call) Return(_a0 error) *MockService_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockService_Stop_Call) RunAndReturn(run func() error) *MockService_Stop_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockService(t mockConstructorTestingTNewMockService) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
