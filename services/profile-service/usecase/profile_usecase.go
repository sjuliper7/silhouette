package usecase

import (
	"github.com/sjuliper7/silhouette/services/profile-service/models"
	"github.com/sjuliper7/silhouette/services/profile-service/repositories"
	"log"
)

type profileUsecase struct {
	repo repositories.Repository
}

func NewProfileUsecase(repo repositories.Repository) ProfileUsecase {
	return profileUsecase{repo: repo}
}

func (usecase profileUsecase) GetUser(userID int64) (profile models.Profile, err error) {
	profile, err = usecase.repo.GetUser(userID)

	if err != nil {
		log.Println("[usecase][GetUser] Error when calling Get User %+v", err)
		return profile, err
	}

	return profile, nil
}
