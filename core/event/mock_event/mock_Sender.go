// Code generated by mockery v2.32.0. DO NOT EDIT.

package mock_event

import (
	app "github.com/anyproto/any-sync/app"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/anyproto/anytype-heart/pb"
)

// MockSender is an autogenerated mock type for the Sender type
type MockSender struct {
	mock.Mock
}

type MockSender_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSender) EXPECT() *MockSender_Expecter {
	return &MockSender_Expecter{mock: &_m.Mock}
}

// Broadcast provides a mock function with given fields: _a0
func (_m *MockSender) Broadcast(_a0 *pb.Event) {
	_m.Called(_a0)
}

// MockSender_Broadcast_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Broadcast'
type MockSender_Broadcast_Call struct {
	*mock.Call
}

// Broadcast is a helper method to define mock.On call
//   - _a0 *pb.Event
func (_e *MockSender_Expecter) Broadcast(_a0 interface{}) *MockSender_Broadcast_Call {
	return &MockSender_Broadcast_Call{Call: _e.mock.On("Broadcast", _a0)}
}

func (_c *MockSender_Broadcast_Call) Run(run func(_a0 *pb.Event)) *MockSender_Broadcast_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*pb.Event))
	})
	return _c
}

func (_c *MockSender_Broadcast_Call) Return() *MockSender_Broadcast_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSender_Broadcast_Call) RunAndReturn(run func(*pb.Event)) *MockSender_Broadcast_Call {
	_c.Call.Return(run)
	return _c
}

// BroadcastToOtherSessions provides a mock function with given fields: token, e
func (_m *MockSender) BroadcastToOtherSessions(token string, e *pb.Event) {
	_m.Called(token, e)
}

// MockSender_BroadcastToOtherSessions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BroadcastToOtherSessions'
type MockSender_BroadcastToOtherSessions_Call struct {
	*mock.Call
}

// BroadcastToOtherSessions is a helper method to define mock.On call
//   - token string
//   - e *pb.Event
func (_e *MockSender_Expecter) BroadcastToOtherSessions(token interface{}, e interface{}) *MockSender_BroadcastToOtherSessions_Call {
	return &MockSender_BroadcastToOtherSessions_Call{Call: _e.mock.On("BroadcastToOtherSessions", token, e)}
}

func (_c *MockSender_BroadcastToOtherSessions_Call) Run(run func(token string, e *pb.Event)) *MockSender_BroadcastToOtherSessions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*pb.Event))
	})
	return _c
}

func (_c *MockSender_BroadcastToOtherSessions_Call) Return() *MockSender_BroadcastToOtherSessions_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSender_BroadcastToOtherSessions_Call) RunAndReturn(run func(string, *pb.Event)) *MockSender_BroadcastToOtherSessions_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockSender) Init(a *app.App) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.App) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSender_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockSender_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockSender_Expecter) Init(a interface{}) *MockSender_Init_Call {
	return &MockSender_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockSender_Init_Call) Run(run func(a *app.App)) *MockSender_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockSender_Init_Call) Return(err error) *MockSender_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockSender_Init_Call) RunAndReturn(run func(*app.App) error) *MockSender_Init_Call {
	_c.Call.Return(run)
	return _c
}

// IsActive provides a mock function with given fields: token
func (_m *MockSender) IsActive(token string) bool {
	ret := _m.Called(token)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockSender_IsActive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsActive'
type MockSender_IsActive_Call struct {
	*mock.Call
}

// IsActive is a helper method to define mock.On call
//   - token string
func (_e *MockSender_Expecter) IsActive(token interface{}) *MockSender_IsActive_Call {
	return &MockSender_IsActive_Call{Call: _e.mock.On("IsActive", token)}
}

func (_c *MockSender_IsActive_Call) Run(run func(token string)) *MockSender_IsActive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockSender_IsActive_Call) Return(_a0 bool) *MockSender_IsActive_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSender_IsActive_Call) RunAndReturn(run func(string) bool) *MockSender_IsActive_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockSender) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockSender_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockSender_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockSender_Expecter) Name() *MockSender_Name_Call {
	return &MockSender_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockSender_Name_Call) Run(run func()) *MockSender_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSender_Name_Call) Return(name string) *MockSender_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockSender_Name_Call) RunAndReturn(run func() string) *MockSender_Name_Call {
	_c.Call.Return(run)
	return _c
}

// SendToSession provides a mock function with given fields: token, _a1
func (_m *MockSender) SendToSession(token string, _a1 *pb.Event) {
	_m.Called(token, _a1)
}

// MockSender_SendToSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendToSession'
type MockSender_SendToSession_Call struct {
	*mock.Call
}

// SendToSession is a helper method to define mock.On call
//   - token string
//   - _a1 *pb.Event
func (_e *MockSender_Expecter) SendToSession(token interface{}, _a1 interface{}) *MockSender_SendToSession_Call {
	return &MockSender_SendToSession_Call{Call: _e.mock.On("SendToSession", token, _a1)}
}

func (_c *MockSender_SendToSession_Call) Run(run func(token string, _a1 *pb.Event)) *MockSender_SendToSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*pb.Event))
	})
	return _c
}

func (_c *MockSender_SendToSession_Call) Return() *MockSender_SendToSession_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSender_SendToSession_Call) RunAndReturn(run func(string, *pb.Event)) *MockSender_SendToSession_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSender creates a new instance of MockSender. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSender(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSender {
	mock := &MockSender{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
