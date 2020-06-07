package mysql

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/sjuliper7/silhouette/services/profile-service/models"
	"github.com/stretchr/testify/assert"
)

type fields struct {
	DB      *sqlx.DB
	Profile models.ProfileTable
	Query   string
}

func TestGet(t *testing.T) {
	tMock := time.Now()
	tMockV2, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Split(tMock.String(), " m=")[0])
	if err != nil {
		t.Errorf("error: %v", err)
	}
	mockDb, mock, err := sqlmock.New()

	if err != nil {
		t.Error("Failed to mock db")
	}

	sx := sqlx.NewDb(mockDb, "mockdb")
	sql := `SELECT id, user_id, name, address, work_at, phone_number, gender, date_of_birth, created_at, updated_at FROM profiles WHERE user_id = ? and is_active = 1`
	rgxQuery := regexp.QuoteMeta(sql)

	successFields := fields{
		DB: sx,
		Profile: models.ProfileTable{
			ID:          1,
			UserID:      1,
			Address:     "Lumban Bulbul",
			Name:        "Juliper Simanjuntak Update",
			WorkAt:      "Start Up Companny",
			PhoneNumber: "082272194654",
			Gender:      "Laki Laki",
			DateOfBirth: "19/07/1997",
			IsActive:    true,
			CreatedAt:   tMockV2,
			UpdatedAt:   tMockV2,
		},
		Query: rgxQuery,
	}

	tests := []struct {
		Name   string
		Fields fields
		Want   []models.ProfileTable
	}{
		{
			Name:   "success-test",
			Fields: successFields,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			repo := NewMysqlProfileRepository(tt.Fields.DB)

			rows := sqlmock.NewRows([]string{"id", "user_id", "address", "name", "work_at", "phone_number", "gender", "date_of_birth", "is_active", "created_at", "updated_at"}).
				AddRow(1, 1, "Lumban Bulbul", "Juliper Simanjuntak Update", "Start Up Companny", "082272194654", "Laki Laki", "19/07/1997", 1, tMock, tMock)

			mock.ExpectQuery(tt.Fields.Query).WillReturnRows(rows)

			if got, _ := repo.Get(1); !reflect.DeepEqual(got, tt.Fields.Profile) {
				t.Errorf("profile_repository.Get(userID) = %v, want %v", got, tt.Fields.Profile)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tMock := time.Now()
	tMockV2, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Split(tMock.String(), " m=")[0])
	if err != nil {
		t.Errorf("error: %v", err)
	}
	mockDb, mock, err := sqlmock.New()

	if err != nil {
		t.Error("Failed to mock db")
	}

	sx := sqlx.NewDb(mockDb, "mockdb")
	sql := `INSERT INTO profiles(user_id, name, address, work_at, phone_number, gender, date_of_birth, is_active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	rgxQuery := regexp.QuoteMeta(sql)

	successFields := fields{
		DB: sx,
		Profile: models.ProfileTable{
			UserID:      1,
			Address:     "Lumban Bulbul",
			Name:        "Juliper Simanjuntak Update",
			WorkAt:      "Start Up Companny",
			PhoneNumber: "082272194654",
			Gender:      "Laki Laki",
			DateOfBirth: "19/07/1997",
			IsActive:    true,
			CreatedAt:   tMockV2,
			UpdatedAt:   tMockV2,
		},
		Query: rgxQuery,
	}

	tests := []struct {
		Name   string
		Fields fields
		Want   []models.ProfileTable
	}{
		{
			Name:   "success-test",
			Fields: successFields,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			input := tt.Fields.Profile
			repo := NewMysqlProfileRepository(tt.Fields.DB)

			mock.ExpectBegin()
			mock.ExpectPrepare(rgxQuery)
			mock.ExpectExec(rgxQuery).WithArgs(input.UserID, input.Name, input.Address, input.WorkAt, input.PhoneNumber, input.Gender, input.DateOfBirth, input.IsActive, input.CreatedAt, input.UpdatedAt).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			err := repo.Add(&tt.Fields.Profile)
			assert.NoError(t, err)
			assert.Equal(t, int64(1), tt.Fields.Profile.ID)
		})
	}
}

func TestUpdate(t *testing.T) {
	tMock := time.Now()

	mockDb, mock, err := sqlmock.New()

	if err != nil {
		t.Error("Failed to mock db")
	}

	sx := sqlx.NewDb(mockDb, "mockdb")
	sql := `UPDATE profiles SET user_id = ?, name = ?, address = ?, work_at = ?, phone_number = ?, gender = ?, date_of_birth = ?, is_active = ? ,created_at = ?, updated_at = ? WHERE id=?`
	rgxQuery := regexp.QuoteMeta(sql)

	successFields := fields{
		DB: sx,
		Profile: models.ProfileTable{
			ID:          1,
			UserID:      1,
			Address:     "Lumban Bulbul",
			Name:        "Juliper Simanjuntak Update",
			WorkAt:      "Start Up Companny",
			PhoneNumber: "082272194654",
			Gender:      "Laki Laki",
			DateOfBirth: "19/07/1997",
			IsActive:    true,
			CreatedAt:   tMock,
			UpdatedAt:   tMock,
		},
		Query: rgxQuery,
	}

	tests := []struct {
		Name   string
		Fields fields
		Want   []models.ProfileTable
	}{
		{
			Name:   "success-test",
			Fields: successFields,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			input := tt.Fields.Profile
			repo := NewMysqlProfileRepository(tt.Fields.DB)

			mock.ExpectBegin()
			mock.ExpectPrepare(rgxQuery)

			mock.ExpectExec(rgxQuery).WithArgs(input.UserID, input.Name, input.Address, input.WorkAt, input.PhoneNumber, input.Gender, input.DateOfBirth, input.IsActive, input.CreatedAt, input.UpdatedAt, input.ID).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			err := repo.Update(&tt.Fields.Profile)
			assert.NoError(t, err)
		})
	}
}

func TestDelete(t *testing.T) {

	mockDb, mock, err := sqlmock.New()

	if err != nil {
		t.Error("Failed to mock db")
	}

	sx := sqlx.NewDb(mockDb, "mockdb")
	sql := `UPDATE profiles SET is_active = false where user_id =?`
	rgxQuery := regexp.QuoteMeta(sql)

	successFields := fields{
		DB: sx,
		Profile: models.ProfileTable{
			UserID: 1,
		},
		Query: rgxQuery,
	}

	tests := []struct {
		Name   string
		Fields fields
		Want   []models.ProfileTable
	}{
		{
			Name:   "success-test",
			Fields: successFields,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			input := tt.Fields.Profile
			repo := NewMysqlProfileRepository(tt.Fields.DB)

			mock.ExpectBegin()
			mock.ExpectPrepare(rgxQuery)

			mock.ExpectExec(rgxQuery).WithArgs(input.UserID).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			_, err := repo.Delete(input.UserID)
			assert.NoError(t, err)
		})
	}
}
