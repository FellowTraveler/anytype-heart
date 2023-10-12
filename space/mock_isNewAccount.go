// Code generated by mockery v2.26.1. DO NOT EDIT.

package space

import (
	app "github.com/anyproto/any-sync/app"
	mock "github.com/stretchr/testify/mock"
)

// MockisNewAccount is an autogenerated mock type for the isNewAccount type
type MockisNewAccount struct {
	mock.Mock
}

type MockisNewAccount_Expecter struct {
	mock *mock.Mock
}

func (_m *MockisNewAccount) EXPECT() *MockisNewAccount_Expecter {
	return &MockisNewAccount_Expecter{mock: &_m.Mock}
}

// Init provides a mock function with given fields: a
func (_m *MockisNewAccount) Init(a *app.App) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.App) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockisNewAccount_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockisNewAccount_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockisNewAccount_Expecter) Init(a interface{}) *MockisNewAccount_Init_Call {
	return &MockisNewAccount_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockisNewAccount_Init_Call) Run(run func(a *app.App)) *MockisNewAccount_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockisNewAccount_Init_Call) Return(err error) *MockisNewAccount_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockisNewAccount_Init_Call) RunAndReturn(run func(*app.App) error) *MockisNewAccount_Init_Call {
	_c.Call.Return(run)
	return _c
}

// IsNewAccount provides a mock function with given fields:
func (_m *MockisNewAccount) IsNewAccount() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockisNewAccount_IsNewAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsNewAccount'
type MockisNewAccount_IsNewAccount_Call struct {
	*mock.Call
}

// IsNewAccount is a helper method to define mock.On call
func (_e *MockisNewAccount_Expecter) IsNewAccount() *MockisNewAccount_IsNewAccount_Call {
	return &MockisNewAccount_IsNewAccount_Call{Call: _e.mock.On("IsNewAccount")}
}

func (_c *MockisNewAccount_IsNewAccount_Call) Run(run func()) *MockisNewAccount_IsNewAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockisNewAccount_IsNewAccount_Call) Return(_a0 bool) *MockisNewAccount_IsNewAccount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockisNewAccount_IsNewAccount_Call) RunAndReturn(run func() bool) *MockisNewAccount_IsNewAccount_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockisNewAccount) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockisNewAccount_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockisNewAccount_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockisNewAccount_Expecter) Name() *MockisNewAccount_Name_Call {
	return &MockisNewAccount_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockisNewAccount_Name_Call) Run(run func()) *MockisNewAccount_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockisNewAccount_Name_Call) Return(name string) *MockisNewAccount_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockisNewAccount_Name_Call) RunAndReturn(run func() string) *MockisNewAccount_Name_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockisNewAccount interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockisNewAccount creates a new instance of MockisNewAccount. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockisNewAccount(t mockConstructorTestingTNewMockisNewAccount) *MockisNewAccount {
	mock := &MockisNewAccount{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
