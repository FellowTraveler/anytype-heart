// Code generated by mockery v2.35.2. DO NOT EDIT.

package mock_techspace

import (
	context "context"

	commonspace "github.com/anyproto/any-sync/commonspace"

	mock "github.com/stretchr/testify/mock"

	objectcache "github.com/anyproto/anytype-heart/core/block/object/objectcache"

	spaceinfo "github.com/anyproto/anytype-heart/space/spaceinfo"

	types "github.com/gogo/protobuf/types"
)

// MockTechSpace is an autogenerated mock type for the TechSpace type
type MockTechSpace struct {
	mock.Mock
}

type MockTechSpace_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTechSpace) EXPECT() *MockTechSpace_Expecter {
	return &MockTechSpace_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: ctx
func (_m *MockTechSpace) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTechSpace_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockTechSpace_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockTechSpace_Expecter) Close(ctx interface{}) *MockTechSpace_Close_Call {
	return &MockTechSpace_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *MockTechSpace_Close_Call) Run(run func(ctx context.Context)) *MockTechSpace_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockTechSpace_Close_Call) Return(err error) *MockTechSpace_Close_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockTechSpace_Close_Call) RunAndReturn(run func(context.Context) error) *MockTechSpace_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: techCoreSpace, objectCache
func (_m *MockTechSpace) Run(techCoreSpace commonspace.Space, objectCache objectcache.Cache) error {
	ret := _m.Called(techCoreSpace, objectCache)

	var r0 error
	if rf, ok := ret.Get(0).(func(commonspace.Space, objectcache.Cache) error); ok {
		r0 = rf(techCoreSpace, objectCache)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTechSpace_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockTechSpace_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - techCoreSpace commonspace.Space
//   - objectCache objectcache.Cache
func (_e *MockTechSpace_Expecter) Run(techCoreSpace interface{}, objectCache interface{}) *MockTechSpace_Run_Call {
	return &MockTechSpace_Run_Call{Call: _e.mock.On("Run", techCoreSpace, objectCache)}
}

func (_c *MockTechSpace_Run_Call) Run(run func(techCoreSpace commonspace.Space, objectCache objectcache.Cache)) *MockTechSpace_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(commonspace.Space), args[1].(objectcache.Cache))
	})
	return _c
}

func (_c *MockTechSpace_Run_Call) Return(err error) *MockTechSpace_Run_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockTechSpace_Run_Call) RunAndReturn(run func(commonspace.Space, objectcache.Cache) error) *MockTechSpace_Run_Call {
	_c.Call.Return(run)
	return _c
}

// SetLocalInfo provides a mock function with given fields: ctx, info
func (_m *MockTechSpace) SetLocalInfo(ctx context.Context, info spaceinfo.SpaceLocalInfo) error {
	ret := _m.Called(ctx, info)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, spaceinfo.SpaceLocalInfo) error); ok {
		r0 = rf(ctx, info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTechSpace_SetLocalInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetLocalInfo'
type MockTechSpace_SetLocalInfo_Call struct {
	*mock.Call
}

// SetLocalInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - info spaceinfo.SpaceLocalInfo
func (_e *MockTechSpace_Expecter) SetLocalInfo(ctx interface{}, info interface{}) *MockTechSpace_SetLocalInfo_Call {
	return &MockTechSpace_SetLocalInfo_Call{Call: _e.mock.On("SetLocalInfo", ctx, info)}
}

func (_c *MockTechSpace_SetLocalInfo_Call) Run(run func(ctx context.Context, info spaceinfo.SpaceLocalInfo)) *MockTechSpace_SetLocalInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(spaceinfo.SpaceLocalInfo))
	})
	return _c
}

func (_c *MockTechSpace_SetLocalInfo_Call) Return(err error) *MockTechSpace_SetLocalInfo_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockTechSpace_SetLocalInfo_Call) RunAndReturn(run func(context.Context, spaceinfo.SpaceLocalInfo) error) *MockTechSpace_SetLocalInfo_Call {
	_c.Call.Return(run)
	return _c
}

// SetPersistentInfo provides a mock function with given fields: ctx, info
func (_m *MockTechSpace) SetPersistentInfo(ctx context.Context, info spaceinfo.SpacePersistentInfo) error {
	ret := _m.Called(ctx, info)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, spaceinfo.SpacePersistentInfo) error); ok {
		r0 = rf(ctx, info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTechSpace_SetPersistentInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPersistentInfo'
type MockTechSpace_SetPersistentInfo_Call struct {
	*mock.Call
}

// SetPersistentInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - info spaceinfo.SpacePersistentInfo
func (_e *MockTechSpace_Expecter) SetPersistentInfo(ctx interface{}, info interface{}) *MockTechSpace_SetPersistentInfo_Call {
	return &MockTechSpace_SetPersistentInfo_Call{Call: _e.mock.On("SetPersistentInfo", ctx, info)}
}

func (_c *MockTechSpace_SetPersistentInfo_Call) Run(run func(ctx context.Context, info spaceinfo.SpacePersistentInfo)) *MockTechSpace_SetPersistentInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(spaceinfo.SpacePersistentInfo))
	})
	return _c
}

func (_c *MockTechSpace_SetPersistentInfo_Call) Return(err error) *MockTechSpace_SetPersistentInfo_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockTechSpace_SetPersistentInfo_Call) RunAndReturn(run func(context.Context, spaceinfo.SpacePersistentInfo) error) *MockTechSpace_SetPersistentInfo_Call {
	_c.Call.Return(run)
	return _c
}

// SpaceViewCreate provides a mock function with given fields: ctx, spaceId, force
func (_m *MockTechSpace) SpaceViewCreate(ctx context.Context, spaceId string, force bool) error {
	ret := _m.Called(ctx, spaceId, force)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) error); ok {
		r0 = rf(ctx, spaceId, force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTechSpace_SpaceViewCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SpaceViewCreate'
type MockTechSpace_SpaceViewCreate_Call struct {
	*mock.Call
}

// SpaceViewCreate is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
//   - force bool
func (_e *MockTechSpace_Expecter) SpaceViewCreate(ctx interface{}, spaceId interface{}, force interface{}) *MockTechSpace_SpaceViewCreate_Call {
	return &MockTechSpace_SpaceViewCreate_Call{Call: _e.mock.On("SpaceViewCreate", ctx, spaceId, force)}
}

func (_c *MockTechSpace_SpaceViewCreate_Call) Run(run func(ctx context.Context, spaceId string, force bool)) *MockTechSpace_SpaceViewCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(bool))
	})
	return _c
}

func (_c *MockTechSpace_SpaceViewCreate_Call) Return(err error) *MockTechSpace_SpaceViewCreate_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockTechSpace_SpaceViewCreate_Call) RunAndReturn(run func(context.Context, string, bool) error) *MockTechSpace_SpaceViewCreate_Call {
	_c.Call.Return(run)
	return _c
}

// SpaceViewExists provides a mock function with given fields: ctx, spaceId
func (_m *MockTechSpace) SpaceViewExists(ctx context.Context, spaceId string) (bool, error) {
	ret := _m.Called(ctx, spaceId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, spaceId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, spaceId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, spaceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTechSpace_SpaceViewExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SpaceViewExists'
type MockTechSpace_SpaceViewExists_Call struct {
	*mock.Call
}

// SpaceViewExists is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
func (_e *MockTechSpace_Expecter) SpaceViewExists(ctx interface{}, spaceId interface{}) *MockTechSpace_SpaceViewExists_Call {
	return &MockTechSpace_SpaceViewExists_Call{Call: _e.mock.On("SpaceViewExists", ctx, spaceId)}
}

func (_c *MockTechSpace_SpaceViewExists_Call) Run(run func(ctx context.Context, spaceId string)) *MockTechSpace_SpaceViewExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockTechSpace_SpaceViewExists_Call) Return(exists bool, err error) *MockTechSpace_SpaceViewExists_Call {
	_c.Call.Return(exists, err)
	return _c
}

func (_c *MockTechSpace_SpaceViewExists_Call) RunAndReturn(run func(context.Context, string) (bool, error)) *MockTechSpace_SpaceViewExists_Call {
	_c.Call.Return(run)
	return _c
}

// SpaceViewId provides a mock function with given fields: id
func (_m *MockTechSpace) SpaceViewId(id string) (string, error) {
	ret := _m.Called(id)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTechSpace_SpaceViewId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SpaceViewId'
type MockTechSpace_SpaceViewId_Call struct {
	*mock.Call
}

// SpaceViewId is a helper method to define mock.On call
//   - id string
func (_e *MockTechSpace_Expecter) SpaceViewId(id interface{}) *MockTechSpace_SpaceViewId_Call {
	return &MockTechSpace_SpaceViewId_Call{Call: _e.mock.On("SpaceViewId", id)}
}

func (_c *MockTechSpace_SpaceViewId_Call) Run(run func(id string)) *MockTechSpace_SpaceViewId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTechSpace_SpaceViewId_Call) Return(_a0 string, _a1 error) *MockTechSpace_SpaceViewId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTechSpace_SpaceViewId_Call) RunAndReturn(run func(string) (string, error)) *MockTechSpace_SpaceViewId_Call {
	_c.Call.Return(run)
	return _c
}

// SpaceViewSetData provides a mock function with given fields: ctx, spaceId, details
func (_m *MockTechSpace) SpaceViewSetData(ctx context.Context, spaceId string, details *types.Struct) error {
	ret := _m.Called(ctx, spaceId, details)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *types.Struct) error); ok {
		r0 = rf(ctx, spaceId, details)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTechSpace_SpaceViewSetData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SpaceViewSetData'
type MockTechSpace_SpaceViewSetData_Call struct {
	*mock.Call
}

// SpaceViewSetData is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
//   - details *types.Struct
func (_e *MockTechSpace_Expecter) SpaceViewSetData(ctx interface{}, spaceId interface{}, details interface{}) *MockTechSpace_SpaceViewSetData_Call {
	return &MockTechSpace_SpaceViewSetData_Call{Call: _e.mock.On("SpaceViewSetData", ctx, spaceId, details)}
}

func (_c *MockTechSpace_SpaceViewSetData_Call) Run(run func(ctx context.Context, spaceId string, details *types.Struct)) *MockTechSpace_SpaceViewSetData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*types.Struct))
	})
	return _c
}

func (_c *MockTechSpace_SpaceViewSetData_Call) Return(err error) *MockTechSpace_SpaceViewSetData_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockTechSpace_SpaceViewSetData_Call) RunAndReturn(run func(context.Context, string, *types.Struct) error) *MockTechSpace_SpaceViewSetData_Call {
	_c.Call.Return(run)
	return _c
}

// TechSpaceId provides a mock function with given fields:
func (_m *MockTechSpace) TechSpaceId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockTechSpace_TechSpaceId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TechSpaceId'
type MockTechSpace_TechSpaceId_Call struct {
	*mock.Call
}

// TechSpaceId is a helper method to define mock.On call
func (_e *MockTechSpace_Expecter) TechSpaceId() *MockTechSpace_TechSpaceId_Call {
	return &MockTechSpace_TechSpaceId_Call{Call: _e.mock.On("TechSpaceId")}
}

func (_c *MockTechSpace_TechSpaceId_Call) Run(run func()) *MockTechSpace_TechSpaceId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTechSpace_TechSpaceId_Call) Return(_a0 string) *MockTechSpace_TechSpaceId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTechSpace_TechSpaceId_Call) RunAndReturn(run func() string) *MockTechSpace_TechSpaceId_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTechSpace creates a new instance of MockTechSpace. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTechSpace(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTechSpace {
	mock := &MockTechSpace{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
