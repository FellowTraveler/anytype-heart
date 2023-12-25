// Code generated by mockery v2.38.0. DO NOT EDIT.

package mock_spacecore

import (
	pool "github.com/anyproto/any-sync/net/pool"
	mock "github.com/stretchr/testify/mock"
)

// MockPoolManager is an autogenerated mock type for the PoolManager type
type MockPoolManager struct {
	mock.Mock
}

type MockPoolManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPoolManager) EXPECT() *MockPoolManager_Expecter {
	return &MockPoolManager_Expecter{mock: &_m.Mock}
}

// StreamPeerPool provides a mock function with given fields:
func (_m *MockPoolManager) StreamPeerPool() pool.Pool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StreamPeerPool")
	}

	var r0 pool.Pool
	if rf, ok := ret.Get(0).(func() pool.Pool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pool.Pool)
		}
	}

	return r0
}

// MockPoolManager_StreamPeerPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StreamPeerPool'
type MockPoolManager_StreamPeerPool_Call struct {
	*mock.Call
}

// StreamPeerPool is a helper method to define mock.On call
func (_e *MockPoolManager_Expecter) StreamPeerPool() *MockPoolManager_StreamPeerPool_Call {
	return &MockPoolManager_StreamPeerPool_Call{Call: _e.mock.On("StreamPeerPool")}
}

func (_c *MockPoolManager_StreamPeerPool_Call) Run(run func()) *MockPoolManager_StreamPeerPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPoolManager_StreamPeerPool_Call) Return(_a0 pool.Pool) *MockPoolManager_StreamPeerPool_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPoolManager_StreamPeerPool_Call) RunAndReturn(run func() pool.Pool) *MockPoolManager_StreamPeerPool_Call {
	_c.Call.Return(run)
	return _c
}

// UnaryPeerPool provides a mock function with given fields:
func (_m *MockPoolManager) UnaryPeerPool() pool.Pool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UnaryPeerPool")
	}

	var r0 pool.Pool
	if rf, ok := ret.Get(0).(func() pool.Pool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pool.Pool)
		}
	}

	return r0
}

// MockPoolManager_UnaryPeerPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnaryPeerPool'
type MockPoolManager_UnaryPeerPool_Call struct {
	*mock.Call
}

// UnaryPeerPool is a helper method to define mock.On call
func (_e *MockPoolManager_Expecter) UnaryPeerPool() *MockPoolManager_UnaryPeerPool_Call {
	return &MockPoolManager_UnaryPeerPool_Call{Call: _e.mock.On("UnaryPeerPool")}
}

func (_c *MockPoolManager_UnaryPeerPool_Call) Run(run func()) *MockPoolManager_UnaryPeerPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPoolManager_UnaryPeerPool_Call) Return(_a0 pool.Pool) *MockPoolManager_UnaryPeerPool_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPoolManager_UnaryPeerPool_Call) RunAndReturn(run func() pool.Pool) *MockPoolManager_UnaryPeerPool_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPoolManager creates a new instance of MockPoolManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPoolManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPoolManager {
	mock := &MockPoolManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
