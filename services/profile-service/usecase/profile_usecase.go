package usecase

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/profile-service/models"
	"github.com/sjuliper7/silhouette/services/profile-service/repositories"
	"time"
)

type profileUsecase struct {
	profileRepository repositories.Repository
}

func NewProfileUseCase(repo repositories.Repository) ProfileUseCase {
	return &profileUsecase{profileRepository: repo}
}

func (profileUsecase *profileUsecase) Get(userID int64) (profile models.ProfileTable, err error) {
	profile, err = profileUsecase.profileRepository.Get(userID)

	if err != nil {
		logrus.Errorf("[usecase][GetUser] error when calling get user [profileRepository][Get], %v ", err)
		return profile, err
	}

	return profile, nil
}

func (profileUsecase *profileUsecase) Add(profile models.ProfileTable) (err error) {
	profile.CreatedAt = time.Now()
	profile.UpdatedAt = time.Now()
	err = profileUsecase.profileRepository.Add(&profile)

	if err != nil {
		logrus.Errorf("[profileUsecase][Add] error when calling [profileRepository][Add] ,%v", err)
		return err
	}

	return nil
}

func (profileUsecase *profileUsecase) Update(profile models.ProfileTable) (err error) {
	err = profileUsecase.profileRepository.Update(&profile)

	if err != nil {
		logrus.Errorf("[profileUsecase][Update] error when calling [profileRepository][Update], %v", err)
		return err
	}

	return  nil
}

func (profileUsecase *profileUsecase) Delete(profileID int64) (err error){
	deleted, err := profileUsecase.profileRepository.Delete(profileID)

	if err != nil {
		logrus.Errorf("[usecase][DeleteProfile] error when calling [profileRepository][Delete], %v", err)
		return  err
	}

	logrus.Infof("[usecase][DeleteProfile] profile deleted :%v", deleted)

	return  nil
}