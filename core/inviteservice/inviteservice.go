package inviteservice

import (
	"context"
	"fmt"

	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/app/logger"
	"github.com/anyproto/any-sync/util/crypto"
	"github.com/gogo/protobuf/proto"
	"github.com/ipfs/go-cid"
	"go.uber.org/zap"

	"github.com/anyproto/anytype-heart/core/anytype/account"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/core/files/fileacl"
	"github.com/anyproto/anytype-heart/core/invitestore"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space"
	"github.com/anyproto/anytype-heart/space/spaceinfo"
	"github.com/anyproto/anytype-heart/space/techspace"
	"github.com/anyproto/anytype-heart/util/encode"
)

type InviteInfo struct {
	InviteFileCid string
	InviteFileKey string
}

const CName = "common.core.inviteservice"

var log = logger.NewNamed(CName)

type InviteView struct {
	SpaceId      string
	SpaceName    string
	SpaceIconCid string
	CreatorName  string
}

type InviteService interface {
	app.ComponentRunnable
	GetPayload(ctx context.Context, inviteCid cid.Cid, inviteFileKey crypto.SymKey) (*model.InvitePayload, error)
	View(ctx context.Context, inviteCid cid.Cid, inviteFileKey crypto.SymKey) (InviteView, error)
	RemoveExisting(ctx context.Context, spaceId string) error
	Generate(ctx context.Context, spaceId string, inviteKey crypto.PrivKey, sendInvite func() error) (InviteInfo, error)
	GetCurrent(ctx context.Context, spaceId string) (InviteInfo, error)
}

var _ InviteService = (*inviteService)(nil)

type inviteService struct {
	inviteStore    invitestore.Service
	fileAcl        fileacl.Service
	accountService account.Service
	spaceService   space.Service
}

func New() InviteService {
	return &inviteService{}
}

func (i *inviteService) Init(a *app.App) (err error) {
	i.inviteStore = app.MustComponent[invitestore.Service](a)
	i.fileAcl = app.MustComponent[fileacl.Service](a)
	i.accountService = app.MustComponent[account.Service](a)
	i.spaceService = app.MustComponent[space.Service](a)
	return
}

func (i *inviteService) Name() (name string) {
	return CName
}

func (i *inviteService) Run(ctx context.Context) (err error) {
	return
}

func (i *inviteService) Close(ctx context.Context) (err error) {
	return
}

func (i *inviteService) View(ctx context.Context, inviteCid cid.Cid, inviteFileKey crypto.SymKey) (InviteView, error) {
	invitePayload, err := i.GetPayload(ctx, inviteCid, inviteFileKey)
	if err != nil {
		return InviteView{}, err
	}
	return InviteView{
		SpaceId:      invitePayload.SpaceId,
		SpaceName:    invitePayload.SpaceName,
		SpaceIconCid: invitePayload.SpaceIconCid,
		CreatorName:  invitePayload.CreatorName,
	}, nil
}

func (i *inviteService) GetCurrent(ctx context.Context, spaceId string) (info InviteInfo, err error) {
	var (
		fileCid, fileKey string
	)
	err = i.spaceService.TechSpace().DoSpaceView(ctx, spaceId, func(spaceView techspace.SpaceView) error {
		fileCid, fileKey = spaceView.GetExistingInviteInfo()
		return nil
	})
	if err != nil {
		err = getInviteError("get existing invite info", err)
		return
	}
	if fileCid == "" {
		err = ErrInviteNotExists
		return
	}
	info.InviteFileCid = fileCid
	info.InviteFileKey = fileKey
	return
}

func (i *inviteService) RemoveExisting(ctx context.Context, spaceId string) (err error) {
	var fileCid string
	err = i.spaceService.TechSpace().DoSpaceView(ctx, spaceId, func(spaceView techspace.SpaceView) error {
		fileCid, err = spaceView.RemoveExistingInviteInfo()
		return err
	})
	if err != nil {
		return removeInviteError("remove existing invite info", err)
	}
	if len(fileCid) == 0 {
		return nil
	}
	invCid, err := cid.Decode(fileCid)
	if err != nil {
		return removeInviteError("decode invite cid", err)
	}
	err = i.inviteStore.RemoveInvite(ctx, invCid)
	if err != nil {
		return removeInviteError("remove invite from store", err)
	}
	return
}

func (i *inviteService) Generate(ctx context.Context, spaceId string, inviteKey crypto.PrivKey, sendInvite func() error) (result InviteInfo, err error) {
	if spaceId == i.accountService.PersonalSpaceID() {
		return InviteInfo{}, ErrPersonalSpace
	}
	var fileCid, fileKey string
	err = i.spaceService.TechSpace().DoSpaceView(ctx, spaceId, func(spaceView techspace.SpaceView) error {
		fileCid, fileKey = spaceView.GetExistingInviteInfo()
		return nil
	})
	if err != nil {
		return InviteInfo{}, generateInviteError("get existing invite info", err)
	}
	if fileCid != "" {
		return InviteInfo{
			InviteFileCid: fileCid,
			InviteFileKey: fileKey,
		}, nil
	}
	invite, err := i.buildInvite(ctx, spaceId, inviteKey)
	if err != nil {
		return InviteInfo{}, generateInviteError("build invite", err)
	}
	inviteFileCid, inviteFileKey, err := i.inviteStore.StoreInvite(ctx, invite)
	if err != nil {
		return InviteInfo{}, generateInviteError("store invite in ipfs", err)
	}
	removeInviteFile := func() {
		err := i.inviteStore.RemoveInvite(ctx, inviteFileCid)
		if err != nil {
			log.Error("remove invite file", zap.Error(err))
		}
	}
	inviteFileKeyRaw, err := encode.EncodeKeyToBase58(inviteFileKey)
	if err != nil {
		removeInviteFile()
		return InviteInfo{}, generateInviteError("encode invite file key", err)
	}
	err = i.spaceService.TechSpace().DoSpaceView(ctx, spaceId, func(spaceView techspace.SpaceView) error {
		return spaceView.SetInviteFileInfo(inviteFileCid.String(), inviteFileKeyRaw)
	})
	if err != nil {
		removeInviteFile()
		return InviteInfo{}, generateInviteError("set invite file info", err)
	}
	err = sendInvite()
	if err != nil {
		removeErr := i.RemoveExisting(ctx, spaceId)
		if removeErr != nil {
			log.Error("remove existing invite", zap.Error(removeErr))
		}
		return InviteInfo{}, generateInviteError("send invite", err)
	}
	return InviteInfo{
		InviteFileCid: inviteFileCid.String(),
		InviteFileKey: inviteFileKeyRaw,
	}, err
}

func (i *inviteService) GetPayload(ctx context.Context, inviteCid cid.Cid, inviteFileKey crypto.SymKey) (md *model.InvitePayload, err error) {
	invite, err := i.inviteStore.GetInvite(ctx, inviteCid, inviteFileKey)
	if err != nil {
		return nil, getInviteError("get invite from store", err)
	}
	var invitePayload model.InvitePayload
	err = proto.Unmarshal(invite.Payload, &invitePayload)
	if err != nil {
		return nil, badContentError("unmarshal invite payload", err)
	}
	creatorIdentity, err := crypto.DecodeAccountAddress(invitePayload.CreatorIdentity)
	if err != nil {
		return nil, badContentError("decode creator identity", err)
	}
	ok, err := creatorIdentity.Verify(invite.Payload, invite.Signature)
	if err != nil {
		return nil, badContentError("verify creator identity", err)
	}
	if !ok {
		return nil, badContentError("verify creator identity", fmt.Errorf("signature is invalid"))
	}
	err = i.fileAcl.StoreFileKeys(domain.FileId(invitePayload.SpaceIconCid), invitePayload.SpaceIconEncryptionKeys)
	if err != nil {
		return nil, getInviteError("store space icon encryption keys", err)
	}
	return &invitePayload, nil
}

func (i *inviteService) buildInvite(ctx context.Context, spaceId string, inviteKey crypto.PrivKey) (*model.Invite, error) {
	invitePayload, err := i.buildInvitePayload(ctx, spaceId, inviteKey)
	if err != nil {
		return nil, fmt.Errorf("build invite payload: %w", err)
	}
	invitePayloadRaw, err := proto.Marshal(invitePayload)
	if err != nil {
		return nil, fmt.Errorf("marshal invite payload: %w", err)
	}
	invitePayloadSignature, err := i.accountService.SignData(invitePayloadRaw)
	if err != nil {
		return nil, fmt.Errorf("sign invite payload: %w", err)
	}
	return &model.Invite{
		Payload:   invitePayloadRaw,
		Signature: invitePayloadSignature,
	}, nil
}

func (i *inviteService) buildInvitePayload(ctx context.Context, spaceId string, inviteKey crypto.PrivKey) (*model.InvitePayload, error) {
	profile, err := i.accountService.ProfileInfo()
	if err != nil {
		return nil, fmt.Errorf("get profile info: %w", err)
	}
	rawInviteKey, err := inviteKey.Marshall()
	if err != nil {
		return nil, fmt.Errorf("marshal invite priv key: %w", err)
	}
	invitePayload := &model.InvitePayload{
		SpaceId:         spaceId,
		CreatorIdentity: i.accountService.AccountID(),
		CreatorName:     profile.Name,
		InviteKey:       rawInviteKey,
	}
	var description spaceinfo.SpaceDescription
	err = i.spaceService.TechSpace().DoSpaceView(ctx, spaceId, func(spaceView techspace.SpaceView) error {
		description = spaceView.GetSpaceDescription()
		return nil
	})
	invitePayload.SpaceName = description.Name
	if err != nil {
		return nil, fmt.Errorf("get space description: %w", err)
	}
	if description.IconImage != "" {
		iconCid, iconEncryptionKeys, err := i.fileAcl.GetInfoForFileSharing(ctx, description.IconImage)
		if err == nil {
			invitePayload.SpaceIconCid = iconCid
			invitePayload.SpaceIconEncryptionKeys = iconEncryptionKeys
		} else {
			log.Error("get space icon info", zap.Error(err))
		}
	}
	return invitePayload, nil
}
