package grpc

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/models"
)

//GetProfile ...
func (profileServer ProfileService) GetProfile(ctx context.Context, params *models.UserGetProfileArguments) (*models.Profile, error) {
	logrus.Infof("request : %v", params)

	UserID := params.UserID
	pf, err := profileServer.profileUsecase.Get(UserID)

	if err != nil {
		logrus.Errorf("[delivery][profile] failed when call [usecase][GetProfile] %v", err)
		return nil, err
	}

	var profile *models.Profile = &models.Profile{}
	profile.ID = pf.ID
	profile.UserID = pf.UserID
	profile.Address = pf.Address
	profile.WorkAt = pf.WorkAt
	profile.PhoneNumber = pf.PhoneNumber
	profile.Gender = pf.Gender
	profile.IsActive = pf.IsActive
	profile.CreatedAt = pf.CreatedAt.String()
	profile.UpdatedAt = pf.UpdatedAt.String()
	profile.Name = pf.Name
	profile.DateOfBirth = pf.DateOfBirth

	return profile, nil
}
