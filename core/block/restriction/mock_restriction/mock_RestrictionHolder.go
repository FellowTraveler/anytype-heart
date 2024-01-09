// Code generated by mockery v2.39.1. DO NOT EDIT.

package mock_restriction

import (
	domain "github.com/anyproto/anytype-heart/core/domain"
	mock "github.com/stretchr/testify/mock"

	model "github.com/anyproto/anytype-heart/pkg/lib/pb/model"

	smartblock "github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
)

// MockRestrictionHolder is an autogenerated mock type for the RestrictionHolder type
type MockRestrictionHolder struct {
	mock.Mock
}

type MockRestrictionHolder_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRestrictionHolder) EXPECT() *MockRestrictionHolder_Expecter {
	return &MockRestrictionHolder_Expecter{mock: &_m.Mock}
}

// Layout provides a mock function with given fields:
func (_m *MockRestrictionHolder) Layout() (model.ObjectTypeLayout, bool) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Layout")
	}

	var r0 model.ObjectTypeLayout
	var r1 bool
	if rf, ok := ret.Get(0).(func() (model.ObjectTypeLayout, bool)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() model.ObjectTypeLayout); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.ObjectTypeLayout)
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockRestrictionHolder_Layout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Layout'
type MockRestrictionHolder_Layout_Call struct {
	*mock.Call
}

// Layout is a helper method to define mock.On call
func (_e *MockRestrictionHolder_Expecter) Layout() *MockRestrictionHolder_Layout_Call {
	return &MockRestrictionHolder_Layout_Call{Call: _e.mock.On("Layout")}
}

func (_c *MockRestrictionHolder_Layout_Call) Run(run func()) *MockRestrictionHolder_Layout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRestrictionHolder_Layout_Call) Return(_a0 model.ObjectTypeLayout, _a1 bool) *MockRestrictionHolder_Layout_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRestrictionHolder_Layout_Call) RunAndReturn(run func() (model.ObjectTypeLayout, bool)) *MockRestrictionHolder_Layout_Call {
	_c.Call.Return(run)
	return _c
}

// Type provides a mock function with given fields:
func (_m *MockRestrictionHolder) Type() smartblock.SmartBlockType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 smartblock.SmartBlockType
	if rf, ok := ret.Get(0).(func() smartblock.SmartBlockType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(smartblock.SmartBlockType)
	}

	return r0
}

// MockRestrictionHolder_Type_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Type'
type MockRestrictionHolder_Type_Call struct {
	*mock.Call
}

// Type is a helper method to define mock.On call
func (_e *MockRestrictionHolder_Expecter) Type() *MockRestrictionHolder_Type_Call {
	return &MockRestrictionHolder_Type_Call{Call: _e.mock.On("Type")}
}

func (_c *MockRestrictionHolder_Type_Call) Run(run func()) *MockRestrictionHolder_Type_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRestrictionHolder_Type_Call) Return(_a0 smartblock.SmartBlockType) *MockRestrictionHolder_Type_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRestrictionHolder_Type_Call) RunAndReturn(run func() smartblock.SmartBlockType) *MockRestrictionHolder_Type_Call {
	_c.Call.Return(run)
	return _c
}

// UniqueKey provides a mock function with given fields:
func (_m *MockRestrictionHolder) UniqueKey() domain.UniqueKey {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UniqueKey")
	}

	var r0 domain.UniqueKey
	if rf, ok := ret.Get(0).(func() domain.UniqueKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.UniqueKey)
		}
	}

	return r0
}

// MockRestrictionHolder_UniqueKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UniqueKey'
type MockRestrictionHolder_UniqueKey_Call struct {
	*mock.Call
}

// UniqueKey is a helper method to define mock.On call
func (_e *MockRestrictionHolder_Expecter) UniqueKey() *MockRestrictionHolder_UniqueKey_Call {
	return &MockRestrictionHolder_UniqueKey_Call{Call: _e.mock.On("UniqueKey")}
}

func (_c *MockRestrictionHolder_UniqueKey_Call) Run(run func()) *MockRestrictionHolder_UniqueKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRestrictionHolder_UniqueKey_Call) Return(_a0 domain.UniqueKey) *MockRestrictionHolder_UniqueKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRestrictionHolder_UniqueKey_Call) RunAndReturn(run func() domain.UniqueKey) *MockRestrictionHolder_UniqueKey_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRestrictionHolder creates a new instance of MockRestrictionHolder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRestrictionHolder(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRestrictionHolder {
	mock := &MockRestrictionHolder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
