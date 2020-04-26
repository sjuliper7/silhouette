package repository

import (
	"github.com/sjuliper7/silhouette/services/user-service/models"
)

//Repository declaration type interface
type UserRepository interface {
	GetAll() (users []models.UserTable, err error)
	Add(user *models.UserTable) (err error)
	Get(userID int64) (user models.UserTable, err error)
	Update(user *models.UserTable) (err error)
	Delete(userID int64) (deleted bool, err error)
}

type ProfileRepository interface {
	Get(userID int64) (profile models.Profile, err error)
}

type KafkaRepository interface {
	PublishMessage(topic string, message []byte) (err error)
}
