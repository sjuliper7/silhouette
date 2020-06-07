package mocks

import (
	"github.com/sjuliper7/silhouette/services/profile-service/models"
	"github.com/stretchr/testify/mock"
)

//ProfileRepositoryMock ...
type ProfileRepositoryMock struct {
	mock.Mock
}

//Get ...
func (m *ProfileRepositoryMock) Get(userID int64) (profile models.ProfileTable, err error) {
	args := m.Called(userID)
	return args.Get(0).(models.ProfileTable), args.Error(1)
}

//Add ...
func (m *ProfileRepositoryMock) Add(profile *models.ProfileTable) (err error) {
	args := m.Called(profile)
	return args.Error(0)
}

//Update ...
func (m *ProfileRepositoryMock) Update(profile *models.ProfileTable) (err error) {
	args := m.Called(profile)
	return args.Error(0)
}

//Delete ...
func (m *ProfileRepositoryMock) Delete(userID int64) (deleted bool, err error) {
	args := m.Called(userID)
	return args.Bool(0), args.Error(1)
}
