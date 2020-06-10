package usecase

import (
	"reflect"
	"testing"
	"time"

	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type fields struct {
	Users      []models.User
	User       models.User
	UserTable  models.UserTable
	Profile    models.Profile
	UserTables []models.UserTable
	ID         int64
}

func TestGetAllUser(t *testing.T) {
	tMock := time.Now()

	userTables := []models.UserTable{
		{
			ID:        1,
			Username:  "mock",
			Email:     "mock@gmail.com",
			Role:      "mock-role",
			IsActive:  1,
			CreatedAt: tMock,
			UpdatedAt: tMock,
		},
	}

	profile := models.Profile{
		ID:          1,
		UserID:      1,
		Address:     "Lumban Bulbul",
		Name:        "Juliper Simanjuntak",
		WorkAt:      "Start Up Companny",
		PhoneNumber: "082272194654",
		Gender:      "Laki Laki",
		DateOfBirth: "19/07/1997",
		CreatedAt:   tMock.String(),
		UpdatedAt:   tMock.String(),
	}

	users := []models.User{
		{
			ID:        1,
			Username:  "mock",
			Email:     "mock@gmail.com",
			Role:      "mock-role",
			CreatedAt: tMock.String(),
			UpdatedAt: tMock.String(),
			Profile:   profile,
		},
	}

	successFields := fields{
		Users:      users,
		UserTables: userTables,
		Profile:    profile,
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
			mockRepo := new(mocks.UserRepositoryMock)
			mockProfileService := new(mocks.ProfileServiceMock)
			mockKafkaRepository := new(mocks.KafkaRepositoryMock)

			mockRepo.On("GetAll").Return(tt.Fields.UserTables, nil)
			mockProfileService.On("Get", mock.AnythingOfType("int64")).Return(tt.Fields.Profile, nil)

			userCase := NewUserUsecase(mockRepo, mockProfileService, mockKafkaRepository)
			got, err := userCase.GetAll()

			assert.NoError(t, err)
			assert.Len(t, got, len(tt.Fields.Users))
			if !reflect.DeepEqual(got, tt.Fields.Users) {
				t.Errorf("TestGetAll = %v, want %v", got, tt.Fields.Users)
			}

		})
	}
}

func TestAdd(t *testing.T) {
	tMock := time.Now()
	userTable := models.UserTable{
		Username:  "mock",
		Email:     "mock@gmail.com",
		Role:      "mock-role",
		IsActive:  1,
		CreatedAt: tMock,
		UpdatedAt: tMock,
	}

	profile := models.Profile{
		UserID:      1,
		Address:     "Lumban Bulbul",
		Name:        "Juliper Simanjuntak",
		WorkAt:      "Start Up Companny",
		PhoneNumber: "082272194654",
		Gender:      "Laki Laki",
		DateOfBirth: "19/07/1997",
		CreatedAt:   tMock.String(),
		UpdatedAt:   tMock.String(),
	}

	user := models.User{
		Username:  "mock",
		Email:     "mock@gmail.com",
		Role:      "mock-role",
		CreatedAt: tMock.String(),
		UpdatedAt: tMock.String(),
		Profile:   profile,
	}

	successFields := fields{
		User:      user,
		UserTable: userTable,
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
			mockRepo := new(mocks.UserRepositoryMock)
			mockProfileService := new(mocks.ProfileServiceMock)
			mockKafkaRepository := new(mocks.KafkaRepositoryMock)

			mockRepo.On("Add", mock.AnythingOfType("*models.UserTable")).Return(nil)
			mockKafkaRepository.On("PublishMessage", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).Return(nil)

			userCase := NewUserUsecase(mockRepo, mockProfileService, mockKafkaRepository)
			err := userCase.Add(&tt.Fields.User)

			assert.NoError(t, err)
		})
	}
}

func TestGet(t *testing.T) {
	tMock := time.Now()

	userTable := models.UserTable{
		ID:        1,
		Username:  "mock",
		Email:     "mock@gmail.com",
		Role:      "mock-role",
		IsActive:  1,
		CreatedAt: tMock,
		UpdatedAt: tMock,
	}

	profile := models.Profile{
		ID:          1,
		UserID:      1,
		Address:     "Lumban Bulbul",
		Name:        "Juliper Simanjuntak",
		WorkAt:      "Start Up Companny",
		PhoneNumber: "082272194654",
		Gender:      "Laki Laki",
		DateOfBirth: "19/07/1997",
		CreatedAt:   tMock.String(),
		UpdatedAt:   tMock.String(),
	}

	user := models.User{
		ID:        1,
		Username:  "mock",
		Email:     "mock@gmail.com",
		Role:      "mock-role",
		CreatedAt: tMock.String(),
		UpdatedAt: tMock.String(),
		Profile:   profile,
	}

	successFields := fields{
		User:      user,
		UserTable: userTable,
		Profile:   profile,
		ID:        1,
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
			mockRepo := new(mocks.UserRepositoryMock)
			mockProfileService := new(mocks.ProfileServiceMock)
			mockKafkaRepository := new(mocks.KafkaRepositoryMock)

			mockRepo.On("Get", mock.AnythingOfType("int64")).Return(tt.Fields.UserTable, nil)
			mockProfileService.On("Get", mock.AnythingOfType("int64")).Return(tt.Fields.Profile, nil)

			userCase := NewUserUsecase(mockRepo, mockProfileService, mockKafkaRepository)
			got, err := userCase.Get(tt.Fields.ID)

			assert.NoError(t, err)
			if !reflect.DeepEqual(got, tt.Fields.User) {
				t.Errorf("TestGet() = %v, want %v", got, tt.Fields.Users)
			}

		})
	}
}

func TestUpdate(t *testing.T) {
	tMock := time.Now()
	userTables := []models.UserTable{
		{
			ID:        1,
			Username:  "mock-update",
			Email:     "mock@gmail.com",
			Role:      "mock-role",
			IsActive:  1,
			CreatedAt: tMock,
			UpdatedAt: tMock,
		},
	}

	profile := models.Profile{
		UserID:      1,
		Address:     "Lumban Bulbul",
		Name:        "Juliper Simanjuntak Update",
		WorkAt:      "Start Up Companny",
		PhoneNumber: "082272194654",
		Gender:      "Laki Laki",
		DateOfBirth: "19/07/1997",
		CreatedAt:   tMock.String(),
		UpdatedAt:   tMock.String(),
	}

	users := []models.User{
		{
			Username:  "mock",
			Email:     "mock@gmail.com",
			Role:      "mock-role",
			CreatedAt: tMock.String(),
			UpdatedAt: tMock.String(),
			Profile:   profile,
		},
	}

	successFields := fields{
		Users:      users,
		UserTables: userTables,
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
			mockRepo := new(mocks.UserRepositoryMock)
			mockProfileService := new(mocks.ProfileServiceMock)
			mockKafkaRepository := new(mocks.KafkaRepositoryMock)

			mockRepo.On("Update", mock.AnythingOfType("*models.UserTable")).Return(nil)
			mockRepo.On("Get", mock.AnythingOfType("int64")).Return(tt.Fields.UserTable, nil)
			mockKafkaRepository.On("PublishMessage", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).Return(nil)
			mockProfileService.On("Get", mock.AnythingOfType("int64")).Return(tt.Fields.Profile, nil)

			userCase := NewUserUsecase(mockRepo, mockProfileService, mockKafkaRepository)
			_, err := userCase.Update(tt.Fields.User)

			assert.NoError(t, err)

		})
	}
}

func TestDelete(t *testing.T) {

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: fields{ID: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			mockRepo := new(mocks.UserRepositoryMock)
			mockProfileService := new(mocks.ProfileServiceMock)
			mockKafkaRepository := new(mocks.KafkaRepositoryMock)

			mockRepo.On("Delete", mock.AnythingOfType("int64")).Return(true, nil)
			mockKafkaRepository.On("PublishMessage", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).Return(nil)

			userCase := NewUserUsecase(mockRepo, mockProfileService, mockKafkaRepository)
			_, err := userCase.Delete(tt.Fields.ID)

			assert.NoError(t, err)

		})
	}

}

func TestLogin(t *testing.T) {

	tMock := time.Now()

	userTable := models.UserTable{
		ID:        1,
		Username:  "mock",
		Email:     "mock@gmail.com",
		Password:  "bycrypt(ganteng)",
		Role:      "mock-role",
		IsActive:  1,
		CreatedAt: tMock,
		UpdatedAt: tMock,
	}

	successFields := fields{
		UserTable: userTable,
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
			mockRepo := new(mocks.UserRepositoryMock)
			mockProfileService := new(mocks.ProfileServiceMock)
			mockKafkaRepository := new(mocks.KafkaRepositoryMock)

			mockRepo.On("GetByEmail", mock.AnythingOfType("string")).Return(tt.Fields.UserTable, nil)

			userMock := models.User{
				Email:    "mock@gmail.com",
				Password: "ganteng",
			}
			userCase := NewUserUsecase(mockRepo, mockProfileService, mockKafkaRepository)
			_, err := userCase.Login(&userMock)
			assert.NoError(t, err)

		})
	}

}
