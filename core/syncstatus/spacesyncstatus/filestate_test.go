package spacesyncstatus

import (
	"testing"

	"github.com/anyproto/any-sync/commonspace/syncstatus"
	"github.com/stretchr/testify/assert"

	"github.com/anyproto/anytype-heart/core/syncstatus/filesyncstatus"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

func TestFileState_GetSyncObjectCount(t *testing.T) {
	t.Run("GetSyncObjectCount", func(t *testing.T) {
		// given
		fileState := NewFileState(nil)

		// when
		fileState.fileSyncCountBySpace["spaceId"] = 1
		objectCount := fileState.GetSyncObjectCount("spaceId")

		// then
		assert.Equal(t, 1, objectCount)
	})
	t.Run("GetSyncObjectCount: zero value", func(t *testing.T) {
		// given
		fileState := NewFileState(nil)

		// when
		objectCount := fileState.GetSyncObjectCount("spaceId")

		// then
		assert.Equal(t, 0, objectCount)
	})
}

func TestFileState_GetSyncStatus(t *testing.T) {
	t.Run("GetSyncStatus", func(t *testing.T) {
		// given
		fileState := NewFileState(nil)

		// when
		fileState.fileSyncStatusBySpace["spaceId"] = syncstatus.Syncing
		syncStatus := fileState.GetSyncStatus("spaceId")

		// then
		assert.Equal(t, syncstatus.Syncing, syncStatus)
	})
	t.Run("GetSyncStatus: zero value", func(t *testing.T) {
		// given
		fileState := NewFileState(nil)

		// when
		syncStatus := fileState.GetSyncStatus("spaceId")

		// then
		assert.Equal(t, syncstatus.Synced, syncStatus)
	})
}

func TestFileState_SetObjectsNumber(t *testing.T) {
	t.Run("SetObjectsNumber", func(t *testing.T) {
		// given
		storeFixture := objectstore.NewStoreFixture(t)
		storeFixture.AddObjects(t, []objectstore.TestObject{
			{
				bundle.RelationKeyId:               pbtypes.String("id1"),
				bundle.RelationKeyFileBackupStatus: pbtypes.Int64(int64(filesyncstatus.Syncing)),
				bundle.RelationKeySpaceId:          pbtypes.String("spaceId"),
			},
			{
				bundle.RelationKeyId:               pbtypes.String("id2"),
				bundle.RelationKeyFileBackupStatus: pbtypes.Int64(int64(filesyncstatus.Synced)),
				bundle.RelationKeySpaceId:          pbtypes.String("spaceId"),
			},
			{
				bundle.RelationKeyId:               pbtypes.String("id3"),
				bundle.RelationKeyFileBackupStatus: pbtypes.Int64(int64(filesyncstatus.Syncing)),
				bundle.RelationKeySpaceId:          pbtypes.String("spaceId"),
			},
		})
		fileState := NewFileState(storeFixture)
		syncStatus := syncstatus.MakeSyncStatus("spaceId", syncstatus.Synced, 0, syncstatus.Null, syncstatus.Files)

		// when
		fileState.SetObjectsNumber(syncStatus)

		// then
		assert.Equal(t, 2, fileState.GetSyncObjectCount("spaceId"))
	})
	t.Run("SetObjectsNumber: no file object", func(t *testing.T) {
		// given
		storeFixture := objectstore.NewStoreFixture(t)
		fileState := NewFileState(storeFixture)
		syncStatus := syncstatus.MakeSyncStatus("spaceId", syncstatus.Synced, 0, syncstatus.Null, syncstatus.Files)

		// when
		fileState.SetObjectsNumber(syncStatus)

		// then
		assert.Equal(t, 0, fileState.GetSyncObjectCount("spaceId"))
	})
}

func TestFileState_IsSyncFinished(t *testing.T) {
	t.Run("IsSyncFinished, sync is finished", func(t *testing.T) {
		// given
		fileState := NewFileState(nil)

		// when
		finished := fileState.IsSyncFinished("spaceId")

		// then
		assert.True(t, finished)
	})
	t.Run("IsSyncFinished, sync is finished", func(t *testing.T) {
		// given
		fileState := NewFileState(nil)

		// when
		syncStatus := syncstatus.MakeSyncStatus("spaceId", syncstatus.Synced, 0, syncstatus.Null, syncstatus.Files)
		fileState.SetSyncStatus(syncStatus)
		finished := fileState.IsSyncFinished("spaceId")

		// then
		assert.True(t, finished)
	})
	t.Run("IsSyncFinished, sync is not finished", func(t *testing.T) {
		// given
		fileState := NewFileState(nil)

		// when
		syncStatus := syncstatus.MakeSyncStatus("spaceId", syncstatus.Offline, 3, syncstatus.Null, syncstatus.Files)
		fileState.SetSyncStatus(syncStatus)
		finished := fileState.IsSyncFinished("spaceId")

		// then
		assert.False(t, finished)
	})
}

func TestFileState_SetSyncStatus(t *testing.T) {
	t.Run("SetSyncStatus, status synced", func(t *testing.T) {
		// given
		fileState := NewFileState(nil)

		// when
		syncStatus := syncstatus.MakeSyncStatus("spaceId", syncstatus.Synced, 0, syncstatus.Null, syncstatus.Files)
		fileState.SetSyncStatus(syncStatus)

		// then
		assert.Equal(t, syncstatus.Synced, fileState.GetSyncStatus("spaceId"))
	})
	t.Run("SetSyncStatus, received status synced, but there are syncing files in store", func(t *testing.T) {
		// given
		storeFixture := objectstore.NewStoreFixture(t)
		storeFixture.AddObjects(t, []objectstore.TestObject{
			{
				bundle.RelationKeyId:               pbtypes.String("id1"),
				bundle.RelationKeyFileBackupStatus: pbtypes.Int64(int64(filesyncstatus.Syncing)),
				bundle.RelationKeySpaceId:          pbtypes.String("spaceId"),
			},
			{
				bundle.RelationKeyId:               pbtypes.String("id2"),
				bundle.RelationKeyFileBackupStatus: pbtypes.Int64(int64(filesyncstatus.Synced)),
				bundle.RelationKeySpaceId:          pbtypes.String("spaceId"),
			},
			{
				bundle.RelationKeyId:               pbtypes.String("id3"),
				bundle.RelationKeyFileBackupStatus: pbtypes.Int64(int64(filesyncstatus.Syncing)),
				bundle.RelationKeySpaceId:          pbtypes.String("spaceId"),
			},
		})
		fileState := NewFileState(storeFixture)

		// when
		syncStatus := syncstatus.MakeSyncStatus("spaceId", syncstatus.Synced, 0, syncstatus.Null, syncstatus.Files)
		fileState.SetObjectsNumber(syncStatus)
		fileState.SetSyncStatus(syncStatus)

		// then
		assert.Equal(t, syncstatus.Syncing, fileState.GetSyncStatus("spaceId"))
	})
	t.Run("SetSyncStatus, sync in progress", func(t *testing.T) {
		// given
		fileState := NewFileState(nil)

		// when
		syncStatus := syncstatus.MakeSyncStatus("spaceId", syncstatus.Syncing, 0, syncstatus.Null, syncstatus.Files)
		fileState.SetSyncStatus(syncStatus)

		// then
		assert.Equal(t, syncstatus.Syncing, fileState.GetSyncStatus("spaceId"))
	})
	t.Run("SetSyncStatus, sync is finished with error", func(t *testing.T) {
		// given
		fileState := NewFileState(nil)

		// when
		syncStatus := syncstatus.MakeSyncStatus("spaceId", syncstatus.Error, 3, syncstatus.Null, syncstatus.Files)
		fileState.SetSyncStatus(syncStatus)

		// then
		assert.Equal(t, syncstatus.Error, fileState.GetSyncStatus("spaceId"))
	})
	t.Run("SetSyncStatus, offline", func(t *testing.T) {
		// given
		fileState := NewFileState(nil)

		// when
		syncStatus := syncstatus.MakeSyncStatus("spaceId", syncstatus.Offline, 3, syncstatus.Null, syncstatus.Files)
		fileState.SetSyncStatus(syncStatus)

		// then
		assert.Equal(t, syncstatus.Offline, fileState.GetSyncStatus("spaceId"))
	})
}
