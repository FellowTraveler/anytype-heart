// Code generated by mockery v2.30.1. DO NOT EDIT.

package mock_objectid

import (
	"context"
	time "time"

	treestorage "github.com/anyproto/any-sync/commonspace/object/tree/treestorage"
	mock "github.com/stretchr/testify/mock"

	"github.com/anyproto/anytype-heart/core/block/import/common"
	"github.com/anyproto/anytype-heart/core/domain/objectorigin"
)

// MockIDGetter is an autogenerated mock type for the IDGetter type
type MockIDGetter struct {
	mock.Mock
}

type MockIDGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIDGetter) EXPECT() *MockIDGetter_Expecter {
	return &MockIDGetter_Expecter{mock: &_m.Mock}
}

// GetID provides a mock function with given fields: spaceID, sn, createdTime, getExisting
func (_m *MockIDGetter) GetIDAndPayload(ctx context.Context, spaceID string, sn *common.Snapshot, createdTime time.Time, getExisting bool, origin objectorigin.ObjectOrigin) (string, treestorage.TreeStorageCreatePayload, error) {
	ret := _m.Called(spaceID, sn, createdTime, getExisting)

	var r0 string
	var r1 treestorage.TreeStorageCreatePayload
	var r2 error
	if rf, ok := ret.Get(0).(func(string, *common.Snapshot, time.Time, bool) (string, treestorage.TreeStorageCreatePayload, error)); ok {
		return rf(spaceID, sn, createdTime, getExisting)
	}
	if rf, ok := ret.Get(0).(func(string, *common.Snapshot, time.Time, bool) string); ok {
		r0 = rf(spaceID, sn, createdTime, getExisting)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, *common.Snapshot, time.Time, bool) treestorage.TreeStorageCreatePayload); ok {
		r1 = rf(spaceID, sn, createdTime, getExisting)
	} else {
		r1 = ret.Get(1).(treestorage.TreeStorageCreatePayload)
	}

	if rf, ok := ret.Get(2).(func(string, *common.Snapshot, time.Time, bool) error); ok {
		r2 = rf(spaceID, sn, createdTime, getExisting)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockIDGetter_GetID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIDAndPayload'
type MockIDGetter_GetID_Call struct {
	*mock.Call
}

// GetID is a helper method to define mock.On call
//   - spaceID string
//   - sn *common.Snapshot
//   - createdTime time.Time
//   - getExisting bool
func (_e *MockIDGetter_Expecter) GetID(spaceID interface{}, sn interface{}, createdTime interface{}, getExisting interface{}) *MockIDGetter_GetID_Call {
	return &MockIDGetter_GetID_Call{Call: _e.mock.On("GetIDAndPayload", spaceID, sn, createdTime, getExisting)}
}

func (_c *MockIDGetter_GetID_Call) Run(run func(spaceID string, sn *common.Snapshot, createdTime time.Time, getExisting bool)) *MockIDGetter_GetID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*common.Snapshot), args[2].(time.Time), args[3].(bool))
	})
	return _c
}

func (_c *MockIDGetter_GetID_Call) Return(_a0 string, _a1 treestorage.TreeStorageCreatePayload, _a2 error) *MockIDGetter_GetID_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockIDGetter_GetID_Call) RunAndReturn(run func(string, *common.Snapshot, time.Time, bool) (string, treestorage.TreeStorageCreatePayload, error)) *MockIDGetter_GetID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIDGetter creates a new instance of MockIDGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIDGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIDGetter {
	mock := &MockIDGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
