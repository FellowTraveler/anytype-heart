// Code generated by mockery v2.35.2. DO NOT EDIT.

package mock_techspace

import (
	context "context"

	app "github.com/anyproto/any-sync/app"

	mock "github.com/stretchr/testify/mock"

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

// Init provides a mock function with given fields: a
func (_m *MockTechSpace) Init(a *app.App) error {
	ret := _m.Called(a)

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

// Run provides a mock function with given fields: ctx
func (_m *MockTechSpace) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
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
//   - ctx context.Context
func (_e *MockTechSpace_Expecter) Run(ctx interface{}) *MockTechSpace_Run_Call {
	return &MockTechSpace_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *MockTechSpace_Run_Call) Run(run func(ctx context.Context)) *MockTechSpace_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockTechSpace_Run_Call) Return(err error) *MockTechSpace_Run_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockTechSpace_Run_Call) RunAndReturn(run func(context.Context) error) *MockTechSpace_Run_Call {
	_c.Call.Return(run)
	return _c
}

// SetInfo provides a mock function with given fields: ctx, info
func (_m *MockTechSpace) SetInfo(ctx context.Context, info spaceinfo.SpaceInfo) error {
	ret := _m.Called(ctx, info)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, spaceinfo.SpaceInfo) error); ok {
		r0 = rf(ctx, info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTechSpace_SetInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetInfo'
type MockTechSpace_SetInfo_Call struct {
	*mock.Call
}

// SetInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - info spaceinfo.SpaceInfo
func (_e *MockTechSpace_Expecter) SetInfo(ctx interface{}, info interface{}) *MockTechSpace_SetInfo_Call {
	return &MockTechSpace_SetInfo_Call{Call: _e.mock.On("SetInfo", ctx, info)}
}

func (_c *MockTechSpace_SetInfo_Call) Run(run func(ctx context.Context, info spaceinfo.SpaceInfo)) *MockTechSpace_SetInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(spaceinfo.SpaceInfo))
	})
	return _c
}

func (_c *MockTechSpace_SetInfo_Call) Return(err error) *MockTechSpace_SetInfo_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockTechSpace_SetInfo_Call) RunAndReturn(run func(context.Context, spaceinfo.SpaceInfo) error) *MockTechSpace_SetInfo_Call {
	_c.Call.Return(run)
	return _c
}

// SpaceViewCreate provides a mock function with given fields: ctx, spaceId
func (_m *MockTechSpace) SpaceViewCreate(ctx context.Context, spaceId string) error {
	ret := _m.Called(ctx, spaceId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, spaceId)
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
func (_e *MockTechSpace_Expecter) SpaceViewCreate(ctx interface{}, spaceId interface{}) *MockTechSpace_SpaceViewCreate_Call {
	return &MockTechSpace_SpaceViewCreate_Call{Call: _e.mock.On("SpaceViewCreate", ctx, spaceId)}
}

func (_c *MockTechSpace_SpaceViewCreate_Call) Run(run func(ctx context.Context, spaceId string)) *MockTechSpace_SpaceViewCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockTechSpace_SpaceViewCreate_Call) Return(err error) *MockTechSpace_SpaceViewCreate_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockTechSpace_SpaceViewCreate_Call) RunAndReturn(run func(context.Context, string) error) *MockTechSpace_SpaceViewCreate_Call {
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
