package objectcache

import (
	"crypto/rand"
	"fmt"

	"github.com/anyproto/any-sync/commonspace/object/tree/objecttree"
	"github.com/anyproto/any-sync/util/crypto"

	"github.com/anyproto/anytype-heart/core/domain"
	coresb "github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space/spacecore"
)

const ChangeType = "anytype.object"

func createChangePayload(sbType coresb.SmartBlockType, key domain.UniqueKey, spaceID string) (data []byte, err error) {
	var keyStr string
	if key != nil {
		if key.SmartblockType() != sbType {
			return nil, fmt.Errorf("uniquekey smartblocktype mismatch")
		}
		keyStr = key.InternalKey()
	}
	payload := &model.ObjectChangePayload{SmartBlockType: model.SmartBlockType(sbType), Key: keyStr}
	if sbType == coresb.SmartBlockTypeSpaceObject {
		mdl := &model.SpaceObjectHeader{SpaceID: spaceID}
		marshalled, err := mdl.Marshal()
		if err != nil {
			return nil, err
		}
		payload.Data = marshalled
	}
	return payload.Marshal()
}

func derivePayload(spaceId string, changePayload []byte) objecttree.ObjectTreeCreatePayload {
	return objecttree.ObjectTreeCreatePayload{
		ChangeType:    spacecore.ChangeType,
		ChangePayload: changePayload,
		SpaceId:       spaceId,
		IsEncrypted:   true,
	}
}

func derivePersonalPayload(spaceId string, signKey crypto.PrivKey, changePayload []byte) objecttree.ObjectTreeCreatePayload {
	return objecttree.ObjectTreeCreatePayload{
		PrivKey:       signKey,
		ChangeType:    spacecore.ChangeType,
		ChangePayload: changePayload,
		SpaceId:       spaceId,
		IsEncrypted:   true,
	}
}

func createPayload(spaceId string, signKey crypto.PrivKey, changePayload []byte, timestamp int64) (objecttree.ObjectTreeCreatePayload, error) {
	seed := make([]byte, 32)
	if _, err := rand.Read(seed); err != nil {
		return objecttree.ObjectTreeCreatePayload{}, err
	}
	return objecttree.ObjectTreeCreatePayload{
		PrivKey:       signKey,
		ChangeType:    spacecore.ChangeType,
		ChangePayload: changePayload,
		SpaceId:       spaceId,
		IsEncrypted:   true,
		Timestamp:     timestamp,
		Seed:          seed,
	}, nil
}
