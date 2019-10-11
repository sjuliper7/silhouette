package repositories

import (
	"github.com/sjuliper7/silhouette/services/user-service/models"
)

//Repository declaration type interface
type Repository interface {
	GetAlluser() []models.User
}
