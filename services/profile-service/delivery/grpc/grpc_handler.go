package grpc

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/models"
)

func (ps ProfileServer) GetProfile(ctx context.Context, params *models.UserGetProfileArguments) (*models.Profile, error) {
	UserID := params.UserID
	pf, err := ps.profileUc.GetProfile(UserID)

	if err != nil {
		logrus.Println("Failed when call [usecase][GetProfile]")
		return nil, err
	}

	var profile *models.Profile = &models.Profile{}
	profile.ID = int64(pf.ID)
	profile.UserID = int64(pf.UserId)
	profile.Address = pf.Address
	profile.WorkAt = pf.WorkAt
	profile.PhoneNumber = pf.PhoneNumber
	profile.Gender = pf.Gender

	return profile, nil
}
