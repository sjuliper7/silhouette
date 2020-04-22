package mysql

import (
	"github.com/jmoiron/sqlx"
)

type fields struct {
	DB *sqlx.DB
}

// func TestMysqlRepository_GetAll(t *testing.T) {

// 	tMock := time.Now()

// 	mockDb, mock, err := sqlmock.New()

// 	if err != nil {
// 		t.Error("Failed mock db")
// 	}

// 	sx := sqlx.NewDb(mockDb, "mockdb")

// 	sql := `SELECT id, username, email, name, role, is_active, created_at, updated_at FROM users where is_active = true`
// 	rgxQuery := regexp.QuoteMeta(sql)

// 	test := []struct {
// 		Name   string
// 		Fields fields
// 		Want   []models.UserTable
// 	}{
// 		{
// 			Name:   "Test-1",
// 			Fields: fields{DB: sx},
// 			Want: []models.UserTable{
// 				models.UserTable{
// 					ID:        1,
// 					Username:  "sjuliper7",
// 					Email:     "sjuliper7@gmail.com",
// 					Name:      "Juliper Simanjuntak",
// 					Role:      "user",
// 					IsActive:  1,
// 					CreatedAt: tMock,
// 					UpdatedAt: tMock,
// 				},
// 				models.UserTable{
// 					ID:        2,
// 					Username:  "yesica",
// 					Email:     "yesica@gmail.com",
// 					Name:      "Yesica Tampubolon",
// 					Role:      "user",
// 					IsActive:  1,
// 					CreatedAt: tMock,
// 					UpdatedAt: tMock,
// 				},
// 			},
// 		},
// 	}

// 	for _, tt := range test {
// 		repo := userMysqlRepository{Conn: tt.Fields.DB}

// 		rows := sqlmock.NewRows([]string{"id", "username", "email", "name", "role", "is_active", "created_at", "updated_at"}).
// 			AddRow(1, "sjuliper7", "sjuliper7@gmail.com", "Juliper Simanjuntak", "user", 1, tMock, tMock).
// 			AddRow(2, "yesica", "yesica@gmail.com", "Yesica Tampubolon", "user", 1, tMock, tMock)

// 		mock.ExpectQuery(rgxQuery).WillReturnRows(rows)

// 		resultMock, err := repo.GetAll()
// 		if err != nil {
// 			t.Errorf("Error When call GetAllUser")
// 		}

// 		if !reflect.DeepEqual(resultMock, tt.Want) {
// 			t.Errorf("workRepo.GetAllWork() = %v, want %v", resultMock, tt.Want)
// 		}
// 	}

// }

// func TestMysqlRepository_Add(t *testing.T) {
// 	mockDB, mock, err := sqlmock.New()

// 	if err != nil {
// 		t.Error("Failed mock db")
// 	}

// 	sx := sqlx.NewDb(mockDB, "mockdb")

// 	sql := `INSERT INTO users(username, email, name, role) VALUES (?, ?, ?, ?)`
// 	rgxQuery := regexp.QuoteMeta(sql)

// 	t.Run("Test-1", func(t *testing.T) {
// 		user := models.UserTable{
// 			ID:       0,
// 			Username: "sjuliper8",
// 			Email:    "sjuliper8@gmail.com",
// 			Name:     "JuliperS",
// 			Role:     "admin",
// 		}

// 		mock.ExpectPrepare(rgxQuery)
// 		mock.ExpectExec(rgxQuery).WithArgs(user.Username, user.Email, user.Name, user.Role).
// 			WillReturnResult(sqlmock.NewResult(1, 1))

// 		r := &userMysqlRepository{Conn: sx}

// 		err := r.Add(&user)
// 		assert.NoError(t, err)
// 		assert.Equal(t, uint64(1), user.ID)
// 	})
// }

// func TestUserMysqlRepository_Update(t *testing.T) {
// 	tMock := time.Now()
// 	dbMock, mock, err := sqlmock.New()

// 	if err != nil {
// 		t.Error("Failed mock db")
// 	}

// 	user := models.UserTable{
// 		ID:        1,
// 		Username:  "sjuliper7",
// 		Email:     "sjuliper7@gmail.com",
// 		Name:      "Juliper Simanjuntak",
// 		Role:      "user",
// 		IsActive:  1,
// 		CreatedAt: tMock,
// 		UpdatedAt: tMock,
// 	}

// 	sx := sqlx.NewDb(dbMock, "mockdb")

// 	sql := `UPDATE users SET username = ?, email = ?, name = ?, role = ?, is_active = ? ,created_at = ?, updated_at = ? WHERE id=?`
// 	rgxQuery := regexp.QuoteMeta(sql)

// 	t.Run("Test-1", func(t *testing.T) {
// 		repo := &userMysqlRepository{Conn: sx}

// 		mock.ExpectPrepare(rgxQuery)
// 		mock.ExpectExec(rgxQuery).WithArgs(user.Username,
// 			user.Email,
// 			user.Name,
// 			user.Role,
// 			user.IsActive,
// 			user.CreatedAt,
// 			user.UpdatedAt,
// 			user.ID).
// 			WillReturnResult(sqlmock.NewResult(1, 1))

// 		err := repo.Update(&user)
// 		assert.NoError(t, err)
// 	})

// }

// func TestMysqlRepository_Get(t *testing.T) {
// 	tMock := time.Now()
// 	mockDb, mock, err := sqlmock.New()
// 	sx := sqlx.NewDb(mockDb, "mockbd")

// 	if err != nil {
// 		t.Error("Failed mock db")
// 		return
// 	}

// 	rgxQuery := "SELECT id, username, email, name, role, is_active, created_at, updated_at FROM users where is_active = true and id = (.+)"

// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   models.UserTable
// 	}{
// 		{
// 			name:   "test-1",
// 			fields: fields{DB: sx},
// 			want: models.UserTable{
// 				ID:        1,
// 				Username:  "sjuliper",
// 				Email:     "simanjuntak.juliper@outlook.com",
// 				Name:      "Juliper Simanjuntak",
// 				Role:      "admin",
// 				IsActive:  1,
// 				CreatedAt: tMock,
// 				UpdatedAt: tMock,
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := &userMysqlRepository{
// 				Conn: tt.fields.DB,
// 			}

// 			rows := sqlmock.NewRows([]string{"id", "username", "email", "name", "role", "is_active", "created_at", "updated_at"}).
// 				AddRow(1, "sjuliper", "simanjuntak.juliper@outlook.com", "Juliper Simanjuntak", "admin", 1, tMock, tMock)

// 			mock.ExpectPrepare(rgxQuery)
// 			mock.ExpectQuery(rgxQuery).WillReturnRows(rows)

// 			if got, _ := r.Get(1); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("mysqlRepository.GetUser(userID) = %v, want %v", got, tt.want)
// 			}

// 		})
// 	}

// }

// func TestUserMysqlRepository_DeleteUser(t *testing.T) {
// 	dbMock, mock, err := sqlmock.New()

// 	sx := sqlx.NewDb(dbMock, "mockdb")

// 	if err != nil {
// 		t.Error("Failed mock db")
// 	}

// 	sql := `UPDATE users SET is_active = false where id =?`
// 	rgxQuery := regexp.QuoteMeta(sql)

// 	t.Run("Test-delete-1", func(t *testing.T) {
// 		repo := userMysqlRepository{Conn: sx}

// 		prep := mock.ExpectPrepare(rgxQuery)
// 		prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

// 		_, err := repo.Delete(1)

// 		assert.NoError(t, err)
// 	})
// }
