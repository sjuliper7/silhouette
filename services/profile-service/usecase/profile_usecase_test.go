package usecase

import (
	"reflect"
	"testing"
	"time"

	"github.com/sjuliper7/silhouette/services/profile-service/models"
	"github.com/sjuliper7/silhouette/services/profile-service/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type fields struct {
	Profile models.ProfileTable
	UserID  int64
}

func TestGet(t *testing.T) {
	tMock := time.Now()

	profile := models.ProfileTable{
		ID:          1,
		UserID:      1,
		Address:     "Lumban Bulbul",
		Name:        "Juliper Simanjuntak",
		WorkAt:      "Start Up Companny",
		PhoneNumber: "082272194654",
		Gender:      "Laki Laki",
		DateOfBirth: "19/07/1997",
		CreatedAt:   tMock,
		UpdatedAt:   tMock,
	}

	successFields := fields{
		Profile: profile,
		UserID:  1,
	}

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: successFields,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			mockRepo := new(mocks.ProfileRepositoryMock)

			mockRepo.On("Get", mock.AnythingOfType("int64")).Return(tt.Fields.Profile, nil)

			profileCase := NewProfileUseCase(mockRepo)
			got, err := profileCase.Get(tt.Fields.UserID)

			assert.NoError(t, err)
			if !reflect.DeepEqual(got, tt.Fields.Profile) {
				t.Errorf("TestGet() = %v, want %v", got, tt.Fields.Profile)
			}

		})
	}
}

func TestAdd(t *testing.T) {
	tMock := time.Now()

	profile := models.ProfileTable{
		UserID:      1,
		Address:     "Lumban Bulbul",
		Name:        "Juliper Simanjuntak",
		WorkAt:      "Start Up Companny",
		PhoneNumber: "082272194654",
		Gender:      "Laki Laki",
		DateOfBirth: "19/07/1997",
		CreatedAt:   tMock,
		UpdatedAt:   tMock,
	}

	successFields := fields{
		Profile: profile,
	}

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: successFields,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			mockRepo := new(mocks.ProfileRepositoryMock)

			mockRepo.On("Add", mock.AnythingOfType("*models.ProfileTable")).Return(nil)

			profileCase := NewProfileUseCase(mockRepo)
			err := profileCase.Add(tt.Fields.Profile)

			assert.NoError(t, err)

		})
	}
}

func TestUpdate(t *testing.T) {
	tMock := time.Now()

	profile := models.ProfileTable{
		ID:          1,
		UserID:      1,
		Address:     "Lumban Bulbul",
		Name:        "Juliper Simanjuntak",
		WorkAt:      "Start Up Companny",
		PhoneNumber: "082272194654",
		Gender:      "Laki Laki",
		DateOfBirth: "19/07/1997",
		CreatedAt:   tMock,
		UpdatedAt:   tMock,
	}

	successFields := fields{
		Profile: profile,
	}

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: successFields,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			mockRepo := new(mocks.ProfileRepositoryMock)

			mockRepo.On("Get", mock.AnythingOfType("int64")).Return(tt.Fields.Profile, nil)
			mockRepo.On("Update", mock.AnythingOfType("*models.ProfileTable")).Return(nil)

			profileCase := NewProfileUseCase(mockRepo)
			err := profileCase.Update(tt.Fields.Profile)

			assert.NoError(t, err)

		})
	}
}

func TestDelete(t *testing.T) {

	successFields := fields{
		UserID: 1,
	}

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: successFields,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			mockRepo := new(mocks.ProfileRepositoryMock)

			mockRepo.On("Delete", mock.AnythingOfType("int64")).Return(true, nil)

			profileCase := NewProfileUseCase(mockRepo)
			err := profileCase.Delete(tt.Fields.UserID)

			assert.NoError(t, err)

		})
	}
}
