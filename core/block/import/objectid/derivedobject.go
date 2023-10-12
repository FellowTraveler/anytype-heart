package objectid

import (
	"context"
	"fmt"
	"time"

	"github.com/anyproto/any-sync/commonspace/object/tree/treestorage"

	"github.com/anyproto/anytype-heart/core/block/import/converter"
	"github.com/anyproto/anytype-heart/core/block/object/payloadcreator"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/space"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

type derivedObject struct {
	existingObject *existingObject
	spaceService   space.SpaceService
}

func newDerivedObject(existingObject *existingObject, spaceService space.SpaceService) *derivedObject {
	return &derivedObject{existingObject: existingObject, spaceService: spaceService}
}

func (r *derivedObject) GetIDAndPayload(ctx context.Context, spaceID string, sn *converter.Snapshot, createdTime time.Time, getExisting bool) (string, treestorage.TreeStorageCreatePayload, error) {
	id, payload, err := r.existingObject.GetIDAndPayload(ctx, spaceID, sn, createdTime, getExisting)
	if err != nil {
		return "", treestorage.TreeStorageCreatePayload{}, err
	}
	if id != "" {
		return id, payload, nil
	}
	id = pbtypes.GetString(sn.Snapshot.Data.Details, bundle.RelationKeyUniqueKey.String())
	uk, err := domain.UnmarshalUniqueKey(id)
	if err != nil {
		id = pbtypes.GetString(sn.Snapshot.Data.Details, bundle.RelationKeyId.String())
		uk, err = domain.UnmarshalUniqueKey(id)
		if err != nil {
			return "", treestorage.TreeStorageCreatePayload{}, fmt.Errorf("get unique key: %w", err)
		}
	}

	spc, err := r.spaceService.Get(ctx, spaceID)
	if err != nil {
		return "", treestorage.TreeStorageCreatePayload{}, fmt.Errorf("get space : %w", err)
	}
	payload, err = spc.DeriveTreePayload(ctx, payloadcreator.PayloadDerivationParams{
		Key: uk,
	})
	if err != nil {
		return "", treestorage.TreeStorageCreatePayload{}, fmt.Errorf("derive tree create payload: %w", err)
	}
	return payload.RootRawChange.Id, payload, nil
}
