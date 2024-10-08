package process

import (
	"github.com/globalsign/mgo/bson"

	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/logging"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
)

var log = logging.Logger("notification-process")

type NotificationService interface {
	CreateAndSend(notification *model.Notification) error
}

type NotificationSender interface {
	SendNotification()
}

type Notificationable interface {
	Progress
	FinishWithNotification(notification *model.Notification, err error)
}

type notificationProcess struct {
	*progress
	notification        *model.Notification
	notificationService NotificationService
}

func NewNotificationProcess(pbType pb.ModelProcessType, notificationService NotificationService) Notificationable {
	return &notificationProcess{progress: &progress{
		id:     bson.NewObjectId().Hex(),
		done:   make(chan struct{}),
		cancel: make(chan struct{}),
		pType:  pbType,
	}, notificationService: notificationService}
}

func (n *notificationProcess) FinishWithNotification(notification *model.Notification, err error) {
	n.notification = notification
	n.Finish(err)
}

func (n *notificationProcess) SendNotification() {
	if n.notification != nil {
		notificationSendErr := n.notificationService.CreateAndSend(n.notification)
		if notificationSendErr != nil {
			log.Errorf("failed to send notification: %v", notificationSendErr)
		}
	}
}
