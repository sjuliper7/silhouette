package usecase

import "github.com/sjuliper7/silhouette/services/notification-service/model"

type notificationUsecase struct {
}

//NewNotificatonUsecase ...
func NewNotificatonUsecase() NotificationUsecase {
	return &notificationUsecase{}
}

func (notificationCase *notificationUsecase) AccountRegisterNotification(notif model.Notification) error {
	return nil
}

func (notificationCase *notificationUsecase) AccountUpdateNotifcation(notif model.Notification) error {
	return nil
}

func (notificationCase *notificationUsecase) AccountDeleteNotification(notif model.Notification) error {
	return nil
}
