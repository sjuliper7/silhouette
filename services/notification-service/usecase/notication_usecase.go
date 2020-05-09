package usecase

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/notification-service/model"
	"github.com/sjuliper7/silhouette/services/notification-service/repository"
)

const (
	email = "email"
)

type notificationUsecase struct {
	emailRepository repository.EmailRepository
}

//NewNotificatonUsecase ...
func NewNotificatonUsecase(emailRepository repository.EmailRepository) NotificationUsecase {
	return &notificationUsecase{emailRepository: emailRepository}
}

func (notificationCase *notificationUsecase) AccountRegisterNotification(notif model.Notification) error {
	if notif.Type == email {
		email := model.Email{}
		email.Receivers = []string{notif.Email}
		email.Sender = ""
		email.Subject = "Welcome Dude"
		email.Message = `Hello, <br> <b>have a nice data<b/> <br> Thank you for join to our system!`

		err := notificationCase.emailRepository.SentEmail(email)
		if err != nil {
			logrus.Infof("[usecase][AccountRegisterNotification] failed when access SentEmail %v", err)
			return err
		}
	}
	return nil
}

func (notificationCase *notificationUsecase) AccountUpdateNotifcation(notif model.Notification) error {
	return nil
}

func (notificationCase *notificationUsecase) AccountDeleteNotification(notif model.Notification) error {
	return nil
}
