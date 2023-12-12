// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore (interfaces: ObjectStore)
//
// Generated by this command:
//
//	mockgen -package testMock -destination objectstore_mock.go github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore ObjectStore
//
// Package testMock is a generated GoMock package.
package testMock

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	coordinatorproto "github.com/anyproto/any-sync/coordinator/coordinatorproto"
	domain "github.com/anyproto/anytype-heart/core/domain"
	relationutils "github.com/anyproto/anytype-heart/core/relationutils"
	database "github.com/anyproto/anytype-heart/pkg/lib/database"
	ftsearch "github.com/anyproto/anytype-heart/pkg/lib/localstore/ftsearch"
	model "github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	pbtypes "github.com/anyproto/anytype-heart/util/pbtypes"
	types "github.com/gogo/protobuf/types"
	gomock "go.uber.org/mock/gomock"
)

// MockObjectStore is a mock of ObjectStore interface.
type MockObjectStore struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStoreMockRecorder
}

// MockObjectStoreMockRecorder is the mock recorder for MockObjectStore.
type MockObjectStoreMockRecorder struct {
	mock *MockObjectStore
}

// NewMockObjectStore creates a new mock instance.
func NewMockObjectStore(ctrl *gomock.Controller) *MockObjectStore {
	mock := &MockObjectStore{ctrl: ctrl}
	mock.recorder = &MockObjectStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectStore) EXPECT() *MockObjectStoreMockRecorder {
	return m.recorder
}

// AddToIndexQueue mocks base method.
func (m *MockObjectStore) AddToIndexQueue(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToIndexQueue", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToIndexQueue indicates an expected call of AddToIndexQueue.
func (mr *MockObjectStoreMockRecorder) AddToIndexQueue(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToIndexQueue", reflect.TypeOf((*MockObjectStore)(nil).AddToIndexQueue), arg0)
}

// Close mocks base method.
func (m *MockObjectStore) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockObjectStoreMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockObjectStore)(nil).Close), arg0)
}

// DeleteDetails mocks base method.
func (m *MockObjectStore) DeleteDetails(arg0 ...string) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteDetails", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDetails indicates an expected call of DeleteDetails.
func (mr *MockObjectStoreMockRecorder) DeleteDetails(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDetails", reflect.TypeOf((*MockObjectStore)(nil).DeleteDetails), arg0...)
}

// DeleteObject mocks base method.
func (m *MockObjectStore) DeleteObject(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockObjectStoreMockRecorder) DeleteObject(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockObjectStore)(nil).DeleteObject), arg0)
}

// DeleteVirtualSpace mocks base method.
func (m *MockObjectStore) DeleteVirtualSpace(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVirtualSpace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVirtualSpace indicates an expected call of DeleteVirtualSpace.
func (mr *MockObjectStoreMockRecorder) DeleteVirtualSpace(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVirtualSpace", reflect.TypeOf((*MockObjectStore)(nil).DeleteVirtualSpace), arg0)
}

// EraseIndexes mocks base method.
func (m *MockObjectStore) EraseIndexes(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EraseIndexes", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EraseIndexes indicates an expected call of EraseIndexes.
func (mr *MockObjectStoreMockRecorder) EraseIndexes(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EraseIndexes", reflect.TypeOf((*MockObjectStore)(nil).EraseIndexes), arg0)
}

// FTSearch mocks base method.
func (m *MockObjectStore) FTSearch() ftsearch.FTSearch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FTSearch")
	ret0, _ := ret[0].(ftsearch.FTSearch)
	return ret0
}

// FTSearch indicates an expected call of FTSearch.
func (mr *MockObjectStoreMockRecorder) FTSearch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FTSearch", reflect.TypeOf((*MockObjectStore)(nil).FTSearch))
}

// FetchRelationByKey mocks base method.
func (m *MockObjectStore) FetchRelationByKey(arg0, arg1 string) (*relationutils.Relation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRelationByKey", arg0, arg1)
	ret0, _ := ret[0].(*relationutils.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRelationByKey indicates an expected call of FetchRelationByKey.
func (mr *MockObjectStoreMockRecorder) FetchRelationByKey(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRelationByKey", reflect.TypeOf((*MockObjectStore)(nil).FetchRelationByKey), arg0, arg1)
}

// FetchRelationByKeys mocks base method.
func (m *MockObjectStore) FetchRelationByKeys(arg0 string, arg1 ...string) (relationutils.Relations, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchRelationByKeys", varargs...)
	ret0, _ := ret[0].(relationutils.Relations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRelationByKeys indicates an expected call of FetchRelationByKeys.
func (mr *MockObjectStoreMockRecorder) FetchRelationByKeys(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRelationByKeys", reflect.TypeOf((*MockObjectStore)(nil).FetchRelationByKeys), varargs...)
}

// FetchRelationByLinks mocks base method.
func (m *MockObjectStore) FetchRelationByLinks(arg0 string, arg1 pbtypes.RelationLinks) (relationutils.Relations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRelationByLinks", arg0, arg1)
	ret0, _ := ret[0].(relationutils.Relations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRelationByLinks indicates an expected call of FetchRelationByLinks.
func (mr *MockObjectStoreMockRecorder) FetchRelationByLinks(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRelationByLinks", reflect.TypeOf((*MockObjectStore)(nil).FetchRelationByLinks), arg0, arg1)
}

// GetAccountStatus mocks base method.
func (m *MockObjectStore) GetAccountStatus() (*coordinatorproto.SpaceStatusPayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountStatus")
	ret0, _ := ret[0].(*coordinatorproto.SpaceStatusPayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountStatus indicates an expected call of GetAccountStatus.
func (mr *MockObjectStoreMockRecorder) GetAccountStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountStatus", reflect.TypeOf((*MockObjectStore)(nil).GetAccountStatus))
}

// GetByIDs mocks base method.
func (m *MockObjectStore) GetByIDs(arg0 string, arg1 []string) ([]*model.ObjectInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", arg0, arg1)
	ret0, _ := ret[0].([]*model.ObjectInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs.
func (mr *MockObjectStoreMockRecorder) GetByIDs(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockObjectStore)(nil).GetByIDs), arg0, arg1)
}

// GetChecksums mocks base method.
func (m *MockObjectStore) GetChecksums(arg0 string) (*model.ObjectStoreChecksums, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChecksums", arg0)
	ret0, _ := ret[0].(*model.ObjectStoreChecksums)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChecksums indicates an expected call of GetChecksums.
func (mr *MockObjectStoreMockRecorder) GetChecksums(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChecksums", reflect.TypeOf((*MockObjectStore)(nil).GetChecksums), arg0)
}

// GetDetails mocks base method.
func (m *MockObjectStore) GetDetails(arg0 string) (*model.ObjectDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetails", arg0)
	ret0, _ := ret[0].(*model.ObjectDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetails indicates an expected call of GetDetails.
func (mr *MockObjectStoreMockRecorder) GetDetails(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetails", reflect.TypeOf((*MockObjectStore)(nil).GetDetails), arg0)
}

// GetGlobalChecksums mocks base method.
func (m *MockObjectStore) GetGlobalChecksums() (*model.ObjectStoreChecksums, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGlobalChecksums")
	ret0, _ := ret[0].(*model.ObjectStoreChecksums)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGlobalChecksums indicates an expected call of GetGlobalChecksums.
func (mr *MockObjectStoreMockRecorder) GetGlobalChecksums() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGlobalChecksums", reflect.TypeOf((*MockObjectStore)(nil).GetGlobalChecksums))
}

// GetInboundLinksByID mocks base method.
func (m *MockObjectStore) GetInboundLinksByID(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInboundLinksByID", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInboundLinksByID indicates an expected call of GetInboundLinksByID.
func (mr *MockObjectStoreMockRecorder) GetInboundLinksByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInboundLinksByID", reflect.TypeOf((*MockObjectStore)(nil).GetInboundLinksByID), arg0)
}

// GetLastIndexedHeadsHash mocks base method.
func (m *MockObjectStore) GetLastIndexedHeadsHash(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastIndexedHeadsHash", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastIndexedHeadsHash indicates an expected call of GetLastIndexedHeadsHash.
func (mr *MockObjectStoreMockRecorder) GetLastIndexedHeadsHash(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastIndexedHeadsHash", reflect.TypeOf((*MockObjectStore)(nil).GetLastIndexedHeadsHash), arg0)
}

// GetObjectByUniqueKey mocks base method.
func (m *MockObjectStore) GetObjectByUniqueKey(arg0 string, arg1 domain.UniqueKey) (*model.ObjectDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectByUniqueKey", arg0, arg1)
	ret0, _ := ret[0].(*model.ObjectDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectByUniqueKey indicates an expected call of GetObjectByUniqueKey.
func (mr *MockObjectStoreMockRecorder) GetObjectByUniqueKey(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectByUniqueKey", reflect.TypeOf((*MockObjectStore)(nil).GetObjectByUniqueKey), arg0, arg1)
}

// GetObjectType mocks base method.
func (m *MockObjectStore) GetObjectType(arg0 string) (*model.ObjectType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectType", arg0)
	ret0, _ := ret[0].(*model.ObjectType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectType indicates an expected call of GetObjectType.
func (mr *MockObjectStoreMockRecorder) GetObjectType(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectType", reflect.TypeOf((*MockObjectStore)(nil).GetObjectType), arg0)
}

// GetOutboundLinksByID mocks base method.
func (m *MockObjectStore) GetOutboundLinksByID(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutboundLinksByID", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutboundLinksByID indicates an expected call of GetOutboundLinksByID.
func (mr *MockObjectStoreMockRecorder) GetOutboundLinksByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutboundLinksByID", reflect.TypeOf((*MockObjectStore)(nil).GetOutboundLinksByID), arg0)
}

// GetRelationByID mocks base method.
func (m *MockObjectStore) GetRelationByID(arg0 string) (*model.Relation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelationByID", arg0)
	ret0, _ := ret[0].(*model.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationByID indicates an expected call of GetRelationByID.
func (mr *MockObjectStoreMockRecorder) GetRelationByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationByID", reflect.TypeOf((*MockObjectStore)(nil).GetRelationByID), arg0)
}

// GetRelationByKey mocks base method.
func (m *MockObjectStore) GetRelationByKey(arg0 string) (*model.Relation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelationByKey", arg0)
	ret0, _ := ret[0].(*model.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationByKey indicates an expected call of GetRelationByKey.
func (mr *MockObjectStoreMockRecorder) GetRelationByKey(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationByKey", reflect.TypeOf((*MockObjectStore)(nil).GetRelationByKey), arg0)
}

// GetRelationLink mocks base method.
func (m *MockObjectStore) GetRelationLink(arg0, arg1 string) (*model.RelationLink, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelationLink", arg0, arg1)
	ret0, _ := ret[0].(*model.RelationLink)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationLink indicates an expected call of GetRelationLink.
func (mr *MockObjectStoreMockRecorder) GetRelationLink(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationLink", reflect.TypeOf((*MockObjectStore)(nil).GetRelationLink), arg0, arg1)
}

// GetUniqueKeyById mocks base method.
func (m *MockObjectStore) GetUniqueKeyById(arg0 string) (domain.UniqueKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUniqueKeyById", arg0)
	ret0, _ := ret[0].(domain.UniqueKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUniqueKeyById indicates an expected call of GetUniqueKeyById.
func (mr *MockObjectStoreMockRecorder) GetUniqueKeyById(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUniqueKeyById", reflect.TypeOf((*MockObjectStore)(nil).GetUniqueKeyById), arg0)
}

// GetWithLinksInfoByID mocks base method.
func (m *MockObjectStore) GetWithLinksInfoByID(arg0, arg1 string) (*model.ObjectInfoWithLinks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithLinksInfoByID", arg0, arg1)
	ret0, _ := ret[0].(*model.ObjectInfoWithLinks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithLinksInfoByID indicates an expected call of GetWithLinksInfoByID.
func (mr *MockObjectStoreMockRecorder) GetWithLinksInfoByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithLinksInfoByID", reflect.TypeOf((*MockObjectStore)(nil).GetWithLinksInfoByID), arg0, arg1)
}

// HasIDs mocks base method.
func (m *MockObjectStore) HasIDs(arg0 ...string) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HasIDs", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasIDs indicates an expected call of HasIDs.
func (mr *MockObjectStoreMockRecorder) HasIDs(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasIDs", reflect.TypeOf((*MockObjectStore)(nil).HasIDs), arg0...)
}

// Init mocks base method.
func (m *MockObjectStore) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockObjectStoreMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockObjectStore)(nil).Init), arg0)
}

// List mocks base method.
func (m *MockObjectStore) List(arg0 string, arg1 bool) ([]*model.ObjectInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*model.ObjectInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockObjectStoreMockRecorder) List(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockObjectStore)(nil).List), arg0, arg1)
}

// ListAllRelations mocks base method.
func (m *MockObjectStore) ListAllRelations(arg0 string) (relationutils.Relations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllRelations", arg0)
	ret0, _ := ret[0].(relationutils.Relations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllRelations indicates an expected call of ListAllRelations.
func (mr *MockObjectStoreMockRecorder) ListAllRelations(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllRelations", reflect.TypeOf((*MockObjectStore)(nil).ListAllRelations), arg0)
}

// ListIDsFromFullTextQueue mocks base method.
func (m *MockObjectStore) ListIDsFromFullTextQueue() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIDsFromFullTextQueue")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIDsFromFullTextQueue indicates an expected call of ListIDsFromFullTextQueue.
func (mr *MockObjectStoreMockRecorder) ListIDsFromFullTextQueue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIDsFromFullTextQueue", reflect.TypeOf((*MockObjectStore)(nil).ListIDsFromFullTextQueue))
}

// ListIds mocks base method.
func (m *MockObjectStore) ListIds() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIds")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIds indicates an expected call of ListIds.
func (mr *MockObjectStoreMockRecorder) ListIds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIds", reflect.TypeOf((*MockObjectStore)(nil).ListIds))
}

// ListIdsBySpace mocks base method.
func (m *MockObjectStore) ListIdsBySpace(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIdsBySpace", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIdsBySpace indicates an expected call of ListIdsBySpace.
func (mr *MockObjectStoreMockRecorder) ListIdsBySpace(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIdsBySpace", reflect.TypeOf((*MockObjectStore)(nil).ListIdsBySpace), arg0)
}

// ListVirtualSpaces mocks base method.
func (m *MockObjectStore) ListVirtualSpaces() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVirtualSpaces")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVirtualSpaces indicates an expected call of ListVirtualSpaces.
func (mr *MockObjectStoreMockRecorder) ListVirtualSpaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualSpaces", reflect.TypeOf((*MockObjectStore)(nil).ListVirtualSpaces))
}

// ModifyObjectDetails mocks base method.
func (m *MockObjectStore) ModifyObjectDetails(arg0 string, arg1 func(*types.Struct) (*types.Struct, error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyObjectDetails", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyObjectDetails indicates an expected call of ModifyObjectDetails.
func (mr *MockObjectStoreMockRecorder) ModifyObjectDetails(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyObjectDetails", reflect.TypeOf((*MockObjectStore)(nil).ModifyObjectDetails), arg0, arg1)
}

// Name mocks base method.
func (m *MockObjectStore) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockObjectStoreMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockObjectStore)(nil).Name))
}

// Query mocks base method.
func (m *MockObjectStore) Query(arg0 database.Query) ([]database.Record, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0)
	ret0, _ := ret[0].([]database.Record)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Query indicates an expected call of Query.
func (mr *MockObjectStoreMockRecorder) Query(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockObjectStore)(nil).Query), arg0)
}

// QueryByID mocks base method.
func (m *MockObjectStore) QueryByID(arg0 []string) ([]database.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryByID", arg0)
	ret0, _ := ret[0].([]database.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryByID indicates an expected call of QueryByID.
func (mr *MockObjectStoreMockRecorder) QueryByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryByID", reflect.TypeOf((*MockObjectStore)(nil).QueryByID), arg0)
}

// QueryByIDAndSubscribeForChanges mocks base method.
func (m *MockObjectStore) QueryByIDAndSubscribeForChanges(arg0 []string, arg1 database.Subscription) ([]database.Record, func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryByIDAndSubscribeForChanges", arg0, arg1)
	ret0, _ := ret[0].([]database.Record)
	ret1, _ := ret[1].(func())
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryByIDAndSubscribeForChanges indicates an expected call of QueryByIDAndSubscribeForChanges.
func (mr *MockObjectStoreMockRecorder) QueryByIDAndSubscribeForChanges(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryByIDAndSubscribeForChanges", reflect.TypeOf((*MockObjectStore)(nil).QueryByIDAndSubscribeForChanges), arg0, arg1)
}

// QueryObjectIDs mocks base method.
func (m *MockObjectStore) QueryObjectIDs(arg0 database.Query) ([]string, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryObjectIDs", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryObjectIDs indicates an expected call of QueryObjectIDs.
func (mr *MockObjectStoreMockRecorder) QueryObjectIDs(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryObjectIDs", reflect.TypeOf((*MockObjectStore)(nil).QueryObjectIDs), arg0)
}

// QueryRaw mocks base method.
func (m *MockObjectStore) QueryRaw(arg0 *database.Filters, arg1, arg2 int) ([]database.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRaw", arg0, arg1, arg2)
	ret0, _ := ret[0].([]database.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryRaw indicates an expected call of QueryRaw.
func (mr *MockObjectStoreMockRecorder) QueryRaw(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRaw", reflect.TypeOf((*MockObjectStore)(nil).QueryRaw), arg0, arg1, arg2)
}

// RemoveIDsFromFullTextQueue mocks base method.
func (m *MockObjectStore) RemoveIDsFromFullTextQueue(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveIDsFromFullTextQueue", arg0)
}

// RemoveIDsFromFullTextQueue indicates an expected call of RemoveIDsFromFullTextQueue.
func (mr *MockObjectStoreMockRecorder) RemoveIDsFromFullTextQueue(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIDsFromFullTextQueue", reflect.TypeOf((*MockObjectStore)(nil).RemoveIDsFromFullTextQueue), arg0)
}

// Run mocks base method.
func (m *MockObjectStore) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockObjectStoreMockRecorder) Run(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockObjectStore)(nil).Run), arg0)
}

// SaveAccountStatus mocks base method.
func (m *MockObjectStore) SaveAccountStatus(arg0 *coordinatorproto.SpaceStatusPayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAccountStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAccountStatus indicates an expected call of SaveAccountStatus.
func (mr *MockObjectStoreMockRecorder) SaveAccountStatus(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAccountStatus", reflect.TypeOf((*MockObjectStore)(nil).SaveAccountStatus), arg0)
}

// SaveChecksums mocks base method.
func (m *MockObjectStore) SaveChecksums(arg0 string, arg1 *model.ObjectStoreChecksums) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveChecksums", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveChecksums indicates an expected call of SaveChecksums.
func (mr *MockObjectStoreMockRecorder) SaveChecksums(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveChecksums", reflect.TypeOf((*MockObjectStore)(nil).SaveChecksums), arg0, arg1)
}

// SaveLastIndexedHeadsHash mocks base method.
func (m *MockObjectStore) SaveLastIndexedHeadsHash(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLastIndexedHeadsHash", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveLastIndexedHeadsHash indicates an expected call of SaveLastIndexedHeadsHash.
func (mr *MockObjectStoreMockRecorder) SaveLastIndexedHeadsHash(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLastIndexedHeadsHash", reflect.TypeOf((*MockObjectStore)(nil).SaveLastIndexedHeadsHash), arg0, arg1)
}

// SaveVirtualSpace mocks base method.
func (m *MockObjectStore) SaveVirtualSpace(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveVirtualSpace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveVirtualSpace indicates an expected call of SaveVirtualSpace.
func (mr *MockObjectStoreMockRecorder) SaveVirtualSpace(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveVirtualSpace", reflect.TypeOf((*MockObjectStore)(nil).SaveVirtualSpace), arg0)
}

// SubscribeForAll mocks base method.
func (m *MockObjectStore) SubscribeForAll(arg0 func(database.Record)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SubscribeForAll", arg0)
}

// SubscribeForAll indicates an expected call of SubscribeForAll.
func (mr *MockObjectStoreMockRecorder) SubscribeForAll(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeForAll", reflect.TypeOf((*MockObjectStore)(nil).SubscribeForAll), arg0)
}

// UpdateObjectDetails mocks base method.
func (m *MockObjectStore) UpdateObjectDetails(arg0 string, arg1 *types.Struct) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateObjectDetails", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateObjectDetails indicates an expected call of UpdateObjectDetails.
func (mr *MockObjectStoreMockRecorder) UpdateObjectDetails(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateObjectDetails", reflect.TypeOf((*MockObjectStore)(nil).UpdateObjectDetails), arg0, arg1)
}

// UpdateObjectLinks mocks base method.
func (m *MockObjectStore) UpdateObjectLinks(arg0 string, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateObjectLinks", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateObjectLinks indicates an expected call of UpdateObjectLinks.
func (mr *MockObjectStoreMockRecorder) UpdateObjectLinks(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateObjectLinks", reflect.TypeOf((*MockObjectStore)(nil).UpdateObjectLinks), arg0, arg1)
}

// UpdateObjectSnippet mocks base method.
func (m *MockObjectStore) UpdateObjectSnippet(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateObjectSnippet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateObjectSnippet indicates an expected call of UpdateObjectSnippet.
func (mr *MockObjectStoreMockRecorder) UpdateObjectSnippet(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateObjectSnippet", reflect.TypeOf((*MockObjectStore)(nil).UpdateObjectSnippet), arg0, arg1)
}

// UpdatePendingLocalDetails mocks base method.
func (m *MockObjectStore) UpdatePendingLocalDetails(arg0 string, arg1 func(*types.Struct) (*types.Struct, error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePendingLocalDetails", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePendingLocalDetails indicates an expected call of UpdatePendingLocalDetails.
func (mr *MockObjectStoreMockRecorder) UpdatePendingLocalDetails(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePendingLocalDetails", reflect.TypeOf((*MockObjectStore)(nil).UpdatePendingLocalDetails), arg0, arg1)
}
