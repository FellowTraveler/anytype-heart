package aclnotifications

import (
	"errors"
	"time"

	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/commonspace/object/acl/aclrecordproto"
	"github.com/anyproto/any-sync/commonspace/object/acl/list"
	"github.com/anyproto/any-sync/commonspace/object/acl/syncacl"
	"github.com/anyproto/any-sync/util/crypto"
	"golang.org/x/net/context"

	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space/clientspace"
	"github.com/anyproto/anytype-heart/space/internal/components/dependencies"
)

const Cname = "spaceNotification"

type NotificationSender interface {
	CreateAndSend(notification *model.Notification) error
	GetLastNotificationId(acl string) string
}

type AclNotification interface {
	SendNotification(ctx context.Context, space clientspace.Space, permissions list.AclPermissions, acl syncacl.SyncAcl) error
}

type AclNotificationSender struct {
	app.Component
	identityService     dependencies.IdentityService
	notificationService NotificationSender
}

func NewAclNotificationSender() *AclNotificationSender {
	return &AclNotificationSender{}
}

func (n *AclNotificationSender) Init(a *app.App) (err error) {
	n.identityService = app.MustComponent[dependencies.IdentityService](a)
	n.notificationService = app.MustComponent[NotificationSender](a)
	return nil
}

func (n *AclNotificationSender) Name() (name string) {
	return Cname
}

func (n *AclNotificationSender) SendNotification(ctx context.Context,
	space clientspace.Space,
	permissions list.AclPermissions,
	acl syncacl.SyncAcl,
) error {
	lastNotificationId := n.notificationService.GetLastNotificationId(acl.Id())
	var err error
	if lastNotificationId != "" {
		acl.IterateFrom(lastNotificationId, func(record *list.AclRecord) (IsContinue bool) {
			if err = n.handleAclContent(ctx, record, permissions, space, acl.Id()); err != nil {
				return false
			}
			return true
		})
		return err
	}
	for _, record := range acl.Records() {
		if err = n.handleAclContent(ctx, record, permissions, space, acl.Id()); err != nil {
			return err
		}
	}
	return nil
}

func (n *AclNotificationSender) handleAclContent(ctx context.Context,
	record *list.AclRecord,
	permissions list.AclPermissions,
	space clientspace.Space,
	aclId string,
) error {
	if aclData, ok := record.Model.(*aclrecordproto.AclData); ok {
		err := n.iterateAclContent(ctx, record, permissions, space, aclId, aclData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *AclNotificationSender) iterateAclContent(ctx context.Context,
	record *list.AclRecord,
	permissions list.AclPermissions,
	space clientspace.Space,
	aclId string,
	aclData *aclrecordproto.AclData,
) error {
	for _, content := range aclData.AclContent {
		if permissions.CanManageAccounts() {
			if reqJoin := content.GetRequestJoin(); reqJoin != nil {
				if err := n.sendJoinRequest(ctx, reqJoin, space, record.Id, aclId); err != nil {
					return err
				}
			}
		}
		if reqApprove := content.GetRequestAccept(); reqApprove != nil {
			if err := n.sendParticipantRequestApprove(reqApprove, space, record.Id, aclId); err != nil {
				return err
			}
		}
	}
	return nil
}

func (n *AclNotificationSender) sendJoinRequest(ctx context.Context,
	reqJoin *aclrecordproto.AclAccountRequestJoin,
	space clientspace.Space, id, aclId string,
) error {
	pubKey, name, icon, err := n.getProfileData(ctx, reqJoin)
	if err != nil {
		return err
	}
	err = n.notificationService.CreateAndSend(&model.Notification{
		Id:      id,
		IsLocal: false,
		Payload: &model.NotificationPayloadOfRequestToJoin{RequestToJoin: &model.NotificationRequestToJoin{
			SpaceId:      space.Id(),
			Identity:     pubKey.Account(),
			IdentityName: name,
			IdentityIcon: icon,
		}},
		Space: space.Id(),
		Acl:   aclId,
	})
	if err != nil {
		return err
	}
	return nil
}

func (n *AclNotificationSender) sendParticipantRequestApprove(reqApprove *aclrecordproto.AclAccountRequestAccept,
	space clientspace.Space,
	id, aclId string,
) error {
	identity, _, _ := n.identityService.GetMyProfileDetails()
	pubKey, err := crypto.UnmarshalEd25519PublicKeyProto(reqApprove.Identity)
	if err != nil {
		return err
	}
	account := pubKey.Account()
	if account != identity {
		return nil
	}
	err = n.notificationService.CreateAndSend(&model.Notification{
		Id:      id,
		IsLocal: false,
		Payload: &model.NotificationPayloadOfParticipantRequestApproved{
			ParticipantRequestApproved: &model.NotificationParticipantRequestApproved{
				SpaceID:    space.Id(),
				Permission: mapProtoPermissionToAcl(reqApprove.Permissions),
			},
		},
		Space: space.Id(),
		Acl:   aclId,
	})
	if err != nil {
		return err
	}
	return nil
}

func (n *AclNotificationSender) getProfileData(ctx context.Context, reqJoin *aclrecordproto.AclAccountRequestJoin,
) (crypto.PubKey, string, string, error) {
	pubKey, err := crypto.UnmarshalEd25519PublicKeyProto(reqJoin.InviteIdentity)
	if err != nil {
		return nil, "", "", err
	}
	ctxWithTimeout, _ := context.WithTimeout(ctx, time.Second*30)
	profile := n.identityService.WaitProfile(ctxWithTimeout, pubKey.Account())
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		return nil, "", "", err
	}
	var (
		name string
		icon string
	)
	if profile != nil {
		name = profile.Name
		icon = profile.IconCid
	}
	return pubKey, name, icon, err
}

func mapProtoPermissionToAcl(permissions aclrecordproto.AclUserPermissions) model.ParticipantPermissions {
	switch permissions {
	case aclrecordproto.AclUserPermissions_Owner:
		return model.ParticipantPermissions_Owner
	case aclrecordproto.AclUserPermissions_None:
		return model.ParticipantPermissions_NoPermissions
	case aclrecordproto.AclUserPermissions_Writer:
		return model.ParticipantPermissions_Writer
	case aclrecordproto.AclUserPermissions_Reader:
		return model.ParticipantPermissions_Reader
	case aclrecordproto.AclUserPermissions_Admin:
		return model.ParticipantPermissions_Owner
	}
	return model.ParticipantPermissions_Reader
}
