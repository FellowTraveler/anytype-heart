// Code generated by mockery v2.32.0. DO NOT EDIT.

package mock_system_object

import (
	context "context"

	app "github.com/anyproto/any-sync/app"

	domain "github.com/anyproto/anytype-heart/core/domain"

	mock "github.com/stretchr/testify/mock"

	model "github.com/anyproto/anytype-heart/pkg/lib/pb/model"

	pbtypes "github.com/anyproto/anytype-heart/util/pbtypes"

	relationutils "github.com/anyproto/anytype-heart/core/system_object/relationutils"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

type MockService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockService) EXPECT() *MockService_Expecter {
	return &MockService_Expecter{mock: &_m.Mock}
}

// FetchRelationByKey provides a mock function with given fields: spaceId, key
func (_m *MockService) FetchRelationByKey(spaceId string, key string) (*relationutils.Relation, error) {
	ret := _m.Called(spaceId, key)

	var r0 *relationutils.Relation
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*relationutils.Relation, error)); ok {
		return rf(spaceId, key)
	}
	if rf, ok := ret.Get(0).(func(string, string) *relationutils.Relation); ok {
		r0 = rf(spaceId, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*relationutils.Relation)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(spaceId, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_FetchRelationByKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchRelationByKey'
type MockService_FetchRelationByKey_Call struct {
	*mock.Call
}

// FetchRelationByKey is a helper method to define mock.On call
//   - spaceId string
//   - key string
func (_e *MockService_Expecter) FetchRelationByKey(spaceId interface{}, key interface{}) *MockService_FetchRelationByKey_Call {
	return &MockService_FetchRelationByKey_Call{Call: _e.mock.On("FetchRelationByKey", spaceId, key)}
}

func (_c *MockService_FetchRelationByKey_Call) Run(run func(spaceId string, key string)) *MockService_FetchRelationByKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockService_FetchRelationByKey_Call) Return(relation *relationutils.Relation, err error) *MockService_FetchRelationByKey_Call {
	_c.Call.Return(relation, err)
	return _c
}

func (_c *MockService_FetchRelationByKey_Call) RunAndReturn(run func(string, string) (*relationutils.Relation, error)) *MockService_FetchRelationByKey_Call {
	_c.Call.Return(run)
	return _c
}

// FetchRelationByKeys provides a mock function with given fields: spaceId, keys
func (_m *MockService) FetchRelationByKeys(spaceId string, keys ...string) (relationutils.Relations, error) {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, spaceId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 relationutils.Relations
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...string) (relationutils.Relations, error)); ok {
		return rf(spaceId, keys...)
	}
	if rf, ok := ret.Get(0).(func(string, ...string) relationutils.Relations); ok {
		r0 = rf(spaceId, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(relationutils.Relations)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(spaceId, keys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_FetchRelationByKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchRelationByKeys'
type MockService_FetchRelationByKeys_Call struct {
	*mock.Call
}

// FetchRelationByKeys is a helper method to define mock.On call
//   - spaceId string
//   - keys ...string
func (_e *MockService_Expecter) FetchRelationByKeys(spaceId interface{}, keys ...interface{}) *MockService_FetchRelationByKeys_Call {
	return &MockService_FetchRelationByKeys_Call{Call: _e.mock.On("FetchRelationByKeys",
		append([]interface{}{spaceId}, keys...)...)}
}

func (_c *MockService_FetchRelationByKeys_Call) Run(run func(spaceId string, keys ...string)) *MockService_FetchRelationByKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockService_FetchRelationByKeys_Call) Return(relations relationutils.Relations, err error) *MockService_FetchRelationByKeys_Call {
	_c.Call.Return(relations, err)
	return _c
}

func (_c *MockService_FetchRelationByKeys_Call) RunAndReturn(run func(string, ...string) (relationutils.Relations, error)) *MockService_FetchRelationByKeys_Call {
	_c.Call.Return(run)
	return _c
}

// FetchRelationByLinks provides a mock function with given fields: spaceId, links
func (_m *MockService) FetchRelationByLinks(spaceId string, links pbtypes.RelationLinks) (relationutils.Relations, error) {
	ret := _m.Called(spaceId, links)

	var r0 relationutils.Relations
	var r1 error
	if rf, ok := ret.Get(0).(func(string, pbtypes.RelationLinks) (relationutils.Relations, error)); ok {
		return rf(spaceId, links)
	}
	if rf, ok := ret.Get(0).(func(string, pbtypes.RelationLinks) relationutils.Relations); ok {
		r0 = rf(spaceId, links)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(relationutils.Relations)
		}
	}

	if rf, ok := ret.Get(1).(func(string, pbtypes.RelationLinks) error); ok {
		r1 = rf(spaceId, links)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_FetchRelationByLinks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchRelationByLinks'
type MockService_FetchRelationByLinks_Call struct {
	*mock.Call
}

// FetchRelationByLinks is a helper method to define mock.On call
//   - spaceId string
//   - links pbtypes.RelationLinks
func (_e *MockService_Expecter) FetchRelationByLinks(spaceId interface{}, links interface{}) *MockService_FetchRelationByLinks_Call {
	return &MockService_FetchRelationByLinks_Call{Call: _e.mock.On("FetchRelationByLinks", spaceId, links)}
}

func (_c *MockService_FetchRelationByLinks_Call) Run(run func(spaceId string, links pbtypes.RelationLinks)) *MockService_FetchRelationByLinks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(pbtypes.RelationLinks))
	})
	return _c
}

func (_c *MockService_FetchRelationByLinks_Call) Return(relations relationutils.Relations, err error) *MockService_FetchRelationByLinks_Call {
	_c.Call.Return(relations, err)
	return _c
}

func (_c *MockService_FetchRelationByLinks_Call) RunAndReturn(run func(string, pbtypes.RelationLinks) (relationutils.Relations, error)) *MockService_FetchRelationByLinks_Call {
	_c.Call.Return(run)
	return _c
}

// GetObjectByUniqueKey provides a mock function with given fields: spaceId, uniqueKey
func (_m *MockService) GetObjectByUniqueKey(spaceId string, uniqueKey domain.UniqueKey) (*model.ObjectDetails, error) {
	ret := _m.Called(spaceId, uniqueKey)

	var r0 *model.ObjectDetails
	var r1 error
	if rf, ok := ret.Get(0).(func(string, domain.UniqueKey) (*model.ObjectDetails, error)); ok {
		return rf(spaceId, uniqueKey)
	}
	if rf, ok := ret.Get(0).(func(string, domain.UniqueKey) *model.ObjectDetails); ok {
		r0 = rf(spaceId, uniqueKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ObjectDetails)
		}
	}

	if rf, ok := ret.Get(1).(func(string, domain.UniqueKey) error); ok {
		r1 = rf(spaceId, uniqueKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetObjectByUniqueKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObjectByUniqueKey'
type MockService_GetObjectByUniqueKey_Call struct {
	*mock.Call
}

// GetObjectByUniqueKey is a helper method to define mock.On call
//   - spaceId string
//   - uniqueKey domain.UniqueKey
func (_e *MockService_Expecter) GetObjectByUniqueKey(spaceId interface{}, uniqueKey interface{}) *MockService_GetObjectByUniqueKey_Call {
	return &MockService_GetObjectByUniqueKey_Call{Call: _e.mock.On("GetObjectByUniqueKey", spaceId, uniqueKey)}
}

func (_c *MockService_GetObjectByUniqueKey_Call) Run(run func(spaceId string, uniqueKey domain.UniqueKey)) *MockService_GetObjectByUniqueKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(domain.UniqueKey))
	})
	return _c
}

func (_c *MockService_GetObjectByUniqueKey_Call) Return(_a0 *model.ObjectDetails, _a1 error) *MockService_GetObjectByUniqueKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_GetObjectByUniqueKey_Call) RunAndReturn(run func(string, domain.UniqueKey) (*model.ObjectDetails, error)) *MockService_GetObjectByUniqueKey_Call {
	_c.Call.Return(run)
	return _c
}

// GetObjectIdByUniqueKey provides a mock function with given fields: ctx, spaceId, key
func (_m *MockService) GetObjectIdByUniqueKey(ctx context.Context, spaceId string, key domain.UniqueKey) (string, error) {
	ret := _m.Called(ctx, spaceId, key)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.UniqueKey) (string, error)); ok {
		return rf(ctx, spaceId, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.UniqueKey) string); ok {
		r0 = rf(ctx, spaceId, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, domain.UniqueKey) error); ok {
		r1 = rf(ctx, spaceId, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetObjectIdByUniqueKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObjectIdByUniqueKey'
type MockService_GetObjectIdByUniqueKey_Call struct {
	*mock.Call
}

// GetObjectIdByUniqueKey is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
//   - key domain.UniqueKey
func (_e *MockService_Expecter) GetObjectIdByUniqueKey(ctx interface{}, spaceId interface{}, key interface{}) *MockService_GetObjectIdByUniqueKey_Call {
	return &MockService_GetObjectIdByUniqueKey_Call{Call: _e.mock.On("GetObjectIdByUniqueKey", ctx, spaceId, key)}
}

func (_c *MockService_GetObjectIdByUniqueKey_Call) Run(run func(ctx context.Context, spaceId string, key domain.UniqueKey)) *MockService_GetObjectIdByUniqueKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(domain.UniqueKey))
	})
	return _c
}

func (_c *MockService_GetObjectIdByUniqueKey_Call) Return(id string, err error) *MockService_GetObjectIdByUniqueKey_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockService_GetObjectIdByUniqueKey_Call) RunAndReturn(run func(context.Context, string, domain.UniqueKey) (string, error)) *MockService_GetObjectIdByUniqueKey_Call {
	_c.Call.Return(run)
	return _c
}

// GetObjectType provides a mock function with given fields: url
func (_m *MockService) GetObjectType(url string) (*model.ObjectType, error) {
	ret := _m.Called(url)

	var r0 *model.ObjectType
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.ObjectType, error)); ok {
		return rf(url)
	}
	if rf, ok := ret.Get(0).(func(string) *model.ObjectType); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ObjectType)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetObjectType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObjectType'
type MockService_GetObjectType_Call struct {
	*mock.Call
}

// GetObjectType is a helper method to define mock.On call
//   - url string
func (_e *MockService_Expecter) GetObjectType(url interface{}) *MockService_GetObjectType_Call {
	return &MockService_GetObjectType_Call{Call: _e.mock.On("GetObjectType", url)}
}

func (_c *MockService_GetObjectType_Call) Run(run func(url string)) *MockService_GetObjectType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockService_GetObjectType_Call) Return(_a0 *model.ObjectType, _a1 error) *MockService_GetObjectType_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_GetObjectType_Call) RunAndReturn(run func(string) (*model.ObjectType, error)) *MockService_GetObjectType_Call {
	_c.Call.Return(run)
	return _c
}

// GetObjectTypes provides a mock function with given fields: urls
func (_m *MockService) GetObjectTypes(urls []string) ([]*model.ObjectType, error) {
	ret := _m.Called(urls)

	var r0 []*model.ObjectType
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]*model.ObjectType, error)); ok {
		return rf(urls)
	}
	if rf, ok := ret.Get(0).(func([]string) []*model.ObjectType); ok {
		r0 = rf(urls)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ObjectType)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(urls)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetObjectTypes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObjectTypes'
type MockService_GetObjectTypes_Call struct {
	*mock.Call
}

// GetObjectTypes is a helper method to define mock.On call
//   - urls []string
func (_e *MockService_Expecter) GetObjectTypes(urls interface{}) *MockService_GetObjectTypes_Call {
	return &MockService_GetObjectTypes_Call{Call: _e.mock.On("GetObjectTypes", urls)}
}

func (_c *MockService_GetObjectTypes_Call) Run(run func(urls []string)) *MockService_GetObjectTypes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *MockService_GetObjectTypes_Call) Return(ots []*model.ObjectType, err error) *MockService_GetObjectTypes_Call {
	_c.Call.Return(ots, err)
	return _c
}

func (_c *MockService_GetObjectTypes_Call) RunAndReturn(run func([]string) ([]*model.ObjectType, error)) *MockService_GetObjectTypes_Call {
	_c.Call.Return(run)
	return _c
}

// GetRelationByID provides a mock function with given fields: id
func (_m *MockService) GetRelationByID(id string) (*model.Relation, error) {
	ret := _m.Called(id)

	var r0 *model.Relation
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Relation, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Relation); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Relation)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetRelationByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRelationByID'
type MockService_GetRelationByID_Call struct {
	*mock.Call
}

// GetRelationByID is a helper method to define mock.On call
//   - id string
func (_e *MockService_Expecter) GetRelationByID(id interface{}) *MockService_GetRelationByID_Call {
	return &MockService_GetRelationByID_Call{Call: _e.mock.On("GetRelationByID", id)}
}

func (_c *MockService_GetRelationByID_Call) Run(run func(id string)) *MockService_GetRelationByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockService_GetRelationByID_Call) Return(relation *model.Relation, err error) *MockService_GetRelationByID_Call {
	_c.Call.Return(relation, err)
	return _c
}

func (_c *MockService_GetRelationByID_Call) RunAndReturn(run func(string) (*model.Relation, error)) *MockService_GetRelationByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetRelationByKey provides a mock function with given fields: key
func (_m *MockService) GetRelationByKey(key string) (*model.Relation, error) {
	ret := _m.Called(key)

	var r0 *model.Relation
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Relation, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Relation); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Relation)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetRelationByKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRelationByKey'
type MockService_GetRelationByKey_Call struct {
	*mock.Call
}

// GetRelationByKey is a helper method to define mock.On call
//   - key string
func (_e *MockService_Expecter) GetRelationByKey(key interface{}) *MockService_GetRelationByKey_Call {
	return &MockService_GetRelationByKey_Call{Call: _e.mock.On("GetRelationByKey", key)}
}

func (_c *MockService_GetRelationByKey_Call) Run(run func(key string)) *MockService_GetRelationByKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockService_GetRelationByKey_Call) Return(relation *model.Relation, err error) *MockService_GetRelationByKey_Call {
	_c.Call.Return(relation, err)
	return _c
}

func (_c *MockService_GetRelationByKey_Call) RunAndReturn(run func(string) (*model.Relation, error)) *MockService_GetRelationByKey_Call {
	_c.Call.Return(run)
	return _c
}

// GetRelationIdByKey provides a mock function with given fields: ctx, spaceId, key
func (_m *MockService) GetRelationIdByKey(ctx context.Context, spaceId string, key domain.RelationKey) (string, error) {
	ret := _m.Called(ctx, spaceId, key)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.RelationKey) (string, error)); ok {
		return rf(ctx, spaceId, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.RelationKey) string); ok {
		r0 = rf(ctx, spaceId, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, domain.RelationKey) error); ok {
		r1 = rf(ctx, spaceId, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetRelationIdByKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRelationIdByKey'
type MockService_GetRelationIdByKey_Call struct {
	*mock.Call
}

// GetRelationIdByKey is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
//   - key domain.RelationKey
func (_e *MockService_Expecter) GetRelationIdByKey(ctx interface{}, spaceId interface{}, key interface{}) *MockService_GetRelationIdByKey_Call {
	return &MockService_GetRelationIdByKey_Call{Call: _e.mock.On("GetRelationIdByKey", ctx, spaceId, key)}
}

func (_c *MockService_GetRelationIdByKey_Call) Run(run func(ctx context.Context, spaceId string, key domain.RelationKey)) *MockService_GetRelationIdByKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(domain.RelationKey))
	})
	return _c
}

func (_c *MockService_GetRelationIdByKey_Call) Return(id string, err error) *MockService_GetRelationIdByKey_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockService_GetRelationIdByKey_Call) RunAndReturn(run func(context.Context, string, domain.RelationKey) (string, error)) *MockService_GetRelationIdByKey_Call {
	_c.Call.Return(run)
	return _c
}

// GetTypeIdByKey provides a mock function with given fields: ctx, spaceId, key
func (_m *MockService) GetTypeIdByKey(ctx context.Context, spaceId string, key domain.TypeKey) (string, error) {
	ret := _m.Called(ctx, spaceId, key)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.TypeKey) (string, error)); ok {
		return rf(ctx, spaceId, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.TypeKey) string); ok {
		r0 = rf(ctx, spaceId, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, domain.TypeKey) error); ok {
		r1 = rf(ctx, spaceId, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetTypeIdByKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTypeIdByKey'
type MockService_GetTypeIdByKey_Call struct {
	*mock.Call
}

// GetTypeIdByKey is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
//   - key domain.TypeKey
func (_e *MockService_Expecter) GetTypeIdByKey(ctx interface{}, spaceId interface{}, key interface{}) *MockService_GetTypeIdByKey_Call {
	return &MockService_GetTypeIdByKey_Call{Call: _e.mock.On("GetTypeIdByKey", ctx, spaceId, key)}
}

func (_c *MockService_GetTypeIdByKey_Call) Run(run func(ctx context.Context, spaceId string, key domain.TypeKey)) *MockService_GetTypeIdByKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(domain.TypeKey))
	})
	return _c
}

func (_c *MockService_GetTypeIdByKey_Call) Return(id string, err error) *MockService_GetTypeIdByKey_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockService_GetTypeIdByKey_Call) RunAndReturn(run func(context.Context, string, domain.TypeKey) (string, error)) *MockService_GetTypeIdByKey_Call {
	_c.Call.Return(run)
	return _c
}

// HasObjectType provides a mock function with given fields: id
func (_m *MockService) HasObjectType(id string) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_HasObjectType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasObjectType'
type MockService_HasObjectType_Call struct {
	*mock.Call
}

// HasObjectType is a helper method to define mock.On call
//   - id string
func (_e *MockService_Expecter) HasObjectType(id interface{}) *MockService_HasObjectType_Call {
	return &MockService_HasObjectType_Call{Call: _e.mock.On("HasObjectType", id)}
}

func (_c *MockService_HasObjectType_Call) Run(run func(id string)) *MockService_HasObjectType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockService_HasObjectType_Call) Return(_a0 bool, _a1 error) *MockService_HasObjectType_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_HasObjectType_Call) RunAndReturn(run func(string) (bool, error)) *MockService_HasObjectType_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockService) Init(a *app.App) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.App) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockService_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockService_Expecter) Init(a interface{}) *MockService_Init_Call {
	return &MockService_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockService_Init_Call) Run(run func(a *app.App)) *MockService_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockService_Init_Call) Return(err error) *MockService_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockService_Init_Call) RunAndReturn(run func(*app.App) error) *MockService_Init_Call {
	_c.Call.Return(run)
	return _c
}

// ListAllRelations provides a mock function with given fields: spaceId
func (_m *MockService) ListAllRelations(spaceId string) (relationutils.Relations, error) {
	ret := _m.Called(spaceId)

	var r0 relationutils.Relations
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (relationutils.Relations, error)); ok {
		return rf(spaceId)
	}
	if rf, ok := ret.Get(0).(func(string) relationutils.Relations); ok {
		r0 = rf(spaceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(relationutils.Relations)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(spaceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_ListAllRelations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAllRelations'
type MockService_ListAllRelations_Call struct {
	*mock.Call
}

// ListAllRelations is a helper method to define mock.On call
//   - spaceId string
func (_e *MockService_Expecter) ListAllRelations(spaceId interface{}) *MockService_ListAllRelations_Call {
	return &MockService_ListAllRelations_Call{Call: _e.mock.On("ListAllRelations", spaceId)}
}

func (_c *MockService_ListAllRelations_Call) Run(run func(spaceId string)) *MockService_ListAllRelations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockService_ListAllRelations_Call) Return(relations relationutils.Relations, err error) *MockService_ListAllRelations_Call {
	_c.Call.Return(relations, err)
	return _c
}

func (_c *MockService_ListAllRelations_Call) RunAndReturn(run func(string) (relationutils.Relations, error)) *MockService_ListAllRelations_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockService) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockService_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockService_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockService_Expecter) Name() *MockService_Name_Call {
	return &MockService_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockService_Name_Call) Run(run func()) *MockService_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_Name_Call) Return(name string) *MockService_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockService_Name_Call) RunAndReturn(run func() string) *MockService_Name_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
