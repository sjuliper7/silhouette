package usecase

import (
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/repositories"
)

type userUsecase struct {
	repo repositories.Repository
}

func NewUserUsecase(repo repositories.Repository) UserUsecase {
	return userUsecase{repo}
}

func (uc userUsecase) GetAlluser() (users []models.User) {
	users = uc.repo.GetAlluser()
	return
}
