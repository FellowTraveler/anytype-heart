package space

import (
	"context"
	"errors"
	"fmt"

	spaceservice "github.com/anyproto/anytype-heart/space/spacecore"
	"github.com/anyproto/anytype-heart/space/spaceinfo"
)

func (s *service) startLoad(ctx context.Context, spaceID string) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	persistentStatus := s.getPersistentStatus(spaceID)
	if persistentStatus.AccountStatus == spaceinfo.AccountStatusDeleted {
		return ErrSpaceDeleted
	}
	localStatus := s.getLocalStatus(spaceID)
	// Do nothing if space is already loading
	if localStatus.LocalStatus != spaceinfo.LocalStatusUnknown {
		return nil
	}

	exists, err := s.techSpace.SpaceViewExists(ctx, spaceID)
	if err != nil {
		return
	}
	if !exists {
		return ErrSpaceNotExists
	}

	info := spaceinfo.SpaceLocalInfo{
		SpaceID:     spaceID,
		LocalStatus: spaceinfo.LocalStatusLoading,
	}
	if err = s.setLocalStatus(ctx, info); err != nil {
		return
	}
	s.loading[spaceID] = s.newLoadingSpace(s.ctx, spaceID)
	return
}

func (s *service) onLoad(spaceID string, sp Space, loadErr error) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch {
	case loadErr == nil:
	case errors.Is(loadErr, spaceservice.ErrSpaceDeletionPending):
		return s.setLocalStatus(s.ctx, spaceinfo.SpaceLocalInfo{
			SpaceID:      spaceID,
			LocalStatus:  spaceinfo.LocalStatusMissing,
			RemoteStatus: spaceinfo.RemoteStatusWaitingDeletion,
		})
	case errors.Is(loadErr, spaceservice.ErrSpaceIsDeleted):
		return s.setLocalStatus(s.ctx, spaceinfo.SpaceLocalInfo{
			SpaceID:      spaceID,
			LocalStatus:  spaceinfo.LocalStatusMissing,
			RemoteStatus: spaceinfo.RemoteStatusDeleted,
		})
	default:
		return s.setLocalStatus(s.ctx, spaceinfo.SpaceLocalInfo{
			SpaceID:      spaceID,
			LocalStatus:  spaceinfo.LocalStatusMissing,
			RemoteStatus: spaceinfo.RemoteStatusError,
		})
	}

	s.loaded[spaceID] = sp
	delete(s.loading, spaceID)

	// TODO: check remote status
	return s.setLocalStatus(s.ctx, spaceinfo.SpaceLocalInfo{
		SpaceID:      spaceID,
		LocalStatus:  spaceinfo.LocalStatusOk,
		RemoteStatus: spaceinfo.RemoteStatusUnknown,
	})
}

func (s *service) preLoad(spc Space) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.loaded[spc.Id()] = spc
	s.localStatuses[spc.Id()] = spaceinfo.SpaceLocalInfo{
		SpaceID:      spc.Id(),
		LocalStatus:  spaceinfo.LocalStatusOk,
		RemoteStatus: spaceinfo.RemoteStatusUnknown,
	}
	s.persistentStatuses[spc.Id()] = spaceinfo.SpacePersistentInfo{
		SpaceID:       spc.Id(),
		AccountStatus: spaceinfo.AccountStatusUnknown,
	}
}

func (s *service) waitLoad(ctx context.Context, spaceID string) (sp Space, err error) {
	s.mu.Lock()
	status := s.getLocalStatus(spaceID)

	switch status.LocalStatus {
	case spaceinfo.LocalStatusUnknown:
		return nil, fmt.Errorf("waitLoad for an unknown space")
	case spaceinfo.LocalStatusLoading:
		// loading in progress, wait channel and retry
		waitCh := s.loading[spaceID].loadCh
		loadErr := s.loading[spaceID].loadErr
		s.mu.Unlock()
		if loadErr != nil {
			return nil, loadErr
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-waitCh:
		}
		return s.waitLoad(ctx, spaceID)
	case spaceinfo.LocalStatusMissing:
		// local missing status means the loader ended with an error
		err = s.loading[spaceID].loadErr
	case spaceinfo.LocalStatusOk:
		sp = s.loaded[spaceID]
	default:
		err = fmt.Errorf("undefined space status: %v", status.LocalStatus)
	}
	s.mu.Unlock()
	return
}
