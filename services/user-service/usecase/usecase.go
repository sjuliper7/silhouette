package usecase

import "github.com/sjuliper7/silhouette/services/user-service/models"

//UserUsecase interface declaration
type UserUsecase interface {
	GetAlluser() (users []models.User, err error)
}
