package mocks

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/stretchr/testify/mock"
)

//ProfileServiceMock ...
type ProfileServiceMock struct {
	mock.Mock
}

//Get ...
func (m *ProfileServiceMock) Get(userID int64) (profile models.Profile, err error) {
	logrus.Infof("getting profile...")
	args := m.Called(userID)
	return args.Get(0).(models.Profile), args.Error(1)
}
