package application

import (
	"context"
	"errors"

	"github.com/anyproto/anytype-heart/core/configfetcher"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space"
)

var (
	ErrAccountIsAlreadyDeleted = errors.New("account is already deleted")
	ErrAccountIsActive         = errors.New("account is active")
)

func (s *Service) AccountDelete(ctx context.Context, req *pb.RpcAccountDeleteRequest) (*model.AccountStatus, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	spaceService := s.app.MustComponent(space.CName).(space.Service)
	resp, err := spaceService.DeleteAccount(ctx, req.Revert)
	status := &model.AccountStatus{
		StatusType:   model.AccountStatusType(resp.Status),
		DeletionDate: resp.DeletionDate.Unix(),
	}

	// so we will receive updated account status
	s.refreshRemoteAccountState()

	switch err {
	case space.ErrSpaceIsDeleted:
		return nil, ErrAccountIsAlreadyDeleted
	case space.ErrSpaceDeletionPending:
		return nil, ErrAccountIsAlreadyDeleted
	case space.ErrSpaceIsCreated:
		return nil, ErrAccountIsActive
	}
	return status, err
}

func (s *Service) refreshRemoteAccountState() {
	fetcher := s.app.MustComponent(configfetcher.CName).(configfetcher.ConfigFetcher)
	fetcher.Refetch()
}
