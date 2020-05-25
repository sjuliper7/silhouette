package usecase

import (
	"os"

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
		email.Sender = os.Getenv("CONFIG_EMAIL")
		email.Subject = "Welcome Dude, " + notif.Name
		email.Message = `Hello, <br> <b>have a nice day<b/> <br> Thank you for join to our system!`

		err := notificationCase.emailRepository.SentEmail(email)
		if err != nil {
			logrus.Infof("[usecase][AccountRegisterNotification] failed when access SentEmail: %v", err)
			return err
		}
	}
	return nil
}

func (notificationCase *notificationUsecase) AccountUpdateNotifcation(notif model.Notification) error {
	if notif.Type == email {
		email := model.Email{}
		email.Receivers = []string{notif.Email}
		email.Sender = os.Getenv("CONFIG_EMAIL")
		email.Subject = "Hey yo.., " + notif.Name
		email.Message = `Hello, <br> <b>have a nice day<b/> <br> You have update your profile/account`

		err := notificationCase.emailRepository.SentEmail(email)
		if err != nil {
			logrus.Infof("[usecase][AccountUpdateNotifcation] failed when access SentEmail: %v", err)
			return err
		}
	}
	return nil
}

func (notificationCase *notificationUsecase) AccountDeleteNotification(notif model.Notification) error {
	if notif.Type == email {
		email := model.Email{}
		email.Receivers = []string{notif.Email}
		email.Sender = os.Getenv("CONFIG_EMAIL")
		email.Subject = "Hey yo.., " + notif.Name
		email.Message = `Hello, <br> <b>have a nice day<b/> <br> Your account is deleted`

		err := notificationCase.emailRepository.SentEmail(email)
		if err != nil {
			logrus.Infof("[usecase][AccountDeleteNotification] failed when access SentEmail: %v", err)
			return err
		}
	}

	return nil
}
