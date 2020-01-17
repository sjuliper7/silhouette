package usecase

import "github.com/sjuliper7/silhouette/services/user-service/models"

//UserUsecase interface declaration
type UserUsecase interface {
	GetAll() (users []models.User, err error)
	Add(user *models.User) (err error)
	Get(userID int64) (user models.User, err error)
	Update(us models.User) (user models.User,err error)
	Delete(userID int64) (deleted bool, err error)
}
