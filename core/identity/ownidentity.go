package identity

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/anyproto/any-sync/identityrepo/identityrepoproto"
	"github.com/anyproto/any-sync/util/crypto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"go.uber.org/zap"

	"github.com/anyproto/anytype-heart/core/anytype/account"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/core/files/fileacl"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	coresb "github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
	"github.com/anyproto/anytype-heart/pkg/lib/database"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

type observerService interface {
	broadcastMyIdentityProfile(identityProfile *model.IdentityProfile)
}

type ownProfileSubscription struct {
	spaceService       space.Service
	objectStore        objectstore.ObjectStore
	accountService     account.Service
	identityRepoClient identityRepoClient
	fileAclService     fileacl.Service
	observerService    observerService

	myIdentity          string
	globalNameUpdatedCh chan string
	gotDetailsCh        chan struct{}

	detailsLock sync.RWMutex
	gotDetails  bool
	details     *types.Struct // save details to batch update operation

	pushIdentityTimer        *time.Timer // timer for batching
	pushIdentityBatchTimeout time.Duration

	componentCtx       context.Context
	componentCtxCancel context.CancelFunc
}

func newOwnProfileSubscription(
	spaceService space.Service,
	objectStore objectstore.ObjectStore,
	accountService account.Service,
	identityRepoClient identityRepoClient,
	fileAclService fileacl.Service,
	observerService observerService,
	pushIdentityBatchTimeout time.Duration,
) *ownProfileSubscription {
	componentCtx, componentCtxCancel := context.WithCancel(context.Background())
	return &ownProfileSubscription{
		spaceService:             spaceService,
		objectStore:              objectStore,
		accountService:           accountService,
		identityRepoClient:       identityRepoClient,
		fileAclService:           fileAclService,
		observerService:          observerService,
		globalNameUpdatedCh:      make(chan string),
		gotDetailsCh:             make(chan struct{}),
		pushIdentityBatchTimeout: pushIdentityBatchTimeout,
		componentCtx:             componentCtx,
		componentCtxCancel:       componentCtxCancel,
	}
}

func (s *ownProfileSubscription) run(ctx context.Context) (err error) {
	s.myIdentity = s.accountService.AccountID()

	uniqueKey, err := domain.NewUniqueKey(coresb.SmartBlockTypeProfilePage, "")
	if err != nil {
		return err
	}
	personalSpace, err := s.spaceService.GetPersonalSpace(ctx)
	if err != nil {
		return fmt.Errorf("get space: %w", err)
	}
	profileObjectId, err := personalSpace.DeriveObjectID(ctx, uniqueKey)
	if err != nil {
		return err
	}

	recordsCh := make(chan *types.Struct)
	sub := database.NewSubscription(nil, recordsCh)

	var (
		records  []database.Record
		closeSub func()
	)

	records, closeSub, err = s.objectStore.QueryByIDAndSubscribeForChanges([]string{profileObjectId}, sub)
	if err != nil {
		return err
	}
	go func() {
		select {
		case <-s.componentCtx.Done():
			closeSub()
			return
		}
	}()

	if len(records) > 0 {
		s.handleOwnProfileDetails(records[0].Details)
	}

	go func() {
		for {
			select {
			case <-s.componentCtx.Done():
				return
			case rec, ok := <-recordsCh:
				if !ok {
					return
				}
				s.handleOwnProfileDetails(rec)

			case globalName := <-s.globalNameUpdatedCh:
				s.handleGlobalNameUpdate(globalName)
			}
		}
	}()

	return nil
}

func (s *ownProfileSubscription) close() {
	s.componentCtxCancel()
	close(s.globalNameUpdatedCh)
}

func (s *ownProfileSubscription) updateGlobalName(globalName string) {
	s.globalNameUpdatedCh <- globalName
}

func (s *ownProfileSubscription) enqueuePush() {
	if s.pushIdentityTimer == nil {
		s.pushIdentityTimer = time.AfterFunc(0, func() {
			pushErr := s.pushProfileToIdentityRegistry(context.Background())
			if pushErr != nil {
				log.Error("push profile to identity registry", zap.Error(pushErr))
			}
		})
	} else {
		s.pushIdentityTimer.Reset(s.pushIdentityBatchTimeout)
	}
}

func (s *ownProfileSubscription) handleOwnProfileDetails(profileDetails *types.Struct) {
	if profileDetails == nil {
		return
	}
	s.detailsLock.Lock()
	if !s.gotDetails {
		close(s.gotDetailsCh)
		s.gotDetails = true
	}

	if s.details == nil {
		s.details = &types.Struct{
			Fields: map[string]*types.Value{},
		}
	}
	for _, key := range []domain.RelationKey{
		bundle.RelationKeyId,
		bundle.RelationKeyName,
		bundle.RelationKeyDescription,
		bundle.RelationKeyIconImage,
	} {
		s.details.Fields[key.String()] = profileDetails.Fields[key.String()]
	}
	identityProfile := s.prepareIdentityProfile()
	s.detailsLock.Unlock()

	s.observerService.broadcastMyIdentityProfile(identityProfile)
	s.enqueuePush()
}

func (s *ownProfileSubscription) handleGlobalNameUpdate(globalName string) {
	s.detailsLock.Lock()
	if s.details == nil {
		s.details = &types.Struct{
			Fields: map[string]*types.Value{},
		}
	}
	s.details.Fields[bundle.RelationKeyGlobalName.String()] = pbtypes.String(globalName)
	identityProfile := s.prepareIdentityProfile()
	s.detailsLock.Unlock()

	s.observerService.broadcastMyIdentityProfile(identityProfile)

	s.enqueuePush()
}

func (s *ownProfileSubscription) prepareIdentityProfile() *model.IdentityProfile {
	return &model.IdentityProfile{
		Identity:    s.myIdentity,
		Name:        pbtypes.GetString(s.details, bundle.RelationKeyName.String()),
		Description: pbtypes.GetString(s.details, bundle.RelationKeyDescription.String()),
		IconCid:     pbtypes.GetString(s.details, bundle.RelationKeyIconImage.String()),
		GlobalName:  pbtypes.GetString(s.details, bundle.RelationKeyGlobalName.String()),
	}
}

func (s *ownProfileSubscription) pushProfileToIdentityRegistry(ctx context.Context) error {
	identityProfile, err := s.prepareOwnIdentityProfile()
	if err != nil {
		return fmt.Errorf("prepare own identity profile: %w", err)
	}
	data, err := proto.Marshal(identityProfile)
	if err != nil {
		return fmt.Errorf("marshal identity profile: %w", err)
	}

	symKey := s.spaceService.AccountMetadataSymKey()
	data, err = symKey.Encrypt(data)
	if err != nil {
		return fmt.Errorf("encrypt data: %w", err)
	}

	signature, err := s.accountService.SignData(data)
	if err != nil {
		return fmt.Errorf("failed to sign profile data: %w", err)
	}

	err = s.identityRepoClient.IdentityRepoPut(ctx, s.myIdentity, []*identityrepoproto.Data{
		{
			Kind:      "profile",
			Data:      data,
			Signature: signature,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to push identity: %w", err)
	}
	return nil
}

func (s *ownProfileSubscription) prepareOwnIdentityProfile() (*model.IdentityProfile, error) {
	s.detailsLock.RLock()
	defer s.detailsLock.RUnlock()

	iconImageObjectId := pbtypes.GetString(s.details, bundle.RelationKeyIconImage.String())
	iconCid, iconEncryptionKeys, err := s.prepareIconImageInfo(iconImageObjectId)
	if err != nil {
		return nil, fmt.Errorf("prepare icon image info: %w", err)
	}

	identity := s.accountService.AccountID()
	return &model.IdentityProfile{
		Identity:           identity,
		Name:               pbtypes.GetString(s.details, bundle.RelationKeyName.String()),
		Description:        pbtypes.GetString(s.details, bundle.RelationKeyDescription.String()),
		IconCid:            iconCid,
		IconEncryptionKeys: iconEncryptionKeys,
		GlobalName:         pbtypes.GetString(s.details, bundle.RelationKeyGlobalName.String()),
	}, nil
}

func (s *ownProfileSubscription) prepareIconImageInfo(iconImageObjectId string) (iconCid string, iconEncryptionKeys []*model.FileEncryptionKey, err error) {
	if iconImageObjectId == "" {
		return "", nil, nil
	}
	return s.fileAclService.GetInfoForFileSharing(iconImageObjectId)
}

func (s *ownProfileSubscription) getDetails(ctx context.Context) (identity string, metadataKey crypto.SymKey, details *types.Struct) {
	select {
	case <-s.gotDetailsCh:

	case <-ctx.Done():
		return "", nil, nil
	case <-s.componentCtx.Done():
		return "", nil, nil
	}
	s.detailsLock.RLock()
	defer s.detailsLock.RUnlock()

	return s.myIdentity, s.spaceService.AccountMetadataSymKey(), s.details
}
