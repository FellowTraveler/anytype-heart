package uniquekey

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
)

const separator = "-"

var smartBlockTypeToKey = map[model.SmartBlockType]string{
	model.SmartBlockType_STType:      "ot",
	model.SmartBlockType_STRelation:  "rel",
	model.SmartBlockType_Workspace:   "ws",
	model.SmartBlockType_Home:        "home",
	model.SmartBlockType_Archive:     "archive",
	model.SmartBlockType_ProfilePage: "profile",
	model.SmartBlockType_Widget:      "widget",
}

type UniqueKey interface {
	Marshal() string
	SmartblockType() model.SmartBlockType
	InternalKey() string // underlying key, e.g. for "ot-page" it's "page"
}

type uniqueKey struct {
	sbt model.SmartBlockType
	key string
}

func New(sbt model.SmartBlockType, key string) (UniqueKey, error) {
	if _, exists := smartBlockTypeToKey[sbt]; !exists {
		return nil, fmt.Errorf("smartblocktype %s not supported", sbt.String())
	}
	return &uniqueKey{
		sbt: sbt,
		key: key,
	}, nil
}

func UnmarshalFromString(raw string) (UniqueKey, error) {
	parts := strings.Split(raw, separator)
	if raw == "" || len(parts) > 2 {
		return nil, errors.New("invalid key format")
	}

	// UniqueKey can be without second component, for example, unique key for Workspace object
	var key string
	if len(parts) == 2 {
		key = parts[1]
	}
	for sbt, sbtString := range smartBlockTypeToKey {
		if sbtString == parts[0] {
			return &uniqueKey{
				sbt: sbt,
				key: key,
			}, nil
		}
	}
	return nil, fmt.Errorf("smartblocktype %s not supported", parts[0])
}

func (uk *uniqueKey) Marshal() string {
	if uk.key == "" {
		return smartBlockTypeToKey[uk.sbt]
	}
	return smartBlockTypeToKey[uk.sbt] + separator + uk.key
}

func (uk *uniqueKey) SmartblockType() model.SmartBlockType {
	return uk.sbt
}

func (uk *uniqueKey) InternalKey() string {
	return uk.key
}
