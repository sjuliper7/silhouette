package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/profile-service/helper"
	"github.com/sjuliper7/silhouette/services/profile-service/models"
	"github.com/sjuliper7/silhouette/services/profile-service/repository"
)

type mysqlRepository struct {
	Conn *sqlx.DB
}

//NewMysqlProfileRepository ...
func NewMysqlProfileRepository(conn *sqlx.DB) repository.Repository {
	return &mysqlRepository{Conn: conn}
}

func (profileRepository *mysqlRepository) Get(userID int64) (profile models.ProfileTable, err error) {
	sql := `SELECT id, user_id, address, work_at, phone_number, gender, created_at, updated_at FROM profiles WHERE user_id = ? and is_active = 1`

	rows, err := profileRepository.Conn.Queryx(sql, userID)

	if err != nil {
		logrus.Errorf("[profileRepository][Get] error when queries: %v", err)
		return profile, nil
	}

	defer rows.Close()

	for rows.Next() {
		temp := models.PortfolioTableScanner{}

		err := rows.StructScan(&temp)
		if err != nil {
			logrus.Errorf("[profileRepository][Get] error when scanning values: %v", err)
			return profile, nil
		}

		profile.ID = temp.ID.Int64
		profile.IsActive = temp.IsActive.Bool
		profile.Gender = temp.Gender.String
		profile.PhoneNumber = temp.PhoneNumber.String
		profile.WorkAt = temp.WorkAt.String
		profile.Address = temp.Address.String
		profile.UserId = temp.UserId.Int64
		profile.CreatedAt = helper.ParseStringToTime(temp.CreatedAt.String)
		profile.UpdatedAt = helper.ParseStringToTime(temp.UpdatedAt.String)
	}

	return profile, nil
}

func (profileRepository *mysqlRepository) Add(profile *models.ProfileTable) (err error) {

	tx, err := profileRepository.Conn.Beginx()
	if err != nil {
		logrus.Errorf("[profileRepository][Add] error creating transactions: %v", err)
		return err
	}

	defer tx.Rollback()

	sql := `INSERT INTO profiles(user_id, address, work_at, phone_number, gender, is_active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	stmt, err := tx.Preparex(sql)

	if err != nil {
		logrus.Errorf("[profileRepository][Add] error when preparing query: %v", err)
		return err
	}

	result, err := stmt.Exec(
		profile.UserId,
		profile.Address,
		profile.WorkAt,
		profile.PhoneNumber,
		profile.Gender,
		profile.IsActive,
		profile.CreatedAt,
		profile.UpdatedAt,
	)

	if err != nil {
		logrus.Errorf("[profileRepository][Add] error when inserting values: %v", err)
		return err
	}

	var temp int64
	temp, err = result.LastInsertId()
	if err != nil {
		logrus.Errorf("[profileRepository][Add] error when inserting values: %v ", err)
		return err
	}

	profile.ID = temp

	err = tx.Commit()
	if err != nil {
		logrus.Errorf("[profileRepository][Add] error when commit transaction: %v", err)
	}

	logrus.Infof("[profileRepository][Add] successfully to add: %v", profile.ID)

	return nil
}

func (profileRepository *mysqlRepository) Update(profile *models.ProfileTable) (err error) {
	logrus.Infof("[profileRepository][Update] start to update : %v", profile.ID)

	tx, err := profileRepository.Conn.Beginx()
	if err != nil {
		logrus.Errorf("[profileRepository][Update] error when creating transaction: %v", err)
		return nil
	}

	defer tx.Rollback()

	sql := `UPDATE profiles SET user_id = ?, address = ?, work_at = ?, phone_number = ?, gender = ?, is_active = ? ,created_at = ?, updated_at = ? WHERE id=?`

	stmt, err := tx.Preparex(sql)
	if err != nil {
		logrus.Errorf("[profileRepository][Update] error when prepare the query: %v", err)
		return err
	}

	_, err = stmt.Exec(profile.UserId,
		profile.Address,
		profile.WorkAt,
		profile.PhoneNumber,
		profile.Gender,
		profile.IsActive,
		profile.CreatedAt,
		profile.UpdatedAt,
		profile.ID,
	)

	if err != nil {
		logrus.Errorf("[profileRepository][Update] error when exec the query with value: %v", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		logrus.Errorf("[profileRepository][Update] error when commit transaction: %v", err)
		return nil
	}

	logrus.Infof("[profileRepository][Update] successfully to updated: %v", profile.ID)

	return nil
}

func (profileRepository *mysqlRepository) Delete(userID int64) (deleted bool, err error) {
	tx, err := profileRepository.Conn.Beginx()
	if err != nil {
		logrus.Errorf("[profileRepository][Delete] error when creating transaction: %v", err)
		return false, err
	}

	defer tx.Rollback()

	sql := `UPDATE profiles SET is_active = false where user_id =?`

	stmt, err := tx.Preparex(sql)
	if err != nil {
		logrus.Errorf("[profileRepository][Delete] error when prepare the query: %v", err)
		return false, err
	}
	_, err = stmt.Exec(userID)

	if err != nil {
		logrus.Errorf("[profileRepository][Delete] error when exec the query with value: %v ", err)
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		logrus.Errorf("[profileRepository][Delete] error when commit transaction: %v", err)
		return false, err
	}

	logrus.Infof("[profileRepository][Delete] successfully to deleted profile user id: %v", userID)

	return true, nil
}
