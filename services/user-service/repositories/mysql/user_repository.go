package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/repositories"
)

type userMysqlRepository struct {
	Conn *sqlx.DB
}

//NewMysqlRepository is to create new instance repository
func NewUserMysqlRepository(conn *sqlx.DB) repositories.UserRepository {
	return &userMysqlRepository{conn}
}

func (repo userMysqlRepository) GetAlluser() (users []models.UserTable, err error) {
	sql := `SELECT id, username, email, name, role, is_active, created_at, updated_at FROM users where is_active = true`
	rows, err := repo.Conn.Queryx(sql)
	if err != nil {
		logrus.Fatalln("Error ", err.Error())
	}

	for rows.Next() {
		temp := models.UserTable{}
		var err = rows.StructScan(&temp)

		if err != nil {
			logrus.Println("Here")
			logrus.Println("Failed when getting result with params ", err)
			return nil, err
		}

		users = append(users, temp)

	}

	return users, nil
}

func (repo userMysqlRepository) AddUser(user *models.UserTable) (err error) {
	sql := `INSERT INTO users(username, email, name, role) VALUES (?, ?, ?, ?)`
	stmt, err := repo.Conn.Preparex(sql)

	if err != nil {
		logrus.Println("Error when preparing query", err)
		return err
	}

	result, err := stmt.Exec(user.Username, user.Email, user.Name, user.Role)
	if err != nil {
		logrus.Println("Error when inserting values :", err)
		return err
	}

	var temp int64
	temp, err = result.LastInsertId()
	if err != nil {
		logrus.Println("Error when inserting values ", err)
		return err
	}

	user.ID = uint64(temp)

	return nil
}

func (repo userMysqlRepository) GetUser(userID int64) (user models.UserTable, err error) {

	sql := `SELECT id, username, email, name, role, is_active, created_at, updated_at FROM users where is_active = true and id = ?`
	stmt, err := repo.Conn.Preparex(sql)

	if err != nil {
		logrus.Println("Error when prepare the query ", err)
	}

	err = stmt.Get(&user, userID)
	if err != nil {
		logrus.Println("Error when getting the value ", err)
	}

	return user, nil
}

func (repo userMysqlRepository) UpdateUser(user *models.UserTable) (err error) {
	sql := `UPDATE users SET username = ?, email = ?, name = ?, role = ?, created_at = ?, updated_at = ? WHERE id=?`

	stmt, err := repo.Conn.Preparex(sql)
	if err != nil {
		logrus.Println("Error when prepare the query ", err)
		return err
	}
	_, err = stmt.Exec(user.Username, user.Email, user.Name, user.Role, user.ID)

	if err != nil {
		logrus.Println("Error when exec the query with value ", err)
		return err
	}

	return nil

}

func (repo userMysqlRepository) DeleteUser(userID int64) (deleted bool, err error) {
	sql := `UPDATE users SET is_active = false where id =?`

	stmt, err := repo.Conn.Preparex(sql)
	if err != nil {
		logrus.Println("Error when prepare the query ", err)
		return false, err
	}
	_, err = stmt.Exec(userID)

	if err != nil {
		logrus.Println("Error when exec the query with value ", err)
		return false, err
	}

	return true, nil
}
