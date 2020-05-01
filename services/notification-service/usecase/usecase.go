package usecase

import "github.com/sjuliper7/silhouette/services/notification-service/model"

//NotificationUsecase ...
type NotificationUsecase interface {
	AccountRegisterNotification(notif model.Notification) error
	AccountUpdateNotifcation(notif model.Notification) error
	AccountDeleteNotification(notif model.Notification) error
}
