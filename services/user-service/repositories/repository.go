package repositories

import (
	"github.com/sjuliper7/silhouette/services/user-service/models"
)

//Repository declaration type interface
type UserRepository interface {
	GetAlluser() (users []models.User, err error)
	AddUser(user *models.UserTable) (err error)
	GetUser(userID int64) (user models.UserTable, err error)
	UpdateUser(user *models.UserTable) (err error)
	DeleteUser(userID int64) (deleted bool, err error)
}

type ProfileRepository interface {
	GetProfile(UserID int64) (profile models.Profile, err error)
}

type KafkaRepository interface {
	SendMessage(profile *models.Profile) (err error)
}
