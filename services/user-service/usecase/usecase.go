package usecase

import "github.com/sjuliper7/silhouette/services/user-service/models"

//UserUsecase interface declaration
type UserUsecase interface {
	GetAlluser() (users []models.User, err error)
	AddUser(user *models.UserTable) (err error)
	GetUser(userID int64) (user models.User, err error)
	UpdateUser(user *models.UserTable) (err error)
	DeleteUser(userID int64) (deleted bool, err error)
}
