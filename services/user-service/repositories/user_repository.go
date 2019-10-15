package repositories

import (
	"github.com/jmoiron/sqlx"

	"log"

	"github.com/sjuliper7/silhouette/services/user-service/models"
)

type mysqlRepository struct {
	Conn *sqlx.DB
}

//NewMysqlRepository is to create new instance repository
func NewMysqlRepository(conn *sqlx.DB) Repository {
	return &mysqlRepository{conn}
}

func (repo mysqlRepository) GetAlluser() (users []models.User, err error) {
	sql := `SELECT id, username, email, name, role FROM users`
	rows, err := repo.Conn.Queryx(sql)
	if err != nil {
		log.Fatalln(err.Error)
	}

	for rows.Next() {
		var user models.User
		var err = rows.StructScan(&user)

		if err != nil {
			log.Println("Failed when getting result with params %+v ", err)
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
		log.Println("Error when inserting values ", err)
		return err
	}

	user.ID = uint64(temp)

	return nil
}

func (repo mysqlRepository) GetUser(userID int64) (user models.User, err error) {

	sql := `SELECT id, username, email, name, role FROM users where id = ?`

	stmt, err := repo.Conn.Preparex(sql)

	if err != nil {
		log.Println("Error when prepare the query %v", err)
	}

	err = stmt.Get(&user, userID)

	if err != nil {
		log.Println("Error when getting the value %v", err)
	}

	return user, nil
}
