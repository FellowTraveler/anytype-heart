// Code generated by mockery v2.39.1. DO NOT EDIT.

package mock_objectcache

import (
	context "context"

	domain "github.com/anyproto/anytype-heart/core/domain"
	mock "github.com/stretchr/testify/mock"

	objectcache "github.com/anyproto/anytype-heart/core/block/object/objectcache"

	payloadcreator "github.com/anyproto/anytype-heart/core/block/object/payloadcreator"

	smartblock "github.com/anyproto/anytype-heart/core/block/editor/smartblock"

	treestorage "github.com/anyproto/any-sync/commonspace/object/tree/treestorage"
)

// MockCache is an autogenerated mock type for the Cache type
type MockCache struct {
	mock.Mock
}

type MockCache_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCache) EXPECT() *MockCache_Expecter {
	return &MockCache_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: ctx
func (_m *MockCache) Close(ctx context.Context) error {
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

// MockCache_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockCache_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockCache_Expecter) Close(ctx interface{}) *MockCache_Close_Call {
	return &MockCache_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *MockCache_Close_Call) Run(run func(ctx context.Context)) *MockCache_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockCache_Close_Call) Return(_a0 error) *MockCache_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_Close_Call) RunAndReturn(run func(context.Context) error) *MockCache_Close_Call {
	_c.Call.Return(run)
	return _c
}

// CloseBlocks provides a mock function with given fields:
func (_m *MockCache) CloseBlocks() {
	_m.Called()
}

// MockCache_CloseBlocks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseBlocks'
type MockCache_CloseBlocks_Call struct {
	*mock.Call
}

// CloseBlocks is a helper method to define mock.On call
func (_e *MockCache_Expecter) CloseBlocks() *MockCache_CloseBlocks_Call {
	return &MockCache_CloseBlocks_Call{Call: _e.mock.On("CloseBlocks")}
}

func (_c *MockCache_CloseBlocks_Call) Run(run func()) *MockCache_CloseBlocks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCache_CloseBlocks_Call) Return() *MockCache_CloseBlocks_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_CloseBlocks_Call) RunAndReturn(run func()) *MockCache_CloseBlocks_Call {
	_c.Call.Return(run)
	return _c
}

// CreateTreeObject provides a mock function with given fields: ctx, params
func (_m *MockCache) CreateTreeObject(ctx context.Context, params objectcache.TreeCreationParams) (smartblock.SmartBlock, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for CreateTreeObject")
	}

	var r0 smartblock.SmartBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, objectcache.TreeCreationParams) (smartblock.SmartBlock, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, objectcache.TreeCreationParams) smartblock.SmartBlock); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(smartblock.SmartBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, objectcache.TreeCreationParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_CreateTreeObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTreeObject'
type MockCache_CreateTreeObject_Call struct {
	*mock.Call
}

// CreateTreeObject is a helper method to define mock.On call
//   - ctx context.Context
//   - params objectcache.TreeCreationParams
func (_e *MockCache_Expecter) CreateTreeObject(ctx interface{}, params interface{}) *MockCache_CreateTreeObject_Call {
	return &MockCache_CreateTreeObject_Call{Call: _e.mock.On("CreateTreeObject", ctx, params)}
}

func (_c *MockCache_CreateTreeObject_Call) Run(run func(ctx context.Context, params objectcache.TreeCreationParams)) *MockCache_CreateTreeObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(objectcache.TreeCreationParams))
	})
	return _c
}

func (_c *MockCache_CreateTreeObject_Call) Return(sb smartblock.SmartBlock, err error) *MockCache_CreateTreeObject_Call {
	_c.Call.Return(sb, err)
	return _c
}

func (_c *MockCache_CreateTreeObject_Call) RunAndReturn(run func(context.Context, objectcache.TreeCreationParams) (smartblock.SmartBlock, error)) *MockCache_CreateTreeObject_Call {
	_c.Call.Return(run)
	return _c
}

// CreateTreeObjectWithPayload provides a mock function with given fields: ctx, payload, initFunc
func (_m *MockCache) CreateTreeObjectWithPayload(ctx context.Context, payload treestorage.TreeStorageCreatePayload, initFunc func(string) *smartblock.InitContext) (smartblock.SmartBlock, error) {
	ret := _m.Called(ctx, payload, initFunc)

	if len(ret) == 0 {
		panic("no return value specified for CreateTreeObjectWithPayload")
	}

	var r0 smartblock.SmartBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, treestorage.TreeStorageCreatePayload, func(string) *smartblock.InitContext) (smartblock.SmartBlock, error)); ok {
		return rf(ctx, payload, initFunc)
	}
	if rf, ok := ret.Get(0).(func(context.Context, treestorage.TreeStorageCreatePayload, func(string) *smartblock.InitContext) smartblock.SmartBlock); ok {
		r0 = rf(ctx, payload, initFunc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(smartblock.SmartBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, treestorage.TreeStorageCreatePayload, func(string) *smartblock.InitContext) error); ok {
		r1 = rf(ctx, payload, initFunc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_CreateTreeObjectWithPayload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTreeObjectWithPayload'
type MockCache_CreateTreeObjectWithPayload_Call struct {
	*mock.Call
}

// CreateTreeObjectWithPayload is a helper method to define mock.On call
//   - ctx context.Context
//   - payload treestorage.TreeStorageCreatePayload
//   - initFunc func(string) *smartblock.InitContext
func (_e *MockCache_Expecter) CreateTreeObjectWithPayload(ctx interface{}, payload interface{}, initFunc interface{}) *MockCache_CreateTreeObjectWithPayload_Call {
	return &MockCache_CreateTreeObjectWithPayload_Call{Call: _e.mock.On("CreateTreeObjectWithPayload", ctx, payload, initFunc)}
}

func (_c *MockCache_CreateTreeObjectWithPayload_Call) Run(run func(ctx context.Context, payload treestorage.TreeStorageCreatePayload, initFunc func(string) *smartblock.InitContext)) *MockCache_CreateTreeObjectWithPayload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(treestorage.TreeStorageCreatePayload), args[2].(func(string) *smartblock.InitContext))
	})
	return _c
}

func (_c *MockCache_CreateTreeObjectWithPayload_Call) Return(sb smartblock.SmartBlock, err error) *MockCache_CreateTreeObjectWithPayload_Call {
	_c.Call.Return(sb, err)
	return _c
}

func (_c *MockCache_CreateTreeObjectWithPayload_Call) RunAndReturn(run func(context.Context, treestorage.TreeStorageCreatePayload, func(string) *smartblock.InitContext) (smartblock.SmartBlock, error)) *MockCache_CreateTreeObjectWithPayload_Call {
	_c.Call.Return(run)
	return _c
}

// CreateTreePayload provides a mock function with given fields: ctx, params
func (_m *MockCache) CreateTreePayload(ctx context.Context, params payloadcreator.PayloadCreationParams) (treestorage.TreeStorageCreatePayload, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for CreateTreePayload")
	}

	var r0 treestorage.TreeStorageCreatePayload
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, payloadcreator.PayloadCreationParams) (treestorage.TreeStorageCreatePayload, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, payloadcreator.PayloadCreationParams) treestorage.TreeStorageCreatePayload); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(treestorage.TreeStorageCreatePayload)
	}

	if rf, ok := ret.Get(1).(func(context.Context, payloadcreator.PayloadCreationParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_CreateTreePayload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTreePayload'
type MockCache_CreateTreePayload_Call struct {
	*mock.Call
}

// CreateTreePayload is a helper method to define mock.On call
//   - ctx context.Context
//   - params payloadcreator.PayloadCreationParams
func (_e *MockCache_Expecter) CreateTreePayload(ctx interface{}, params interface{}) *MockCache_CreateTreePayload_Call {
	return &MockCache_CreateTreePayload_Call{Call: _e.mock.On("CreateTreePayload", ctx, params)}
}

func (_c *MockCache_CreateTreePayload_Call) Run(run func(ctx context.Context, params payloadcreator.PayloadCreationParams)) *MockCache_CreateTreePayload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(payloadcreator.PayloadCreationParams))
	})
	return _c
}

func (_c *MockCache_CreateTreePayload_Call) Return(_a0 treestorage.TreeStorageCreatePayload, _a1 error) *MockCache_CreateTreePayload_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_CreateTreePayload_Call) RunAndReturn(run func(context.Context, payloadcreator.PayloadCreationParams) (treestorage.TreeStorageCreatePayload, error)) *MockCache_CreateTreePayload_Call {
	_c.Call.Return(run)
	return _c
}

// DeriveObjectID provides a mock function with given fields: ctx, uniqueKey
func (_m *MockCache) DeriveObjectID(ctx context.Context, uniqueKey domain.UniqueKey) (string, error) {
	ret := _m.Called(ctx, uniqueKey)

	if len(ret) == 0 {
		panic("no return value specified for DeriveObjectID")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.UniqueKey) (string, error)); ok {
		return rf(ctx, uniqueKey)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.UniqueKey) string); ok {
		r0 = rf(ctx, uniqueKey)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.UniqueKey) error); ok {
		r1 = rf(ctx, uniqueKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_DeriveObjectID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeriveObjectID'
type MockCache_DeriveObjectID_Call struct {
	*mock.Call
}

// DeriveObjectID is a helper method to define mock.On call
//   - ctx context.Context
//   - uniqueKey domain.UniqueKey
func (_e *MockCache_Expecter) DeriveObjectID(ctx interface{}, uniqueKey interface{}) *MockCache_DeriveObjectID_Call {
	return &MockCache_DeriveObjectID_Call{Call: _e.mock.On("DeriveObjectID", ctx, uniqueKey)}
}

func (_c *MockCache_DeriveObjectID_Call) Run(run func(ctx context.Context, uniqueKey domain.UniqueKey)) *MockCache_DeriveObjectID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.UniqueKey))
	})
	return _c
}

func (_c *MockCache_DeriveObjectID_Call) Return(id string, err error) *MockCache_DeriveObjectID_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockCache_DeriveObjectID_Call) RunAndReturn(run func(context.Context, domain.UniqueKey) (string, error)) *MockCache_DeriveObjectID_Call {
	_c.Call.Return(run)
	return _c
}

// DeriveObjectIdWithAccountSignature provides a mock function with given fields: ctx, uniqueKey
func (_m *MockCache) DeriveObjectIdWithAccountSignature(ctx context.Context, uniqueKey domain.UniqueKey) (string, error) {
	ret := _m.Called(ctx, uniqueKey)

	if len(ret) == 0 {
		panic("no return value specified for DeriveObjectIdWithAccountSignature")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.UniqueKey) (string, error)); ok {
		return rf(ctx, uniqueKey)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.UniqueKey) string); ok {
		r0 = rf(ctx, uniqueKey)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.UniqueKey) error); ok {
		r1 = rf(ctx, uniqueKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_DeriveObjectIdWithAccountSignature_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeriveObjectIdWithAccountSignature'
type MockCache_DeriveObjectIdWithAccountSignature_Call struct {
	*mock.Call
}

// DeriveObjectIdWithAccountSignature is a helper method to define mock.On call
//   - ctx context.Context
//   - uniqueKey domain.UniqueKey
func (_e *MockCache_Expecter) DeriveObjectIdWithAccountSignature(ctx interface{}, uniqueKey interface{}) *MockCache_DeriveObjectIdWithAccountSignature_Call {
	return &MockCache_DeriveObjectIdWithAccountSignature_Call{Call: _e.mock.On("DeriveObjectIdWithAccountSignature", ctx, uniqueKey)}
}

func (_c *MockCache_DeriveObjectIdWithAccountSignature_Call) Run(run func(ctx context.Context, uniqueKey domain.UniqueKey)) *MockCache_DeriveObjectIdWithAccountSignature_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.UniqueKey))
	})
	return _c
}

func (_c *MockCache_DeriveObjectIdWithAccountSignature_Call) Return(id string, err error) *MockCache_DeriveObjectIdWithAccountSignature_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockCache_DeriveObjectIdWithAccountSignature_Call) RunAndReturn(run func(context.Context, domain.UniqueKey) (string, error)) *MockCache_DeriveObjectIdWithAccountSignature_Call {
	_c.Call.Return(run)
	return _c
}

// DeriveTreeObject provides a mock function with given fields: ctx, params
func (_m *MockCache) DeriveTreeObject(ctx context.Context, params objectcache.TreeDerivationParams) (smartblock.SmartBlock, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for DeriveTreeObject")
	}

	var r0 smartblock.SmartBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, objectcache.TreeDerivationParams) (smartblock.SmartBlock, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, objectcache.TreeDerivationParams) smartblock.SmartBlock); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(smartblock.SmartBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, objectcache.TreeDerivationParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_DeriveTreeObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeriveTreeObject'
type MockCache_DeriveTreeObject_Call struct {
	*mock.Call
}

// DeriveTreeObject is a helper method to define mock.On call
//   - ctx context.Context
//   - params objectcache.TreeDerivationParams
func (_e *MockCache_Expecter) DeriveTreeObject(ctx interface{}, params interface{}) *MockCache_DeriveTreeObject_Call {
	return &MockCache_DeriveTreeObject_Call{Call: _e.mock.On("DeriveTreeObject", ctx, params)}
}

func (_c *MockCache_DeriveTreeObject_Call) Run(run func(ctx context.Context, params objectcache.TreeDerivationParams)) *MockCache_DeriveTreeObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(objectcache.TreeDerivationParams))
	})
	return _c
}

func (_c *MockCache_DeriveTreeObject_Call) Return(sb smartblock.SmartBlock, err error) *MockCache_DeriveTreeObject_Call {
	_c.Call.Return(sb, err)
	return _c
}

func (_c *MockCache_DeriveTreeObject_Call) RunAndReturn(run func(context.Context, objectcache.TreeDerivationParams) (smartblock.SmartBlock, error)) *MockCache_DeriveTreeObject_Call {
	_c.Call.Return(run)
	return _c
}

// DeriveTreeObjectWithAccountSignature provides a mock function with given fields: ctx, params
func (_m *MockCache) DeriveTreeObjectWithAccountSignature(ctx context.Context, params objectcache.TreeDerivationParams) (smartblock.SmartBlock, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for DeriveTreeObjectWithAccountSignature")
	}

	var r0 smartblock.SmartBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, objectcache.TreeDerivationParams) (smartblock.SmartBlock, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, objectcache.TreeDerivationParams) smartblock.SmartBlock); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(smartblock.SmartBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, objectcache.TreeDerivationParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_DeriveTreeObjectWithAccountSignature_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeriveTreeObjectWithAccountSignature'
type MockCache_DeriveTreeObjectWithAccountSignature_Call struct {
	*mock.Call
}

// DeriveTreeObjectWithAccountSignature is a helper method to define mock.On call
//   - ctx context.Context
//   - params objectcache.TreeDerivationParams
func (_e *MockCache_Expecter) DeriveTreeObjectWithAccountSignature(ctx interface{}, params interface{}) *MockCache_DeriveTreeObjectWithAccountSignature_Call {
	return &MockCache_DeriveTreeObjectWithAccountSignature_Call{Call: _e.mock.On("DeriveTreeObjectWithAccountSignature", ctx, params)}
}

func (_c *MockCache_DeriveTreeObjectWithAccountSignature_Call) Run(run func(ctx context.Context, params objectcache.TreeDerivationParams)) *MockCache_DeriveTreeObjectWithAccountSignature_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(objectcache.TreeDerivationParams))
	})
	return _c
}

func (_c *MockCache_DeriveTreeObjectWithAccountSignature_Call) Return(sb smartblock.SmartBlock, err error) *MockCache_DeriveTreeObjectWithAccountSignature_Call {
	_c.Call.Return(sb, err)
	return _c
}

func (_c *MockCache_DeriveTreeObjectWithAccountSignature_Call) RunAndReturn(run func(context.Context, objectcache.TreeDerivationParams) (smartblock.SmartBlock, error)) *MockCache_DeriveTreeObjectWithAccountSignature_Call {
	_c.Call.Return(run)
	return _c
}

// DeriveTreePayload provides a mock function with given fields: ctx, params
func (_m *MockCache) DeriveTreePayload(ctx context.Context, params payloadcreator.PayloadDerivationParams) (treestorage.TreeStorageCreatePayload, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for DeriveTreePayload")
	}

	var r0 treestorage.TreeStorageCreatePayload
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, payloadcreator.PayloadDerivationParams) (treestorage.TreeStorageCreatePayload, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, payloadcreator.PayloadDerivationParams) treestorage.TreeStorageCreatePayload); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(treestorage.TreeStorageCreatePayload)
	}

	if rf, ok := ret.Get(1).(func(context.Context, payloadcreator.PayloadDerivationParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_DeriveTreePayload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeriveTreePayload'
type MockCache_DeriveTreePayload_Call struct {
	*mock.Call
}

// DeriveTreePayload is a helper method to define mock.On call
//   - ctx context.Context
//   - params payloadcreator.PayloadDerivationParams
func (_e *MockCache_Expecter) DeriveTreePayload(ctx interface{}, params interface{}) *MockCache_DeriveTreePayload_Call {
	return &MockCache_DeriveTreePayload_Call{Call: _e.mock.On("DeriveTreePayload", ctx, params)}
}

func (_c *MockCache_DeriveTreePayload_Call) Run(run func(ctx context.Context, params payloadcreator.PayloadDerivationParams)) *MockCache_DeriveTreePayload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(payloadcreator.PayloadDerivationParams))
	})
	return _c
}

func (_c *MockCache_DeriveTreePayload_Call) Return(storagePayload treestorage.TreeStorageCreatePayload, err error) *MockCache_DeriveTreePayload_Call {
	_c.Call.Return(storagePayload, err)
	return _c
}

func (_c *MockCache_DeriveTreePayload_Call) RunAndReturn(run func(context.Context, payloadcreator.PayloadDerivationParams) (treestorage.TreeStorageCreatePayload, error)) *MockCache_DeriveTreePayload_Call {
	_c.Call.Return(run)
	return _c
}

// DoLockedIfNotExists provides a mock function with given fields: objectID, proc
func (_m *MockCache) DoLockedIfNotExists(objectID string, proc func() error) error {
	ret := _m.Called(objectID, proc)

	if len(ret) == 0 {
		panic("no return value specified for DoLockedIfNotExists")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, func() error) error); ok {
		r0 = rf(objectID, proc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCache_DoLockedIfNotExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DoLockedIfNotExists'
type MockCache_DoLockedIfNotExists_Call struct {
	*mock.Call
}

// DoLockedIfNotExists is a helper method to define mock.On call
//   - objectID string
//   - proc func() error
func (_e *MockCache_Expecter) DoLockedIfNotExists(objectID interface{}, proc interface{}) *MockCache_DoLockedIfNotExists_Call {
	return &MockCache_DoLockedIfNotExists_Call{Call: _e.mock.On("DoLockedIfNotExists", objectID, proc)}
}

func (_c *MockCache_DoLockedIfNotExists_Call) Run(run func(objectID string, proc func() error)) *MockCache_DoLockedIfNotExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(func() error))
	})
	return _c
}

func (_c *MockCache_DoLockedIfNotExists_Call) Return(_a0 error) *MockCache_DoLockedIfNotExists_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_DoLockedIfNotExists_Call) RunAndReturn(run func(string, func() error) error) *MockCache_DoLockedIfNotExists_Call {
	_c.Call.Return(run)
	return _c
}

// GetObject provides a mock function with given fields: ctx, id
func (_m *MockCache) GetObject(ctx context.Context, id string) (smartblock.SmartBlock, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetObject")
	}

	var r0 smartblock.SmartBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (smartblock.SmartBlock, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) smartblock.SmartBlock); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(smartblock.SmartBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObject'
type MockCache_GetObject_Call struct {
	*mock.Call
}

// GetObject is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockCache_Expecter) GetObject(ctx interface{}, id interface{}) *MockCache_GetObject_Call {
	return &MockCache_GetObject_Call{Call: _e.mock.On("GetObject", ctx, id)}
}

func (_c *MockCache_GetObject_Call) Run(run func(ctx context.Context, id string)) *MockCache_GetObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCache_GetObject_Call) Return(sb smartblock.SmartBlock, err error) *MockCache_GetObject_Call {
	_c.Call.Return(sb, err)
	return _c
}

func (_c *MockCache_GetObject_Call) RunAndReturn(run func(context.Context, string) (smartblock.SmartBlock, error)) *MockCache_GetObject_Call {
	_c.Call.Return(run)
	return _c
}

// GetObjectWithTimeout provides a mock function with given fields: ctx, id
func (_m *MockCache) GetObjectWithTimeout(ctx context.Context, id string) (smartblock.SmartBlock, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetObjectWithTimeout")
	}

	var r0 smartblock.SmartBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (smartblock.SmartBlock, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) smartblock.SmartBlock); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(smartblock.SmartBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetObjectWithTimeout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObjectWithTimeout'
type MockCache_GetObjectWithTimeout_Call struct {
	*mock.Call
}

// GetObjectWithTimeout is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockCache_Expecter) GetObjectWithTimeout(ctx interface{}, id interface{}) *MockCache_GetObjectWithTimeout_Call {
	return &MockCache_GetObjectWithTimeout_Call{Call: _e.mock.On("GetObjectWithTimeout", ctx, id)}
}

func (_c *MockCache_GetObjectWithTimeout_Call) Run(run func(ctx context.Context, id string)) *MockCache_GetObjectWithTimeout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCache_GetObjectWithTimeout_Call) Return(sb smartblock.SmartBlock, err error) *MockCache_GetObjectWithTimeout_Call {
	_c.Call.Return(sb, err)
	return _c
}

func (_c *MockCache_GetObjectWithTimeout_Call) RunAndReturn(run func(context.Context, string) (smartblock.SmartBlock, error)) *MockCache_GetObjectWithTimeout_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: ctx, objectID
func (_m *MockCache) Remove(ctx context.Context, objectID string) error {
	ret := _m.Called(ctx, objectID)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, objectID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCache_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type MockCache_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - ctx context.Context
//   - objectID string
func (_e *MockCache_Expecter) Remove(ctx interface{}, objectID interface{}) *MockCache_Remove_Call {
	return &MockCache_Remove_Call{Call: _e.mock.On("Remove", ctx, objectID)}
}

func (_c *MockCache_Remove_Call) Run(run func(ctx context.Context, objectID string)) *MockCache_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCache_Remove_Call) Return(_a0 error) *MockCache_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_Remove_Call) RunAndReturn(run func(context.Context, string) error) *MockCache_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCache creates a new instance of MockCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCache {
	mock := &MockCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
