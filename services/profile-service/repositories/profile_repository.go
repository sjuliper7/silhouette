package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/sjuliper7/silhouette/services/profile-service/models"
	"log"
)

type mysqlRepository struct {
	Conn *sqlx.DB
}

func NewMysqlRepository(conn *sqlx.DB) Repository {
	return &mysqlRepository{Conn: conn}
}

func (repo mysqlRepository) GetProfile(userID int64) (profile models.Profile, err error) {
	sql := `SELECT id, user_id, address, work_at, phone_number, gender FROM profiles WHERE user_id = ?`

	stmt, err := repo.Conn.Preparex(sql)

	if err != nil {
		log.Println("Error when prepare the query %+v", err)
	}

	err = stmt.Get(&profile, userID)
	if err != nil {
		log.Println("Error when getting the value %+v", err)
	}

	return profile, nil
}
