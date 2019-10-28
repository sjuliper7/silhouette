package services

import (
	"context"
	"github.com/sjuliper7/silhouette/common/config"
	"github.com/sjuliper7/silhouette/common/protocs"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/repositories"
	"google.golang.org/grpc"
	"log"
)

type profileRepository struct {
	clientProfile protocs.ProfilesClient
}

func NewProfileRepository() (repositories.ProfileRepository, error) {
	repo := profileRepository{}
	profilePort := config.SERVICE_PROFILE_PORT
	conn, err := grpc.Dial(profilePort, grpc.WithInsecure())

	if err != nil {
		log.Fatalln("Could not connect to profile service", profilePort)
		return nil, err
	}

	repo.clientProfile = protocs.NewProfilesClient(conn)

	return repo, nil
}

func (repo profileRepository) GetProfile(UserID int64) (profile models.Profile, err error) {
	profile = models.Profile{}
	result, err := repo.clientProfile.GetProfile(context.Background(), &protocs.UserGetProfileArguments{
		UserID: 1,
	})

	if err != nil {
		log.Println("[repository][profile-service][GetProfile] while grpc GetProfile")
		return profile, err
	}

	profile.ID = result.ID
	profile.UserId = result.UserID
	profile.Gender = profile.Gender
	profile.Address = profile.Address
	profile.PhoneNumber = profile.PhoneNumber
	profile.WorkAt = profile.WorkAt

	return profile, nil
}
