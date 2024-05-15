package syncstatus

import (
	"context"
	"fmt"

	"github.com/anyproto/anytype-heart/core/block/cache"
	"github.com/anyproto/anytype-heart/core/block/editor/basic"
	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/syncstatus/filesyncstatus"
	"github.com/anyproto/anytype-heart/core/syncstatus/helpers"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

func (s *service) OnFileUploadStarted(objectId string) error {
	return s.indexFileSyncStatus(objectId, filesyncstatus.Syncing)
}

func (s *service) OnFileUploaded(objectId string) error {
	return s.indexFileSyncStatus(objectId, filesyncstatus.Synced)
}

func (s *service) OnFileLimited(objectId string) error {
	return s.indexFileSyncStatus(objectId, filesyncstatus.Limited)
}

func (s *service) indexFileSyncStatus(fileObjectId string, status filesyncstatus.Status) error {
	var spaceId string
	err := cache.Do(s.objectGetter, fileObjectId, func(sb smartblock.SmartBlock) (err error) {
		spaceId = sb.SpaceID()
		prevStatus := pbtypes.GetInt64(sb.Details(), bundle.RelationKeyFileBackupStatus.String())
		newStatus := int64(status)
		if prevStatus == newStatus {
			return nil
		}
		detailsSetter, ok := sb.(basic.DetailsSettable)
		if !ok {
			return fmt.Errorf("setting of details is not supported for %T", sb)
		}
		return detailsSetter.SetDetails(nil, []*model.Detail{
			{
				Key:   bundle.RelationKeyFileBackupStatus.String(),
				Value: pbtypes.Int64(newStatus),
			},
		}, true)
	})
	if err != nil {
		return fmt.Errorf("get object: %w", err)
	}

	err = s.updateReceiver.UpdateTree(context.Background(), fileObjectId, status.ToSyncStatus())
	if err != nil {
		return fmt.Errorf("update tree: %w", err)
	}

	s.sendSpaceStatusUpdate(status, spaceId)
	return nil
}

func (s *service) sendSpaceStatusUpdate(status filesyncstatus.Status, spaceId string) {
	var (
		spaceStatus helpers.SpaceSyncStatus
		spaceError  helpers.SpaceSyncError
	)
	switch status {
	case filesyncstatus.Synced:
		spaceStatus = helpers.Synced
	case filesyncstatus.Syncing:
		spaceStatus = helpers.Syncing
	case filesyncstatus.Limited:
		spaceStatus = helpers.Error
		spaceError = helpers.StorageLimitExceed
	case filesyncstatus.Unknown:
		spaceStatus = helpers.Error
		spaceError = helpers.NetworkError
	}

	syncStatus := helpers.MakeSyncStatus(spaceId, spaceStatus, 0, spaceError, helpers.Files)
	s.spaceSyncStatus.SendUpdate(syncStatus)
}
