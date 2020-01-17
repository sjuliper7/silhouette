package mocks

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetAll() (users []models.UserTable, err error) {
	logrus.Infof("getting all user..")
	args := m.Called()
	return args.Get(0).([] models.UserTable), args.Error(1)
}

func (m *UserRepositoryMock) Add(user *models.UserTable) (err error) {
	logrus.Infof("add user..")
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) Get(userID int64) (user models.UserTable, err error) {
	logrus.Infof("get user..")
	args := m.Called(userID)
	return args.Get(0).(models.UserTable), args.Error(1)
}

func (m *UserRepositoryMock) Update(user *models.UserTable) (err error) {
	logrus.Infof("update user..")
	args := m.Called(user)
	return args.Error(1)
}

func (m *UserRepositoryMock) Delete(userID int64) (deleted bool, err error)  {
	logrus.Infof("delete user..")
	args := m.Called(userID)
	return args.Bool(0), args.Error(1)
}


