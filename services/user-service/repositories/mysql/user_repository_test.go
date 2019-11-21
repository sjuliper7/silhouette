package mysql

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/stretchr/testify/assert"
	"reflect"
	"regexp"
	"testing"
)

type fields struct {
	DB *sqlx.DB
}

func TestMysqlRepository_GetAlluser(t *testing.T) {

}

func TestMysqlRepository_AddUser(t *testing.T) {
	mockDB, mock, err := sqlmock.New()

	if err != nil {
		t.Error("Failed mock db")
	}

	sx := sqlx.NewDb(mockDB, "mockdb")

	sql := `INSERT INTO users(username, email, name, role) VALUES (?, ?, ?, ?)`

	rgxQuery := regexp.QuoteMeta(sql)

	t.Run("Test-1", func(t *testing.T) {
		user := models.UserTable{
			ID:       0,
			Username: "sjuliper8",
			Email:    "sjuliper8@gmail.com",
			Name:     "JuliperS",
			Role:     "admin",
		}

		mock.ExpectPrepare(rgxQuery)
		mock.ExpectExec(rgxQuery).WithArgs(user.Username, user.Email, user.Name, user.Role).
			WillReturnResult(sqlmock.NewResult(1, 1))

		r := &mysqlRepository{Conn: sx}

		err := r.AddUser(&user)
		assert.NoError(t, err)
		assert.Equal(t, uint64(1), user.ID)
	})
}

func TestMysqlRepository_GetUser(t *testing.T) {
	mockDb, mock, err := sqlmock.New()
	sx := sqlx.NewDb(mockDb, "mockbd")

	if err != nil {
		t.Error("Failed mock db")
		return
	}

	rgxQuery := "SELECT id, username, email, name, role FROM users where is_active = true and id = (.+)"

	tests := []struct {
		name   string
		fields fields
		want   models.UserTable
	}{
		{
			name:   "test-1",
			fields: fields{DB: sx},
			want: models.UserTable{
				ID:       1,
				Username: "sjuliper",
				Email:    "simanjuntak.juliper@outlook.com",
				Name:     "Juliper Simanjuntak",
				Role:     "admin",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mysqlRepository{
				Conn: tt.fields.DB,
			}

			rows := sqlmock.NewRows([]string{"id", "username", "email", "name", "role"}).
				AddRow(1, "sjuliper", "simanjuntak.juliper@outlook.com", "Juliper Simanjuntak", "admin")

			mock.ExpectPrepare(rgxQuery)
			mock.ExpectQuery(rgxQuery).WillReturnRows(rows)

			if got, _ := r.GetUser(1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlRepository.GetUser(userID) = %v, want %v", got, tt.want)
			}

		})
	}

}
