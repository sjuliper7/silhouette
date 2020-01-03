package usecase

import (
	"errors"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestUserUsecase_GetAllUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepositoryMock)

	mockListUser := make([]models.UserTable, 0)
	mockUser := models.UserTable{
		ID:        1,
		Username:  "mock",
		Email:     "mock@gmail.com",
		Name:      "mock-test",
		Role:      "mock-role",
		IsActive:  1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	mockListUser = append(mockListUser, mockUser)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetAllUser").Return(
			mockListUser,
			nil).
			Once()

		mockProfile := models.Profile{
			ID:          1,
			Address:     "mock-address",
			WorkAt:      "mock-work-at",
			PhoneNumber: "082272194654",
			Gender:      "Male",
		}

		mockProfileRepo := new(mocks.ProfileServiceMock)
		mockProfileRepo.On("GetProfile", mock.AnythingOfType("int64")).Return(mockProfile, nil)

		userCase := NewUserUsecase(mockUserRepo, mockProfileRepo)

		listOfUser, err := userCase.GetAllUser()
		assert.NoError(t, err)
		assert.Len(t, listOfUser, len(mockListUser))

		mockUserRepo.AssertExpectations(t)
		mockProfileRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		uu := make([]models.UserTable, 0)
		mockUserRepo.On("GetAllUser").Return(
			uu,
			errors.New("Unexpected Error")).
			Once()

		mockProfileRepo := new(mocks.ProfileServiceMock)

		userCase := NewUserUsecase(mockUserRepo, mockProfileRepo)
		listOfUser, err := userCase.GetAllUser()
		assert.NoError(t, err)
		assert.Len(t, listOfUser, 0)

		mockUserRepo.AssertExpectations(t)
		mockProfileRepo.AssertExpectations(t)
	})

}