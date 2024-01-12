// Code generated by mockery. DO NOT EDIT.

package mock_database

import (
	database "github.com/anyproto/anytype-heart/pkg/lib/database"
	mock "github.com/stretchr/testify/mock"
)

// MockOrder is an autogenerated mock type for the Order type
type MockOrder struct {
	mock.Mock
}

type MockOrder_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOrder) EXPECT() *MockOrder_Expecter {
	return &MockOrder_Expecter{mock: &_m.Mock}
}

// Compare provides a mock function with given fields: a, b
func (_m *MockOrder) Compare(a database.Getter, b database.Getter) int {
	ret := _m.Called(a, b)

	if len(ret) == 0 {
		panic("no return value specified for Compare")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func(database.Getter, database.Getter) int); ok {
		r0 = rf(a, b)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockOrder_Compare_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Compare'
type MockOrder_Compare_Call struct {
	*mock.Call
}

// Compare is a helper method to define mock.On call
//   - a database.Getter
//   - b database.Getter
func (_e *MockOrder_Expecter) Compare(a interface{}, b interface{}) *MockOrder_Compare_Call {
	return &MockOrder_Compare_Call{Call: _e.mock.On("Compare", a, b)}
}

func (_c *MockOrder_Compare_Call) Run(run func(a database.Getter, b database.Getter)) *MockOrder_Compare_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(database.Getter), args[1].(database.Getter))
	})
	return _c
}

func (_c *MockOrder_Compare_Call) Return(_a0 int) *MockOrder_Compare_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOrder_Compare_Call) RunAndReturn(run func(database.Getter, database.Getter) int) *MockOrder_Compare_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockOrder) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockOrder_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockOrder_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockOrder_Expecter) String() *MockOrder_String_Call {
	return &MockOrder_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockOrder_String_Call) Run(run func()) *MockOrder_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockOrder_String_Call) Return(_a0 string) *MockOrder_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOrder_String_Call) RunAndReturn(run func() string) *MockOrder_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOrder creates a new instance of MockOrder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOrder(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOrder {
	mock := &MockOrder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
