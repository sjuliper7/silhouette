package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/repository"
)

type userMysqlRepository struct {
	Conn *sqlx.DB
}

//NewMysqlRepository is to create new instance repository
func NewUserMysqlRepository(conn *sqlx.DB) repository.UserRepository {
	return &userMysqlRepository{conn}
}

func (repo *userMysqlRepository) GetAll() (users []models.UserTable, err error) {
	sql := `SELECT id, username, email, name, role, is_active, created_at, updated_at FROM users where is_active = true`
	rows, err := repo.Conn.Queryx(sql)
	if err != nil {
		logrus.Errorf("[user-repository][GetAll] error while querying: %v", err)
	}

	for rows.Next() {
		temp := models.UserTable{}
		var err = rows.StructScan(&temp)

		if err != nil {
			logrus.Errorf("[user-repository][GetAll] failed when getting result with params: %v", err)
			return nil, err
		}

		users = append(users, temp)

	}

	return users, nil
}

func (repo *userMysqlRepository) Add(user *models.UserTable) (err error) {
	tx, err := repo.Conn.Beginx()
	if err != nil {
		logrus.Errorf("[user-repository][Add] error while creating transactions: %v", err)
		return err
	}

	defer tx.Rollback()

	sql := `INSERT INTO users(username, email, name, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	stmt, err := tx.Preparex(sql)

	if err != nil {
		logrus.Errorf("[user-repository][Add] error when preparing query: %v", err)
		return err
	}

	result, err := stmt.Exec(user.Username, user.Email, user.Name, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		logrus.Errorf("[user-repository][Add] error when inserting values: %v", err)
		return err
	}

	var temp int64
	temp, err = result.LastInsertId()
	if err != nil {
		logrus.Errorf("[user-repository][Add] error when inserting values: %v", err)
		return err
	}

	user.ID = temp

	err = tx.Commit()
	if err != nil {
		logrus.Errorf("[user-repository][Add] error while commit transaction: %v", err)
		return nil
	}

	logrus.Infof("[user-repository][Add] successfully to add user: %v", user.ID)

	return nil
}

func (repo *userMysqlRepository) Get(userID int64) (user models.UserTable, err error) {

	sql := `SELECT id, username, email, name, role, is_active, created_at, updated_at FROM users where is_active = true and id = ?`
	stmt, err := repo.Conn.Preparex(sql)

	if err != nil {
		logrus.Errorf("[user-repository][Get] error when prepare the query: %v", err)
	}

	err = stmt.Get(&user, userID)
	if err != nil {
		logrus.Errorf("[user-repository][Get] error when getting the value: %v", err)
	}

	return user, nil
}

func (repo *userMysqlRepository) Update(user *models.UserTable) (err error) {
	logrus.Infof("[user-repository][Update] start update user: %v", user.ID)
	logrus.Infof("user : %v", user)
	tx, err := repo.Conn.Beginx()
	if err != nil {
		logrus.Errorf("[user-repository][Update] error while creating transactions %v", err)
		return err
	}

	defer tx.Rollback()

	sql := `UPDATE users SET username = ?, email = ?, name = ?, role = ?, is_active = ? ,created_at = ?, updated_at = ? WHERE id=?`

	stmt, err := tx.Preparex(sql)
	if err != nil {
		logrus.Errorf("[user-repository][Update] error when prepare the query: %v", err)
		return err
	}

	_, err = stmt.Exec(user.Username,
		user.Email,
		user.Name,
		user.Role,
		user.IsActive,
		user.CreatedAt,
		user.UpdatedAt,
		user.ID)

	if err != nil {
		logrus.Errorf("[user-repository][Update] error when exec the query with value: %v", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		logrus.Errorf("[user-repository][Update] error when commit transactions: %v", err)
		return nil
	}

	logrus.Infof("[user-repository][Update] successfully update user: %v", user.ID)

	return nil
}

func (repo *userMysqlRepository) Delete(userID int64) (deleted bool, err error) {
	tx, err := repo.Conn.Beginx()
	if err != nil {
		logrus.Errorf("[user-repository][Delete] error while creating transactions: %v", err)
		return false, err
	}

	defer tx.Rollback()

	sql := `UPDATE users SET is_active = false where id =?`

	stmt, err := tx.Preparex(sql)
	if err != nil {
		logrus.Errorf("[user-repository][Delete] error when prepare the query: %v", err)
		return false, err
	}
	_, err = stmt.Exec(userID)

	if err != nil {
		logrus.Errorf("[user-repository][Delete] error  exec the query with value: %v", err)
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		logrus.Errorf("[user-repository][Delete] error while commit transactions: %v", err)
		return false, err
	}

	logrus.Infof("successfully to deleted user: %v", userID)
	return true, nil
}
