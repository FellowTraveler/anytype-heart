// Code generated by mockery. DO NOT EDIT.

package mock_typeprovider

import (
	app "github.com/anyproto/any-sync/app"
	mock "github.com/stretchr/testify/mock"

	smartblock "github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
)

// MockSmartBlockTypeProvider is an autogenerated mock type for the SmartBlockTypeProvider type
type MockSmartBlockTypeProvider struct {
	mock.Mock
}

type MockSmartBlockTypeProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSmartBlockTypeProvider) EXPECT() *MockSmartBlockTypeProvider_Expecter {
	return &MockSmartBlockTypeProvider_Expecter{mock: &_m.Mock}
}

// Init provides a mock function with given fields: a
func (_m *MockSmartBlockTypeProvider) Init(a *app.App) error {
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

// MockSmartBlockTypeProvider_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockSmartBlockTypeProvider_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockSmartBlockTypeProvider_Expecter) Init(a interface{}) *MockSmartBlockTypeProvider_Init_Call {
	return &MockSmartBlockTypeProvider_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockSmartBlockTypeProvider_Init_Call) Run(run func(a *app.App)) *MockSmartBlockTypeProvider_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockSmartBlockTypeProvider_Init_Call) Return(err error) *MockSmartBlockTypeProvider_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockSmartBlockTypeProvider_Init_Call) RunAndReturn(run func(*app.App) error) *MockSmartBlockTypeProvider_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockSmartBlockTypeProvider) Name() string {
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

// MockSmartBlockTypeProvider_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockSmartBlockTypeProvider_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockSmartBlockTypeProvider_Expecter) Name() *MockSmartBlockTypeProvider_Name_Call {
	return &MockSmartBlockTypeProvider_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockSmartBlockTypeProvider_Name_Call) Run(run func()) *MockSmartBlockTypeProvider_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSmartBlockTypeProvider_Name_Call) Return(name string) *MockSmartBlockTypeProvider_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockSmartBlockTypeProvider_Name_Call) RunAndReturn(run func() string) *MockSmartBlockTypeProvider_Name_Call {
	_c.Call.Return(run)
	return _c
}

// PartitionIDsByType provides a mock function with given fields: spaceId, ids
func (_m *MockSmartBlockTypeProvider) PartitionIDsByType(spaceId string, ids []string) (map[smartblock.SmartBlockType][]string, error) {
	ret := _m.Called(spaceId, ids)

	if len(ret) == 0 {
		panic("no return value specified for PartitionIDsByType")
	}

	var r0 map[smartblock.SmartBlockType][]string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string) (map[smartblock.SmartBlockType][]string, error)); ok {
		return rf(spaceId, ids)
	}
	if rf, ok := ret.Get(0).(func(string, []string) map[smartblock.SmartBlockType][]string); ok {
		r0 = rf(spaceId, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[smartblock.SmartBlockType][]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(spaceId, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSmartBlockTypeProvider_PartitionIDsByType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PartitionIDsByType'
type MockSmartBlockTypeProvider_PartitionIDsByType_Call struct {
	*mock.Call
}

// PartitionIDsByType is a helper method to define mock.On call
//   - spaceId string
//   - ids []string
func (_e *MockSmartBlockTypeProvider_Expecter) PartitionIDsByType(spaceId interface{}, ids interface{}) *MockSmartBlockTypeProvider_PartitionIDsByType_Call {
	return &MockSmartBlockTypeProvider_PartitionIDsByType_Call{Call: _e.mock.On("PartitionIDsByType", spaceId, ids)}
}

func (_c *MockSmartBlockTypeProvider_PartitionIDsByType_Call) Run(run func(spaceId string, ids []string)) *MockSmartBlockTypeProvider_PartitionIDsByType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string))
	})
	return _c
}

func (_c *MockSmartBlockTypeProvider_PartitionIDsByType_Call) Return(_a0 map[smartblock.SmartBlockType][]string, _a1 error) *MockSmartBlockTypeProvider_PartitionIDsByType_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSmartBlockTypeProvider_PartitionIDsByType_Call) RunAndReturn(run func(string, []string) (map[smartblock.SmartBlockType][]string, error)) *MockSmartBlockTypeProvider_PartitionIDsByType_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterStaticType provides a mock function with given fields: id, tp
func (_m *MockSmartBlockTypeProvider) RegisterStaticType(id string, tp smartblock.SmartBlockType) {
	_m.Called(id, tp)
}

// MockSmartBlockTypeProvider_RegisterStaticType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterStaticType'
type MockSmartBlockTypeProvider_RegisterStaticType_Call struct {
	*mock.Call
}

// RegisterStaticType is a helper method to define mock.On call
//   - id string
//   - tp smartblock.SmartBlockType
func (_e *MockSmartBlockTypeProvider_Expecter) RegisterStaticType(id interface{}, tp interface{}) *MockSmartBlockTypeProvider_RegisterStaticType_Call {
	return &MockSmartBlockTypeProvider_RegisterStaticType_Call{Call: _e.mock.On("RegisterStaticType", id, tp)}
}

func (_c *MockSmartBlockTypeProvider_RegisterStaticType_Call) Run(run func(id string, tp smartblock.SmartBlockType)) *MockSmartBlockTypeProvider_RegisterStaticType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(smartblock.SmartBlockType))
	})
	return _c
}

func (_c *MockSmartBlockTypeProvider_RegisterStaticType_Call) Return() *MockSmartBlockTypeProvider_RegisterStaticType_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSmartBlockTypeProvider_RegisterStaticType_Call) RunAndReturn(run func(string, smartblock.SmartBlockType)) *MockSmartBlockTypeProvider_RegisterStaticType_Call {
	_c.Call.Return(run)
	return _c
}

// Type provides a mock function with given fields: spaceID, id
func (_m *MockSmartBlockTypeProvider) Type(spaceID string, id string) (smartblock.SmartBlockType, error) {
	ret := _m.Called(spaceID, id)

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 smartblock.SmartBlockType
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (smartblock.SmartBlockType, error)); ok {
		return rf(spaceID, id)
	}
	if rf, ok := ret.Get(0).(func(string, string) smartblock.SmartBlockType); ok {
		r0 = rf(spaceID, id)
	} else {
		r0 = ret.Get(0).(smartblock.SmartBlockType)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(spaceID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSmartBlockTypeProvider_Type_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Type'
type MockSmartBlockTypeProvider_Type_Call struct {
	*mock.Call
}

// Type is a helper method to define mock.On call
//   - spaceID string
//   - id string
func (_e *MockSmartBlockTypeProvider_Expecter) Type(spaceID interface{}, id interface{}) *MockSmartBlockTypeProvider_Type_Call {
	return &MockSmartBlockTypeProvider_Type_Call{Call: _e.mock.On("Type", spaceID, id)}
}

func (_c *MockSmartBlockTypeProvider_Type_Call) Run(run func(spaceID string, id string)) *MockSmartBlockTypeProvider_Type_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockSmartBlockTypeProvider_Type_Call) Return(_a0 smartblock.SmartBlockType, _a1 error) *MockSmartBlockTypeProvider_Type_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSmartBlockTypeProvider_Type_Call) RunAndReturn(run func(string, string) (smartblock.SmartBlockType, error)) *MockSmartBlockTypeProvider_Type_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSmartBlockTypeProvider creates a new instance of MockSmartBlockTypeProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSmartBlockTypeProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSmartBlockTypeProvider {
	mock := &MockSmartBlockTypeProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
