package usecase

import "github.com/sjuliper7/silhouette/services/profile-service/models"

type ProfileUseCase interface {
	Get(userID int64) (profile models.ProfileTable, err error)
	Add(profile models.ProfileTable) (err error)
	Update(profile models.ProfileTable) (err error)
	Delete(profileID int64) (err error)
}
