package mocks

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/stretchr/testify/mock"
)

// UserUsecaseMock ...
type UserUsecaseMock struct {
	mock.Mock
}

//GetAll ...
func (m *UserUsecaseMock) GetAll() (users []models.User, err error) {
	logrus.Infof("get all users")
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

//Add ...
func (m *UserUsecaseMock) Add(user *models.User) (err error) {
	logrus.Infof("add users")
	args := m.Called(user)
	return args.Error(0)
}

//Get ...
func (m *UserUsecaseMock) Get(userID int64) (user models.User, err error) {
	logrus.Infof("get user")
	args := m.Called(userID)
	return args.Get(0).(models.User), args.Error(1)
}

//Update ...
func (m *UserUsecaseMock) Update(us models.User) (user models.User, err error) {
	logrus.Infof("update user")
	args := m.Called(us)
	return args.Get(0).(models.User), args.Error(1)
}

//Delete ...
func (m *UserUsecaseMock) Delete(userID int64) (deleted bool, err error) {
	logrus.Infof("delete user")
	args := m.Called(userID)
	return args.Bool(0), args.Error(1)
}

//Login
func (m *UserUsecaseMock) Login(user *models.User) (string, error) {
	logrus.Infof("login...")
	args := m.Called(user)
	return args.String(0), args.Error(1)
}
