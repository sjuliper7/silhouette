package delivery

import (
	"context"
	"github.com/sjuliper7/silhouette/common/protocs"
	"log"
)

func (ps ProfileServer) GetProfile(ctx context.Context, params *protocs.UserGetProfileArguments) (*protocs.Profile, error) {
	UserID := params.UserID
	pf, err := ps.usecase.GetUser(UserID)

	if err != nil {
		log.Println("Failed when call [usecase][GetProfile]")
		return nil, err
	}

	var profile *protocs.Profile = &protocs.Profile{}
	profile.ID = int64(pf.ID)
	profile.UserID = int64(pf.UserId)
	profile.Address = pf.Address
	profile.WorkAt = pf.WorkAt
	profile.PhoneNumber = pf.PhoneNumber
	profile.Gender = pf.Gender

	return profile, nil
}
