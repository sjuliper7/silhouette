package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/repositories"
)

type mysqlRepository struct {
	Conn *sqlx.DB
}

//NewMysqlRepository is to create new instance repository
func NewMysqlRepository(conn *sqlx.DB) repositories.UserRepository {
	return &mysqlRepository{conn}
}

func (repo mysqlRepository) GetAlluser() (users []models.User, err error) {
	sql := `SELECT id, username, email, name, role FROM users where is_active = true`
	rows, err := repo.Conn.Queryx(sql)
	if err != nil {
		logrus.Fatalln(err.Error)
	}

	for rows.Next() {
		var user models.User
		var err = rows.StructScan(&user)

		if err != nil {
			logrus.Println("Failed when getting result with params %+v ", err)
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil
}

func (repo mysqlRepository) AddUser(user *models.User) (err error) {
	sql := `INSERT INTO users(username, email, name, role) VALUES(?, ?, ?, ?)`

	result := repo.Conn.MustExec(sql, user.Username, user.Email, user.Name, user.Role)
	var temp int64
	temp, err = result.LastInsertId()
	if err != nil {
		logrus.Println("Error when inserting values %+v", err)
		return err
	}

	user.ID = uint64(temp)

	return nil
}

func (repo mysqlRepository) GetUser(userID int64) (user models.User, err error) {

	sql := `SELECT id, username, email, name, role FROM users where is_active = true and id = ?`
	stmt, err := repo.Conn.Preparex(sql)

	if err != nil {
		logrus.Println("Error when prepare the query %+v", err)
	}

	err = stmt.Get(&user, userID)
	if err != nil {
		logrus.Println("Error when getting the value %+v", err)
	}

	return user, nil
}

func (repo mysqlRepository) UpdateUser(user *models.User) (err error) {
	sql := `UPDATE users SET username = ?, email = ?, name = ?, role = ? WHERE id=?`

	stmt, err := repo.Conn.Preparex(sql)
	if err != nil {
		logrus.Println("Error when prepare the query %+v", err)
		return err
	}
	_, err = stmt.Exec(user.Username, user.Email, user.Name, user.Role, user.ID)

	if err != nil {
		logrus.Println("Error when exec the query with value %+v", err)
		return err
	}

	return nil

}

func (repo mysqlRepository) DeleteUser(userID int64) (deleted bool, err error) {
	sql := `UPDATE users SET is_active = false where id =?`

	stmt, err := repo.Conn.Preparex(sql)
	if err != nil {
		logrus.Println("Error when prepare the query %+v", err)
		return false, err
	}
	_, err = stmt.Exec(userID)

	if err != nil {
		logrus.Println("Error when exec the query with value %+v", err)
		return false, err
	}

	return true, nil
}
