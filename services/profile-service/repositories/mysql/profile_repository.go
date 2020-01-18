package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/profile-service/models"
	"github.com/sjuliper7/silhouette/services/profile-service/repositories"
)

type mysqlRepository struct {
	Conn *sqlx.DB
}

func NewMysqlProfileRepository(conn *sqlx.DB) repositories.Repository {
	return &mysqlRepository{Conn: conn}
}

func (profileRepository *mysqlRepository) Get(userID int64) (profile models.ProfileTable, err error) {
	sql := `SELECT id, user_id, address, work_at, phone_number, gender, created_at, updated_at FROM profiles WHERE user_id = ?`

	stmt, err := profileRepository.Conn.Preparex(sql)

	if err != nil {
		logrus.Errorf("[profileRepository][Get] error when prepare the query %v", err)
	}

	err = stmt.Get(&profile, userID)
	if err != nil {
		logrus.Errorf("[profileRepository][Get] error when getting value, %v", err)
	}

	return profile, nil
}

func (profileRepository *mysqlRepository) Add(profile *models.ProfileTable) (err error) {
	sql := `INSERT INTO profiles(user_id, address, work_at, phone_number, gender, is_active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	stmt, err := profileRepository.Conn.Preparex(sql)

	if err != nil {
		logrus.Errorf("[profileRepository][Add] error when preparing query, %v", err)
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
		logrus.Errorf("[profileRepository][Add] error when inserting values, %v", err)
		return err
	}

	var temp int64
	temp, err = result.LastInsertId()
	if err != nil {
		logrus.Errorf("[profileRepository][Add] error when inserting values, %v ", err)
		return err
	}

	profile.ID = uint64(temp)

	return nil
}

func (profileRepository *mysqlRepository) Update(profile *models.ProfileTable) (err error) {
	sql := `UPDATE profiles SET user_id = ?, address = ?, work_at = ?, phone_number = ?, gender = ?, is_active = ? ,created_at = ?, updated_at = ? WHERE id=?`

	stmt, err := profileRepository.Conn.Preparex(sql)
	if err != nil {
		logrus.Errorf("[profileRepository][Update] error when prepare the query, %v", err)
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
		logrus.Errorf("[profileRepository][Update] error when exec the query with value, %v", err)
		return err
	}

	return nil
}

func (profileRepository *mysqlRepository) Delete(profileID int64)(deleted bool, err error) {
	sql := `UPDATE profiles SET is_active = false where id =?`

	stmt, err := profileRepository.Conn.Preparex(sql)
	if err != nil {
		logrus.Errorf("[profileRepository][Delete] error when prepare the query, %v", err)
		return false, err
	}
	_, err = stmt.Exec(profileID)

	if err != nil {
		logrus.Errorf("[profileRepository][Delete] error when exec the query with value, %v ", err)
		return false, err
	}

	return true, nil
}