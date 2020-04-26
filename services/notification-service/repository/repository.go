package repository

import "github.com/sjuliper7/silhouette/services/notification-service/model"

// EmailRepository ...
type EmailRepository interface {
	SentEmail(email model.Email) error
}
