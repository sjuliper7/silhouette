package repositories

import (
	"github.com/sjuliper7/silhouette/services/user-service/models"
)

//Repository declaration type interface
type Repository interface {
	GetAlluser() (users []models.User, err error)
	AddUser(user *models.User) (err error)
	GetUser(userID int64) (user models.User, err error)
}
