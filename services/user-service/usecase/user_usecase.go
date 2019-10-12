package usecase

import (
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/repositories"
	"log"
)

type userUsecase struct {
	repo repositories.Repository
}

func NewUserUsecase(repo repositories.Repository) UserUsecase {
	return userUsecase{repo}
}

func (uc userUsecase) GetAlluser() (users []models.User, err error) {
	users, err = uc.repo.GetAlluser()

	if err != nil {
		log.Println("Failed when call [repositories][GetAlluser] ", err)
		return nil, err
	}

	return users, err
}
