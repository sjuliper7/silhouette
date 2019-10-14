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
