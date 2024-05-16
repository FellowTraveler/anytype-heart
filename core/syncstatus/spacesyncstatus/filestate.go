package spacesyncstatus

import (
	"github.com/anyproto/anytype-heart/core/syncstatus/filesyncstatus"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/database"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

type FileState struct {
	fileSyncCountBySpace  map[string]int
	fileSyncStatusBySpace map[string]SpaceSyncStatus

	store objectstore.ObjectStore
}

func NewFileState(store objectstore.ObjectStore) *FileState {
	return &FileState{
		fileSyncCountBySpace:  make(map[string]int, 0),
		fileSyncStatusBySpace: make(map[string]SpaceSyncStatus, 0),

		store: store,
	}
}

func (f *FileState) SetObjectsNumber(status *SpaceSync) {
	records, err := f.store.Query(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				RelationKey: bundle.RelationKeyFileBackupStatus.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.Int64(int64(filesyncstatus.Syncing)),
			},
			{
				RelationKey: bundle.RelationKeySpaceId.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String(status.SpaceId),
			},
		},
	})
	if err != nil {
		log.Errorf("failed to query file status: %s", err)
	}
	f.fileSyncCountBySpace[status.SpaceId] = len(records)
}

func (f *FileState) SetSyncStatus(status *SpaceSync) {
	switch status.Status {
	case Synced:
		f.fileSyncStatusBySpace[status.SpaceId] = Synced
		if number := f.fileSyncCountBySpace[status.SpaceId]; number > 0 {
			f.fileSyncStatusBySpace[status.SpaceId] = Syncing
		}
	case Error, Syncing, Offline:
		f.fileSyncStatusBySpace[status.SpaceId] = status.Status
	}
}

func (f *FileState) GetSyncStatus(spaceId string) SpaceSyncStatus {
	return f.fileSyncStatusBySpace[spaceId]
}

func (f *FileState) GetSyncObjectCount(spaceId string) int {
	return f.fileSyncCountBySpace[spaceId]
}

func (f *FileState) IsSyncFinished(spaceId string) bool {
	if _, ok := f.fileSyncStatusBySpace[spaceId]; !ok {
		return false
	}
	status := f.fileSyncStatusBySpace[spaceId]
	count := f.fileSyncCountBySpace[spaceId]
	return count == 0 && status == Synced
}
