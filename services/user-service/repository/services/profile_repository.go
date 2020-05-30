package services

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/config"
	models2 "github.com/sjuliper7/silhouette/commons/models"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/repository"
	"google.golang.org/grpc"
)

type profileRepository struct {
	clientProfile models2.ProfilesClient
}

//NewProfileRepository ...
func NewProfileRepository() (repository.ProfileRepository, error) {
	repo := profileRepository{}
	profilePort := config.SERVICE_PROFILE_PORT
	conn, err := grpc.Dial(profilePort, grpc.WithInsecure())

	if err != nil {
		logrus.Errorf("Could not connect to profile service: %v", profilePort)
		return nil, err
	}

	repo.clientProfile = models2.NewProfilesClient(conn)
	logrus.Infof("Yay connected to profile services...")

	return repo, nil
}

func (repo profileRepository) Get(userID int64) (profile models.Profile, err error) {
	profile = models.Profile{}
	rProfile, err := repo.clientProfile.GetProfile(context.Background(), &models2.UserGetProfileArguments{
		UserID: userID,
	})

	if err != nil {
		logrus.Errorf("[repository][profile-service][GetProfile] while grpc GetProfile: %v", err)
		return profile, err
	}

	profile.ID = rProfile.ID
	profile.Address = rProfile.Address
	profile.WorkAt = rProfile.WorkAt
	profile.PhoneNumber = rProfile.PhoneNumber
	profile.Gender = rProfile.Gender
	profile.IsActive = rProfile.IsActive
	profile.UserID = rProfile.UserID
	profile.CreatedAt = rProfile.CreatedAt
	profile.UpdatedAt = rProfile.UpdatedAt

	return profile, nil
}
