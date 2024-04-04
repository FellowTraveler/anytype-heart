package spacecontroller

import (
	"context"

	"github.com/anyproto/anytype-heart/space/internal/spaceprocess/mode"
	"github.com/anyproto/anytype-heart/space/spaceinfo"
)

type SpaceController interface {
	SpaceId() string
	Start(ctx context.Context) error
	Mode() mode.Mode
	Current() any
	UpdateInfo(ctx context.Context, info spaceinfo.SpacePersistentInfo) error
	SetInfo(ctx context.Context, info spaceinfo.SpacePersistentInfo) error
	SetRemoteStatus(ctx context.Context, status spaceinfo.SpaceRemoteStatusInfo) error
	Close(ctx context.Context) error
	GetStatus() spaceinfo.AccountStatus
}

type DeleteController interface {
	Delete(ctx context.Context) error
}
