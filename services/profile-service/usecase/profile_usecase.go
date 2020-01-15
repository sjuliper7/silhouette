package usecase

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/profile-service/models"
	"github.com/sjuliper7/silhouette/services/profile-service/repositories"
)

type profileUsecase struct {
	repo repositories.Repository
}

func NewProfileUseCase(repo repositories.Repository) ProfileUseCase {
	return &profileUsecase{repo: repo}
}

func (uc *profileUsecase) GetProfile(userID int64) (profile models.ProfileTable, err error) {
	profile, err = uc.repo.GetProfile(userID)

	if err != nil {
		logrus.Errorf("[usecase][GetUser] error when calling get user [repository][GetProfile], %v ", err)
		return profile, err
	}

	return profile, nil
}

func (uc *profileUsecase) AddProfile(profile models.ProfileTable) (err error) {
	err = uc.repo.AddProfile(&profile)

	if err != nil {
		logrus.Errorf("[usecase][AddProfile] error when calling [repository][AddProfile] ,%v", err)
		return err
	}

	return nil
}

func (uc *profileUsecase) UpdateProfile(profile models.ProfileTable) (err error) {
	err = uc.repo.UpdateProfile(&profile)

	if err != nil {
		logrus.Errorf("[usecase][UpdateProfile] error when calling [repository][UpdateProfile], %v", err)
		return err
	}

	return  nil
}

func (uc *profileUsecase) DeleteProfile(profileID int64) (err error){
	deleted, err := uc.repo.DeleteProfile(profileID)

	if err != nil {
		logrus.Errorf("[usecase][DeleteProfile] error when calling [repository][DeleteProfile], %v", err)
		return  err
	}

	logrus.Infof("profile deleted :%v", deleted)

	return  nil
}