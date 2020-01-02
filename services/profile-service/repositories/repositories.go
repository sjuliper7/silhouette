package repositories

import "github.com/sjuliper7/silhouette/services/profile-service/models"

type Repository interface {
	GetProfile(userID int64) (profile models.ProfileTable, err error)
}
