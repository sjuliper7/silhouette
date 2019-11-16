package usecase

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/profile-service/models"
	"github.com/sjuliper7/silhouette/services/profile-service/repositories"
)

type profileUsecase struct {
	repo repositories.Repository
}

func NewProfileUsecase(repo repositories.Repository) ProfileUsecase {
	return profileUsecase{repo: repo}
}

func (uc profileUsecase) GetProfile(userID int64) (profile models.Profile, err error) {
	profile, err = uc.repo.GetProfile(userID)

	if err != nil {
		logrus.Println("[usecase][GetUser] Error when calling Get User %+v", err)
		return profile, err
	}

	return profile, nil
}
