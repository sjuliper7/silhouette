package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type fields struct {
	Users []models.User
	User  models.User
	ID    int
}

func TestFetchUser(t *testing.T) {
	tMock := time.Now()

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

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: fields{Users: users},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			usecaseMock := new(mocks.UserUsecaseMock)
			usecaseMock.On("GetAll").Return(tt.Fields.Users, nil)

			deliveryHandler := NewUserDelivery(usecaseMock)
			req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(deliveryHandler.Resource)
			handler.ServeHTTP(rec, req)

			if status := rec.Code; status != http.StatusOK {
				assert.Errorf(t, nil, "returned wrong status code: got %v want %v", status, http.StatusOK)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	tMock := time.Now()

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

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: fields{User: user},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			usecaseMock := new(mocks.UserUsecaseMock)
			usecaseMock.On("Get", mock.AnythingOfType("int64")).Return(tt.Fields.User, nil)

			id := 1

			deliveryHandler := NewUserDelivery(usecaseMock)
			req, err := http.NewRequest(http.MethodGet, "/api/v1/users/"+strconv.Itoa(id), nil)
			assert.NoError(t, err)

			rec := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/api/v1/users/{id}", deliveryHandler.Resource)
			router.ServeHTTP(rec, req)

			if status := rec.Code; status != http.StatusOK {
				assert.Errorf(t, nil, "returned wrong status code: got %v want %v", status, http.StatusOK)
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	tMock := time.Now()

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

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: fields{User: user},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			usecaseMock := new(mocks.UserUsecaseMock)
			usecaseMock.On("Get", mock.AnythingOfType("int64")).Return(tt.Fields.User, nil)
			usecaseMock.On("Add", mock.AnythingOfType("*models.User")).Return(nil)

			payload := map[string]interface{}{
				"username":      "mock",
				"email":         "mock@gmail.com",
				"role":          "mock-role",
				"address":       "Lumban Bulbul",
				"name":          "Juliper Simanjuntak",
				"work_at":       "Start Up Companny",
				"phone_number":  "082272194654",
				"gender":        "Laki Laki",
				"date_of_birth": "19/07/1997",
			}

			jsonPayload, err := json.Marshal(payload)
			assert.NoError(t, err)

			deliveryHandler := NewUserDelivery(usecaseMock)
			req, err := http.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(string(jsonPayload)))
			assert.NoError(t, err)

			rec := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/api/v1/users", deliveryHandler.Resource)
			router.ServeHTTP(rec, req)

			if status := rec.Code; status != http.StatusOK {
				assert.Errorf(t, nil, "returned wrong status code: got %v want %v", status, http.StatusOK)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	tMock := time.Now()

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

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: fields{User: user, ID: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			usecaseMock := new(mocks.UserUsecaseMock)
			usecaseMock.On("Update", mock.AnythingOfType("models.User")).Return(tt.Fields.User, nil)

			payload := map[string]interface{}{
				"username":      "mock",
				"email":         "mock@gmail.com",
				"role":          "mock-role",
				"address":       "Lumban Bulbul",
				"name":          "Juliper Simanjuntak",
				"work_at":       "Start Up Companny",
				"phone_number":  "082272194654",
				"gender":        "Laki Laki",
				"date_of_birth": "19/07/1997",
			}

			jsonPayload, err := json.Marshal(payload)
			assert.NoError(t, err)

			deliveryHandler := NewUserDelivery(usecaseMock)
			req, err := http.NewRequest(http.MethodPut, "/api/v1/users/"+strconv.Itoa(tt.Fields.ID), strings.NewReader(string(jsonPayload)))
			assert.NoError(t, err)

			rec := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/api/v1/users/{id}", deliveryHandler.Resource)
			router.ServeHTTP(rec, req)

			if status := rec.Code; status != http.StatusOK {
				assert.Errorf(t, nil, "returned wrong status code: got %v want %v", status, http.StatusOK)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {

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
			usecaseMock := new(mocks.UserUsecaseMock)
			usecaseMock.On("Delete", mock.AnythingOfType("int64")).Return(true, nil)

			deliveryHandler := NewUserDelivery(usecaseMock)
			req, err := http.NewRequest(http.MethodDelete, "/api/v1/users/"+strconv.Itoa(tt.Fields.ID), nil)
			assert.NoError(t, err)

			rec := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/api/v1/users/{id}", deliveryHandler.Resource)
			router.ServeHTTP(rec, req)

			if status := rec.Code; status != http.StatusOK {
				assert.Errorf(t, nil, "returned wrong status code: got %v want %v", status, http.StatusOK)
			}
		})
	}
}
