package mocks

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/stretchr/testify/mock"
)

type ProfileServiceMock struct {
	mock.Mock
}

func (m *ProfileServiceMock) GetProfile(userID int64) (profile models.Profile, err error) {
	logrus.Infof("getting profile with user_id %v", userID)
	args := m.Called(userID)
	return args.Get(0).(models.Profile), args.Error(1)
}


