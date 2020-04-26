package repository

import "github.com/sjuliper7/silhouette/services/profile-service/models"

type Repository interface {
	Get(userID int64) (profile models.ProfileTable, err error)
	Add(profile *models.ProfileTable) (err error)
	Update(profile *models.ProfileTable) (err error)
	Delete(userID int64)(deleted bool, err error)
}
