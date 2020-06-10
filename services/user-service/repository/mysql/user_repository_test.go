package mysql

import (
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/stretchr/testify/assert"
)

type fields struct {
	DB   *sqlx.DB
	User *models.UserTable
}

func TestGetAll(t *testing.T) {

	tMock := time.Now()

	mockDb, mock, err := sqlmock.New()

	if err != nil {
		t.Error("Failed to mock db")
	}

	sx := sqlx.NewDb(mockDb, "mockdb")

	sql := `SELECT id, username, email, role, is_active, created_at, updated_at FROM users where is_active = true`
	rgxQuery := regexp.QuoteMeta(sql)

	test := []struct {
		Name   string
		Fields fields
		Want   []models.UserTable
	}{
		{
			Name:   "success-test",
			Fields: fields{DB: sx},
			Want: []models.UserTable{
				models.UserTable{
					ID:        1,
					Username:  "sjuliper7",
					Email:     "sjuliper7@gmail.com",
					Role:      "user",
					IsActive:  1,
					CreatedAt: tMock,
					UpdatedAt: tMock,
				},
				models.UserTable{
					ID:        2,
					Username:  "yesica",
					Email:     "yesica@gmail.com",
					Role:      "user",
					IsActive:  1,
					CreatedAt: tMock,
					UpdatedAt: tMock,
				},
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.Name, func(t *testing.T) {
			repo := NewUserMysqlRepository(tt.Fields.DB)

			rows := sqlmock.NewRows([]string{"id", "username", "email", "role", "is_active", "created_at", "updated_at"}).
				AddRow(1, "sjuliper7", "sjuliper7@gmail.com", "user", 1, tMock, tMock).
				AddRow(2, "yesica", "yesica@gmail.com", "user", 1, tMock, tMock)

			mock.ExpectQuery(rgxQuery).WillReturnRows(rows)

			resultMock, err := repo.GetAll()
			if err != nil {
				t.Errorf("Error When call GetAllUser")
			}

			if !reflect.DeepEqual(resultMock, tt.Want) {
				t.Errorf("workRepo.GetAllWork() = %v, want %v", resultMock, tt.Want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tMock := time.Now()

	mockDB, mock, err := sqlmock.New()

	if err != nil {
		t.Error("Failed to mock db")
	}

	sx := sqlx.NewDb(mockDB, "mockdb")

	sql := `INSERT INTO users(password, username, email, role, is_active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	rgxQuery := regexp.QuoteMeta(sql)

	user := &models.UserTable{
		ID:        0,
		Username:  "sjuliper8",
		Email:     "sjuliper8@gmail.com",
		Role:      "admin",
		IsActive:  1,
		CreatedAt: tMock,
		UpdatedAt: tMock,
	}

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: fields{DB: sx, User: user},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			u := tt.Fields.User

			mock.ExpectBegin()
			mock.ExpectPrepare(rgxQuery)
			mock.ExpectExec(rgxQuery).WithArgs(u.Password, u.Username, u.Email, u.Role, u.IsActive, u.CreatedAt, u.UpdatedAt).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			r := NewUserMysqlRepository(tt.Fields.DB)

			err := r.Add(tt.Fields.User)
			assert.NoError(t, err)
			assert.Equal(t, int64(1), user.ID)
		})
	}
}

func TestUpdate(t *testing.T) {
	tMock := time.Now()
	dbMock, mock, err := sqlmock.New()

	if err != nil {
		t.Error("Failed to mock db")
	}

	user := &models.UserTable{
		ID:        1,
		Username:  "sjuliper7",
		Email:     "sjuliper7@gmail.com",
		Role:      "user",
		IsActive:  1,
		CreatedAt: tMock,
		UpdatedAt: tMock,
	}

	sx := sqlx.NewDb(dbMock, "mockdb")

	sql := `UPDATE users SET username = ?, email = ?, role = ?, is_active = ? ,created_at = ?, updated_at = ? WHERE id=?`
	rgxQuery := regexp.QuoteMeta(sql)

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: fields{DB: sx, User: user},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			u := tt.Fields.User
			repo := NewUserMysqlRepository(tt.Fields.DB)

			mock.ExpectBegin()
			mock.ExpectPrepare(rgxQuery)
			mock.ExpectExec(rgxQuery).WithArgs(u.Username,
				u.Email,
				u.Role,
				u.IsActive,
				u.CreatedAt,
				u.UpdatedAt,
				u.ID).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			err := repo.Update(tt.Fields.User)
			assert.NoError(t, err)
		})

	}
}

func TestGet(t *testing.T) {
	tMock := time.Now()

	mockDb, mock, err := sqlmock.New()
	sx := sqlx.NewDb(mockDb, "mockbd")

	if err != nil {
		t.Error("Failed to mock db")
		return
	}

	sql := `SELECT id, username, email, role, is_active, created_at, updated_at FROM users where is_active = true and id = ?`
	rgxQuery := regexp.QuoteMeta(sql)

	tests := []struct {
		name   string
		fields fields
		want   models.UserTable
	}{
		{
			name:   "success-test",
			fields: fields{DB: sx},
			want: models.UserTable{
				ID:        1,
				Username:  "sjuliper",
				Email:     "simanjuntak.juliper@outlook.com",
				Role:      "admin",
				IsActive:  1,
				CreatedAt: tMock,
				UpdatedAt: tMock,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewUserMysqlRepository(tt.fields.DB)

			rows := sqlmock.NewRows([]string{"id", "username", "email", "role", "is_active", "created_at", "updated_at"}).
				AddRow(1, "sjuliper", "simanjuntak.juliper@outlook.com", "admin", 1, tMock, tMock)

			mock.ExpectPrepare(rgxQuery)
			mock.ExpectQuery(rgxQuery).WillReturnRows(rows)

			if got, _ := r.Get(1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlRepository.GetUser(userID) = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestDelet(t *testing.T) {
	dbMock, mock, err := sqlmock.New()

	sx := sqlx.NewDb(dbMock, "mockdb")

	if err != nil {
		t.Error("Failed to mock db")
	}

	sql := `UPDATE users SET is_active = false where id =?`
	rgxQuery := regexp.QuoteMeta(sql)

	user := &models.UserTable{
		ID: 1,
	}

	tests := []struct {
		Name   string
		Fields fields
	}{
		{
			Name:   "success-test",
			Fields: fields{DB: sx, User: user},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			u := tt.Fields.User
			repo := NewUserMysqlRepository(tt.Fields.DB)

			mock.ExpectBegin()
			prep := mock.ExpectPrepare(rgxQuery)
			prep.ExpectExec().WithArgs(u.ID).WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			_, err := repo.Delete(1)

			assert.NoError(t, err)
		})
	}
}

func TestGetByEmail(t *testing.T) {
	tMock := time.Now()
	dbMock, mock, err := sqlmock.New()

	sx := sqlx.NewDb(dbMock, "mockdb")

	if err != nil {
		t.Error("Failed to mock db")
	}

	sql := `SELECT id, password, username, email, role, is_active, created_at, updated_at FROM users where is_active = true and email = ?`
	rgxQuery := regexp.QuoteMeta(sql)

	tests := []struct {
		name   string
		fields fields
		want   models.UserTable
	}{
		{
			name:   "success-test",
			fields: fields{DB: sx},
			want: models.UserTable{
				ID:        1,
				Username:  "sjuliper",
				Email:     "simanjuntak.juliper@outlook.com",
				Role:      "admin",
				IsActive:  1,
				CreatedAt: tMock,
				UpdatedAt: tMock,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewUserMysqlRepository(tt.fields.DB)

			rows := sqlmock.NewRows([]string{"id", "username", "email", "role", "is_active", "created_at", "updated_at"}).
				AddRow(1, "sjuliper", "simanjuntak.juliper@outlook.com", "admin", 1, tMock, tMock)

			mock.ExpectPrepare(rgxQuery)
			mock.ExpectQuery(rgxQuery).WillReturnRows(rows)

			if got, _ := r.GetByEmail("simanjuntak.juliper@outlook.com"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlRepository.GetUser(userID) = %v, want %v", got, tt.want)
			}

		})
	}

}
