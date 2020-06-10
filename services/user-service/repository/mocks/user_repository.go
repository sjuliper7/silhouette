package mocks

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/stretchr/testify/mock"
)

//UserRepositoryMock ...
type UserRepositoryMock struct {
	mock.Mock
}

// GetAll is mock function
func (m *UserRepositoryMock) GetAll() (users []models.UserTable, err error) {
	logrus.Infof("getting all user...")
	args := m.Called()
	return args.Get(0).([]models.UserTable), args.Error(1)
}

// Add is mock function
func (m *UserRepositoryMock) Add(user *models.UserTable) (err error) {
	logrus.Infof("add user...")
	args := m.Called(user)
	return args.Error(0)
}

// Get is mock function
func (m *UserRepositoryMock) Get(userID int64) (user models.UserTable, err error) {
	logrus.Infof("get user...")
	args := m.Called(userID)
	return args.Get(0).(models.UserTable), args.Error(1)
}

// Update is mock function
func (m *UserRepositoryMock) Update(user *models.UserTable) (err error) {
	logrus.Infof("update user...")
	args := m.Called(user)
	return args.Error(0)
}

// Delete is mock function
func (m *UserRepositoryMock) Delete(userID int64) (deleted bool, err error) {
	logrus.Infof("delete user...")
	args := m.Called(userID)
	return args.Bool(0), args.Error(1)
}

//GetByEmail ...
func (m *UserRepositoryMock) GetByEmail(email string) (user models.UserTable, err error) {
	logrus.Infof("get user by email...")
	args := m.Called(email)
	return args.Get(0).(models.UserTable), args.Error(1)
}
