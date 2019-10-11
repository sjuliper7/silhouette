package repositories

import (
	"database/sql"

	"log"

	"github.com/sjuliper7/silhouette/services/user-service/repositories/queries"

	"github.com/sjuliper7/silhouette/services/user-service/models"
)

type mysqlRepository struct {
	Conn *sql.DB
}

//NewMysqlRepository is to create new instance repository
func NewMysqlRepository(conn *sql.DB) Repository {
	return &mysqlRepository{conn}
}

func (repo mysqlRepository) GetAlluser() []models.User {
	sql := queries.Q(queries.QueryGetAllUser, map[string]string{})
	rows, err := repo.Conn.Query(sql)
	if err != nil {
		log.Fatalln(err.Error)
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user = models.User{}
		var err = rows.Scan(&user.ID, &user.Name, &user.LastName)

		if err != nil {
			log.Fatalln(err.Error)
		}

		users = append(users, user)

	}

	return users
}
