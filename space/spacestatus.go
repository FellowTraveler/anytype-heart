package space

import (
	"errors"
	"github.com/anytypeio/any-sync/coordinator/coordinatorproto"
	"github.com/anytypeio/any-sync/net/rpc/rpcerr"
	"time"
)

type SpaceStatus int32

var (
	ErrSpaceDeleteUnexpected = errors.New("unexpected error while deleting space")
	ErrSpaceIsDeleted        = errors.New("space is deleted")
	ErrSpaceIsCreated        = errors.New("space is created")
	ErrSpaceDeletionPending  = errors.New("space deletion is pending")
)

const (
	SpaceStatusCreated SpaceStatus = iota
	SpaceStatusPendingDeletion
	SpaceStatusDeletionStarted
	SpaceStatusDeleted
)

type SpaceStatusPayload struct {
	Status       SpaceStatus
	DeletionDate time.Time
}

func newSpaceStatus(payload *coordinatorproto.SpaceStatusPayload) SpaceStatusPayload {
	return SpaceStatusPayload{
		Status:       SpaceStatus(payload.Status),
		DeletionDate: time.Unix(0, payload.DeletionTimestamp),
	}
}

func coordError(err error) error {
	err = rpcerr.Unwrap(err)
	switch err {
	case coordinatorproto.ErrSpaceDeletionPending:
		return ErrSpaceDeletionPending
	case coordinatorproto.ErrSpaceIsDeleted:
		return ErrSpaceIsDeleted
	case coordinatorproto.ErrSpaceIsCreated:
		return ErrSpaceIsCreated
	default:
		return ErrSpaceDeleteUnexpected
	}
}
