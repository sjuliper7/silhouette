package grpc

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/models"
)

func (profileServer ProfileServer) GetProfile(ctx context.Context, params *models.UserGetProfileArguments) (*models.Profile, error) {
	logrus.Infof("request : %v", params)

	UserID := params.UserID
	pf, err := profileServer.profileUsecase.Get(UserID)

	if err != nil {
		logrus.Errorf("[delivery][profile] failed when call [usecase][GetProfile] %v", err)
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
