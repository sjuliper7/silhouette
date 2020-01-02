package mocks

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetAllUser() (users []models.UserTable, err error) {
	logrus.Println("getting all user..")
	args := m.Called()
	return args.Get(0).([] models.UserTable), args.Error(1)
}

func (m *UserRepositoryMock) AddUser(user *models.UserTable) (err error) {
	logrus.Println("add user..")
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) GetUser(userID int64) (user models.UserTable, err error) {
	logrus.Println("get user..")
	args := m.Called(userID)
	return args.Get(0).(models.UserTable), args.Error(1)
}

func (m *UserRepositoryMock) UpdateUser(user *models.UserTable) (err error) {
	logrus.Println("update user..")
	args := m.Called(user)
	return args.Error(1)
}

func (m *UserRepositoryMock) DeleteUser(userID int64) (deleted bool, err error)  {
	logrus.Println("delete user..")
	args := m.Called(userID)
	return args.Bool(0), args.Error(1)
}


