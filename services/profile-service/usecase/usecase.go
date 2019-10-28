package usecase

import "github.com/sjuliper7/silhouette/services/profile-service/models"

type ProfileUsecase interface {
	GetProfile(userID int64) (profile models.Profile, err error)
}
