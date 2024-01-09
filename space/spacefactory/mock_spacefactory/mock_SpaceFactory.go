// Code generated by mockery v2.39.1. DO NOT EDIT.

package mock_spacefactory

import (
	context "context"

	app "github.com/anyproto/any-sync/app"
	clientspace "github.com/anyproto/anytype-heart/space/clientspace"

	mock "github.com/stretchr/testify/mock"

	spacecontroller "github.com/anyproto/anytype-heart/space/internal/spacecontroller"

	spaceinfo "github.com/anyproto/anytype-heart/space/spaceinfo"
)

// MockSpaceFactory is an autogenerated mock type for the SpaceFactory type
type MockSpaceFactory struct {
	mock.Mock
}

type MockSpaceFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSpaceFactory) EXPECT() *MockSpaceFactory_Expecter {
	return &MockSpaceFactory_Expecter{mock: &_m.Mock}
}

// CreateAndSetTechSpace provides a mock function with given fields: ctx
func (_m *MockSpaceFactory) CreateAndSetTechSpace(ctx context.Context) (*clientspace.TechSpace, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CreateAndSetTechSpace")
	}

	var r0 *clientspace.TechSpace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*clientspace.TechSpace, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *clientspace.TechSpace); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientspace.TechSpace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceFactory_CreateAndSetTechSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAndSetTechSpace'
type MockSpaceFactory_CreateAndSetTechSpace_Call struct {
	*mock.Call
}

// CreateAndSetTechSpace is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSpaceFactory_Expecter) CreateAndSetTechSpace(ctx interface{}) *MockSpaceFactory_CreateAndSetTechSpace_Call {
	return &MockSpaceFactory_CreateAndSetTechSpace_Call{Call: _e.mock.On("CreateAndSetTechSpace", ctx)}
}

func (_c *MockSpaceFactory_CreateAndSetTechSpace_Call) Run(run func(ctx context.Context)) *MockSpaceFactory_CreateAndSetTechSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSpaceFactory_CreateAndSetTechSpace_Call) Return(_a0 *clientspace.TechSpace, _a1 error) *MockSpaceFactory_CreateAndSetTechSpace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSpaceFactory_CreateAndSetTechSpace_Call) RunAndReturn(run func(context.Context) (*clientspace.TechSpace, error)) *MockSpaceFactory_CreateAndSetTechSpace_Call {
	_c.Call.Return(run)
	return _c
}

// CreateMarketplaceSpace provides a mock function with given fields: ctx
func (_m *MockSpaceFactory) CreateMarketplaceSpace(ctx context.Context) (spacecontroller.SpaceController, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CreateMarketplaceSpace")
	}

	var r0 spacecontroller.SpaceController
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (spacecontroller.SpaceController, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) spacecontroller.SpaceController); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spacecontroller.SpaceController)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceFactory_CreateMarketplaceSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMarketplaceSpace'
type MockSpaceFactory_CreateMarketplaceSpace_Call struct {
	*mock.Call
}

// CreateMarketplaceSpace is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSpaceFactory_Expecter) CreateMarketplaceSpace(ctx interface{}) *MockSpaceFactory_CreateMarketplaceSpace_Call {
	return &MockSpaceFactory_CreateMarketplaceSpace_Call{Call: _e.mock.On("CreateMarketplaceSpace", ctx)}
}

func (_c *MockSpaceFactory_CreateMarketplaceSpace_Call) Run(run func(ctx context.Context)) *MockSpaceFactory_CreateMarketplaceSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSpaceFactory_CreateMarketplaceSpace_Call) Return(sp spacecontroller.SpaceController, err error) *MockSpaceFactory_CreateMarketplaceSpace_Call {
	_c.Call.Return(sp, err)
	return _c
}

func (_c *MockSpaceFactory_CreateMarketplaceSpace_Call) RunAndReturn(run func(context.Context) (spacecontroller.SpaceController, error)) *MockSpaceFactory_CreateMarketplaceSpace_Call {
	_c.Call.Return(run)
	return _c
}

// CreatePersonalSpace provides a mock function with given fields: ctx
func (_m *MockSpaceFactory) CreatePersonalSpace(ctx context.Context) (spacecontroller.SpaceController, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CreatePersonalSpace")
	}

	var r0 spacecontroller.SpaceController
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (spacecontroller.SpaceController, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) spacecontroller.SpaceController); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spacecontroller.SpaceController)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceFactory_CreatePersonalSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatePersonalSpace'
type MockSpaceFactory_CreatePersonalSpace_Call struct {
	*mock.Call
}

// CreatePersonalSpace is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSpaceFactory_Expecter) CreatePersonalSpace(ctx interface{}) *MockSpaceFactory_CreatePersonalSpace_Call {
	return &MockSpaceFactory_CreatePersonalSpace_Call{Call: _e.mock.On("CreatePersonalSpace", ctx)}
}

func (_c *MockSpaceFactory_CreatePersonalSpace_Call) Run(run func(ctx context.Context)) *MockSpaceFactory_CreatePersonalSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSpaceFactory_CreatePersonalSpace_Call) Return(sp spacecontroller.SpaceController, err error) *MockSpaceFactory_CreatePersonalSpace_Call {
	_c.Call.Return(sp, err)
	return _c
}

func (_c *MockSpaceFactory_CreatePersonalSpace_Call) RunAndReturn(run func(context.Context) (spacecontroller.SpaceController, error)) *MockSpaceFactory_CreatePersonalSpace_Call {
	_c.Call.Return(run)
	return _c
}

// CreateShareableSpace provides a mock function with given fields: ctx, id
func (_m *MockSpaceFactory) CreateShareableSpace(ctx context.Context, id string) (spacecontroller.SpaceController, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for CreateShareableSpace")
	}

	var r0 spacecontroller.SpaceController
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (spacecontroller.SpaceController, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) spacecontroller.SpaceController); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spacecontroller.SpaceController)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceFactory_CreateShareableSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateShareableSpace'
type MockSpaceFactory_CreateShareableSpace_Call struct {
	*mock.Call
}

// CreateShareableSpace is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockSpaceFactory_Expecter) CreateShareableSpace(ctx interface{}, id interface{}) *MockSpaceFactory_CreateShareableSpace_Call {
	return &MockSpaceFactory_CreateShareableSpace_Call{Call: _e.mock.On("CreateShareableSpace", ctx, id)}
}

func (_c *MockSpaceFactory_CreateShareableSpace_Call) Run(run func(ctx context.Context, id string)) *MockSpaceFactory_CreateShareableSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSpaceFactory_CreateShareableSpace_Call) Return(sp spacecontroller.SpaceController, err error) *MockSpaceFactory_CreateShareableSpace_Call {
	_c.Call.Return(sp, err)
	return _c
}

func (_c *MockSpaceFactory_CreateShareableSpace_Call) RunAndReturn(run func(context.Context, string) (spacecontroller.SpaceController, error)) *MockSpaceFactory_CreateShareableSpace_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockSpaceFactory) Init(a *app.App) error {
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

// MockSpaceFactory_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockSpaceFactory_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockSpaceFactory_Expecter) Init(a interface{}) *MockSpaceFactory_Init_Call {
	return &MockSpaceFactory_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockSpaceFactory_Init_Call) Run(run func(a *app.App)) *MockSpaceFactory_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockSpaceFactory_Init_Call) Return(err error) *MockSpaceFactory_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockSpaceFactory_Init_Call) RunAndReturn(run func(*app.App) error) *MockSpaceFactory_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockSpaceFactory) Name() string {
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

// MockSpaceFactory_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockSpaceFactory_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockSpaceFactory_Expecter) Name() *MockSpaceFactory_Name_Call {
	return &MockSpaceFactory_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockSpaceFactory_Name_Call) Run(run func()) *MockSpaceFactory_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSpaceFactory_Name_Call) Return(name string) *MockSpaceFactory_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockSpaceFactory_Name_Call) RunAndReturn(run func() string) *MockSpaceFactory_Name_Call {
	_c.Call.Return(run)
	return _c
}

// NewPersonalSpace provides a mock function with given fields: ctx
func (_m *MockSpaceFactory) NewPersonalSpace(ctx context.Context) (spacecontroller.SpaceController, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for NewPersonalSpace")
	}

	var r0 spacecontroller.SpaceController
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (spacecontroller.SpaceController, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) spacecontroller.SpaceController); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spacecontroller.SpaceController)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceFactory_NewPersonalSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewPersonalSpace'
type MockSpaceFactory_NewPersonalSpace_Call struct {
	*mock.Call
}

// NewPersonalSpace is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSpaceFactory_Expecter) NewPersonalSpace(ctx interface{}) *MockSpaceFactory_NewPersonalSpace_Call {
	return &MockSpaceFactory_NewPersonalSpace_Call{Call: _e.mock.On("NewPersonalSpace", ctx)}
}

func (_c *MockSpaceFactory_NewPersonalSpace_Call) Run(run func(ctx context.Context)) *MockSpaceFactory_NewPersonalSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSpaceFactory_NewPersonalSpace_Call) Return(_a0 spacecontroller.SpaceController, _a1 error) *MockSpaceFactory_NewPersonalSpace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSpaceFactory_NewPersonalSpace_Call) RunAndReturn(run func(context.Context) (spacecontroller.SpaceController, error)) *MockSpaceFactory_NewPersonalSpace_Call {
	_c.Call.Return(run)
	return _c
}

// NewShareableSpace provides a mock function with given fields: ctx, id, status
func (_m *MockSpaceFactory) NewShareableSpace(ctx context.Context, id string, status spaceinfo.AccountStatus) (spacecontroller.SpaceController, error) {
	ret := _m.Called(ctx, id, status)

	if len(ret) == 0 {
		panic("no return value specified for NewShareableSpace")
	}

	var r0 spacecontroller.SpaceController
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, spaceinfo.AccountStatus) (spacecontroller.SpaceController, error)); ok {
		return rf(ctx, id, status)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, spaceinfo.AccountStatus) spacecontroller.SpaceController); ok {
		r0 = rf(ctx, id, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spacecontroller.SpaceController)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, spaceinfo.AccountStatus) error); ok {
		r1 = rf(ctx, id, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceFactory_NewShareableSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewShareableSpace'
type MockSpaceFactory_NewShareableSpace_Call struct {
	*mock.Call
}

// NewShareableSpace is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - status spaceinfo.AccountStatus
func (_e *MockSpaceFactory_Expecter) NewShareableSpace(ctx interface{}, id interface{}, status interface{}) *MockSpaceFactory_NewShareableSpace_Call {
	return &MockSpaceFactory_NewShareableSpace_Call{Call: _e.mock.On("NewShareableSpace", ctx, id, status)}
}

func (_c *MockSpaceFactory_NewShareableSpace_Call) Run(run func(ctx context.Context, id string, status spaceinfo.AccountStatus)) *MockSpaceFactory_NewShareableSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(spaceinfo.AccountStatus))
	})
	return _c
}

func (_c *MockSpaceFactory_NewShareableSpace_Call) Return(_a0 spacecontroller.SpaceController, _a1 error) *MockSpaceFactory_NewShareableSpace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSpaceFactory_NewShareableSpace_Call) RunAndReturn(run func(context.Context, string, spaceinfo.AccountStatus) (spacecontroller.SpaceController, error)) *MockSpaceFactory_NewShareableSpace_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSpaceFactory creates a new instance of MockSpaceFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSpaceFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSpaceFactory {
	mock := &MockSpaceFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
