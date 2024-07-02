// Code generated by mockery. DO NOT EDIT.

package objectsyncstatus

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockUpdateReceiver is an autogenerated mock type for the UpdateReceiver type
type MockUpdateReceiver struct {
	mock.Mock
}

type MockUpdateReceiver_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUpdateReceiver) EXPECT() *MockUpdateReceiver_Expecter {
	return &MockUpdateReceiver_Expecter{mock: &_m.Mock}
}

// UpdateNodeStatus provides a mock function with given fields:
func (_m *MockUpdateReceiver) UpdateNodeStatus() {
	_m.Called()
}

// MockUpdateReceiver_UpdateNodeStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateNodeStatus'
type MockUpdateReceiver_UpdateNodeStatus_Call struct {
	*mock.Call
}

// UpdateNodeStatus is a helper method to define mock.On call
func (_e *MockUpdateReceiver_Expecter) UpdateNodeStatus() *MockUpdateReceiver_UpdateNodeStatus_Call {
	return &MockUpdateReceiver_UpdateNodeStatus_Call{Call: _e.mock.On("UpdateNodeStatus")}
}

func (_c *MockUpdateReceiver_UpdateNodeStatus_Call) Run(run func()) *MockUpdateReceiver_UpdateNodeStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUpdateReceiver_UpdateNodeStatus_Call) Return() *MockUpdateReceiver_UpdateNodeStatus_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockUpdateReceiver_UpdateNodeStatus_Call) RunAndReturn(run func()) *MockUpdateReceiver_UpdateNodeStatus_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTree provides a mock function with given fields: ctx, treeId, status
func (_m *MockUpdateReceiver) UpdateTree(ctx context.Context, treeId string, status SyncStatus) error {
	ret := _m.Called(ctx, treeId, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTree")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, SyncStatus) error); ok {
		r0 = rf(ctx, treeId, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUpdateReceiver_UpdateTree_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTree'
type MockUpdateReceiver_UpdateTree_Call struct {
	*mock.Call
}

// UpdateTree is a helper method to define mock.On call
//   - ctx context.Context
//   - treeId string
//   - status SyncStatus
func (_e *MockUpdateReceiver_Expecter) UpdateTree(ctx interface{}, treeId interface{}, status interface{}) *MockUpdateReceiver_UpdateTree_Call {
	return &MockUpdateReceiver_UpdateTree_Call{Call: _e.mock.On("UpdateTree", ctx, treeId, status)}
}

func (_c *MockUpdateReceiver_UpdateTree_Call) Run(run func(ctx context.Context, treeId string, status SyncStatus)) *MockUpdateReceiver_UpdateTree_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(SyncStatus))
	})
	return _c
}

func (_c *MockUpdateReceiver_UpdateTree_Call) Return(err error) *MockUpdateReceiver_UpdateTree_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockUpdateReceiver_UpdateTree_Call) RunAndReturn(run func(context.Context, string, SyncStatus) error) *MockUpdateReceiver_UpdateTree_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUpdateReceiver creates a new instance of MockUpdateReceiver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUpdateReceiver(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUpdateReceiver {
	mock := &MockUpdateReceiver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
