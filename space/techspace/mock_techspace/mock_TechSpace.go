// Code generated by mockery. DO NOT EDIT.

package mock_techspace

import (
	app "github.com/anyproto/any-sync/app"
	commonspace "github.com/anyproto/any-sync/commonspace"

	context "context"

	mock "github.com/stretchr/testify/mock"

	objectcache "github.com/anyproto/anytype-heart/core/block/object/objectcache"

	spaceinfo "github.com/anyproto/anytype-heart/space/spaceinfo"

	techspace "github.com/anyproto/anytype-heart/space/techspace"

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

// DoSpaceView provides a mock function with given fields: ctx, spaceID, apply
func (_m *MockTechSpace) DoSpaceView(ctx context.Context, spaceID string, apply func(techspace.SpaceView) error) error {
	ret := _m.Called(ctx, spaceID, apply)

	if len(ret) == 0 {
		panic("no return value specified for DoSpaceView")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, func(techspace.SpaceView) error) error); ok {
		r0 = rf(ctx, spaceID, apply)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTechSpace_DoSpaceView_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DoSpaceView'
type MockTechSpace_DoSpaceView_Call struct {
	*mock.Call
}

// DoSpaceView is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceID string
//   - apply func(techspace.SpaceView) error
func (_e *MockTechSpace_Expecter) DoSpaceView(ctx interface{}, spaceID interface{}, apply interface{}) *MockTechSpace_DoSpaceView_Call {
	return &MockTechSpace_DoSpaceView_Call{Call: _e.mock.On("DoSpaceView", ctx, spaceID, apply)}
}

func (_c *MockTechSpace_DoSpaceView_Call) Run(run func(ctx context.Context, spaceID string, apply func(techspace.SpaceView) error)) *MockTechSpace_DoSpaceView_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(func(techspace.SpaceView) error))
	})
	return _c
}

func (_c *MockTechSpace_DoSpaceView_Call) Return(err error) *MockTechSpace_DoSpaceView_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockTechSpace_DoSpaceView_Call) RunAndReturn(run func(context.Context, string, func(techspace.SpaceView) error) error) *MockTechSpace_DoSpaceView_Call {
	_c.Call.Return(run)
	return _c
}

// GetSpaceView provides a mock function with given fields: ctx, spaceId
func (_m *MockTechSpace) GetSpaceView(ctx context.Context, spaceId string) (techspace.SpaceView, error) {
	ret := _m.Called(ctx, spaceId)

	if len(ret) == 0 {
		panic("no return value specified for GetSpaceView")
	}

	var r0 techspace.SpaceView
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (techspace.SpaceView, error)); ok {
		return rf(ctx, spaceId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) techspace.SpaceView); ok {
		r0 = rf(ctx, spaceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(techspace.SpaceView)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, spaceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTechSpace_GetSpaceView_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSpaceView'
type MockTechSpace_GetSpaceView_Call struct {
	*mock.Call
}

// GetSpaceView is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
func (_e *MockTechSpace_Expecter) GetSpaceView(ctx interface{}, spaceId interface{}) *MockTechSpace_GetSpaceView_Call {
	return &MockTechSpace_GetSpaceView_Call{Call: _e.mock.On("GetSpaceView", ctx, spaceId)}
}

func (_c *MockTechSpace_GetSpaceView_Call) Run(run func(ctx context.Context, spaceId string)) *MockTechSpace_GetSpaceView_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockTechSpace_GetSpaceView_Call) Return(_a0 techspace.SpaceView, _a1 error) *MockTechSpace_GetSpaceView_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTechSpace_GetSpaceView_Call) RunAndReturn(run func(context.Context, string) (techspace.SpaceView, error)) *MockTechSpace_GetSpaceView_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockTechSpace) Init(a *app.App) error {
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

// MockTechSpace_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockTechSpace_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockTechSpace_Expecter) Init(a interface{}) *MockTechSpace_Init_Call {
	return &MockTechSpace_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockTechSpace_Init_Call) Run(run func(a *app.App)) *MockTechSpace_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockTechSpace_Init_Call) Return(err error) *MockTechSpace_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockTechSpace_Init_Call) RunAndReturn(run func(*app.App) error) *MockTechSpace_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockTechSpace) Name() string {
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

// MockTechSpace_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockTechSpace_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockTechSpace_Expecter) Name() *MockTechSpace_Name_Call {
	return &MockTechSpace_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockTechSpace_Name_Call) Run(run func()) *MockTechSpace_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTechSpace_Name_Call) Return(name string) *MockTechSpace_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockTechSpace_Name_Call) RunAndReturn(run func() string) *MockTechSpace_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: techCoreSpace, objectCache
func (_m *MockTechSpace) Run(techCoreSpace commonspace.Space, objectCache objectcache.Cache) error {
	ret := _m.Called(techCoreSpace, objectCache)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

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

	if len(ret) == 0 {
		panic("no return value specified for SetLocalInfo")
	}

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

	if len(ret) == 0 {
		panic("no return value specified for SetPersistentInfo")
	}

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

// SpaceViewCreate provides a mock function with given fields: ctx, spaceId, force, info
func (_m *MockTechSpace) SpaceViewCreate(ctx context.Context, spaceId string, force bool, info spaceinfo.SpacePersistentInfo) error {
	ret := _m.Called(ctx, spaceId, force, info)

	if len(ret) == 0 {
		panic("no return value specified for SpaceViewCreate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool, spaceinfo.SpacePersistentInfo) error); ok {
		r0 = rf(ctx, spaceId, force, info)
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
//   - info spaceinfo.SpacePersistentInfo
func (_e *MockTechSpace_Expecter) SpaceViewCreate(ctx interface{}, spaceId interface{}, force interface{}, info interface{}) *MockTechSpace_SpaceViewCreate_Call {
	return &MockTechSpace_SpaceViewCreate_Call{Call: _e.mock.On("SpaceViewCreate", ctx, spaceId, force, info)}
}

func (_c *MockTechSpace_SpaceViewCreate_Call) Run(run func(ctx context.Context, spaceId string, force bool, info spaceinfo.SpacePersistentInfo)) *MockTechSpace_SpaceViewCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(bool), args[3].(spaceinfo.SpacePersistentInfo))
	})
	return _c
}

func (_c *MockTechSpace_SpaceViewCreate_Call) Return(err error) *MockTechSpace_SpaceViewCreate_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockTechSpace_SpaceViewCreate_Call) RunAndReturn(run func(context.Context, string, bool, spaceinfo.SpacePersistentInfo) error) *MockTechSpace_SpaceViewCreate_Call {
	_c.Call.Return(run)
	return _c
}

// SpaceViewExists provides a mock function with given fields: ctx, spaceId
func (_m *MockTechSpace) SpaceViewExists(ctx context.Context, spaceId string) (bool, error) {
	ret := _m.Called(ctx, spaceId)

	if len(ret) == 0 {
		panic("no return value specified for SpaceViewExists")
	}

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

	if len(ret) == 0 {
		panic("no return value specified for SpaceViewId")
	}

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

	if len(ret) == 0 {
		panic("no return value specified for SpaceViewSetData")
	}

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

	if len(ret) == 0 {
		panic("no return value specified for TechSpaceId")
	}

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

// WakeUpViews provides a mock function with given fields:
func (_m *MockTechSpace) WakeUpViews() {
	_m.Called()
}

// MockTechSpace_WakeUpViews_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WakeUpViews'
type MockTechSpace_WakeUpViews_Call struct {
	*mock.Call
}

// WakeUpViews is a helper method to define mock.On call
func (_e *MockTechSpace_Expecter) WakeUpViews() *MockTechSpace_WakeUpViews_Call {
	return &MockTechSpace_WakeUpViews_Call{Call: _e.mock.On("WakeUpViews")}
}

func (_c *MockTechSpace_WakeUpViews_Call) Run(run func()) *MockTechSpace_WakeUpViews_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTechSpace_WakeUpViews_Call) Return() *MockTechSpace_WakeUpViews_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTechSpace_WakeUpViews_Call) RunAndReturn(run func()) *MockTechSpace_WakeUpViews_Call {
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
