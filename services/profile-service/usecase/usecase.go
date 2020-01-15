package usecase

import "github.com/sjuliper7/silhouette/services/profile-service/models"

type ProfileUseCase interface {
	GetProfile(userID int64) (profile models.ProfileTable, err error)
	AddProfile(profile models.ProfileTable) (err error)
	UpdateProfile(profile models.ProfileTable) (err error)
	DeleteProfile(profileID int64) (err error)
}
