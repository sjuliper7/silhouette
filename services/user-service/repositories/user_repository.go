package repositories

import (
	"database/sql"

	"log"

	"github.com/sjuliper7/silhouette/services/user-service/models"
)

type mysqlRepository struct {
	Conn *sql.DB
}

//NewMysqlRepository is to create new instance repository
func NewMysqlRepository(conn *sql.DB) Repository {
	return &mysqlRepository{conn}
}

func (repo mysqlRepository) GetAlluser() (users []models.User, err error) {
	sql := `SELECT id, name, last_name FROM users`
	rows, err := repo.Conn.Query(sql)
	if err != nil {
		log.Fatalln(err.Error)
	}

	defer rows.Close()

	for rows.Next() {
		var user = models.User{}
		var err = rows.Scan(&user.ID, &user.Name, &user.LastName)

		if err != nil {
			log.Println("Failed when getting result with params %v ", err)
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil
}
