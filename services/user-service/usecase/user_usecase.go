package usecase

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/repositories"
	"time"
)

type userUsecase struct {
	userRepo    repositories.UserRepository
	profileRepo repositories.ProfileRepository
	kafkaRepo   repositories.KafkaRepository
}

func NewUserUsecase(userRepo repositories.UserRepository, profileRepo repositories.ProfileRepository, kafkaRepo repositories.KafkaRepository) UserUsecase {
	return userUsecase{userRepo, profileRepo, kafkaRepo}
}

func (uc userUsecase) GetAll() (users []models.User, err error) {
	usersTable, err := uc.userRepo.GetAll()

	if err != nil {
		logrus.Errorf("Failed when call [repositories][GetAlluser] ", err)
		return nil, err
	}

	for _, u := range usersTable {
		user := models.User{}
		user.ID = u.ID
		user.Username = u.Username
		user.Email = u.Email
		user.Name = u.Name
		user.Role = u.Role
		user.CreatedAt = u.CreatedAt
		user.UpdatedAt = u.UpdatedAt

		users = append(users, user)
	}

	users, err = uc.fillProfileDetails(users)
	if err != nil {
		logrus.Errorf("Failed when call [usecase][fillProfileDetails] ", err)
		return nil, err
	}

	return users, err
}

func (uc userUsecase) Add(user *models.User) (err error) {

	userTable := models.UserTable{
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		IsActive:  1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	
	err = uc.userRepo.Add(&userTable)
	if err != nil {
		logrus.Errorf("[usecase][AddUser] Error when calling repository to save")
	}

	profile := models.Profile{
		Address:     user.Profile.Address,
		UserID:      userTable.ID,
		WorkAt:      user.Profile.WorkAt,
		PhoneNumber: user.Profile.PhoneNumber,
		Gender:      user.Profile.Gender,
	}

	jsonData, err := json.Marshal(profile)

	if err != nil {
		logrus.Errorf("[usecase][AddUser] error when marshall data ", err)
		return err
	}

	err = uc.kafkaRepo.PublishMessage("registration-finish", jsonData)

	if err != nil {
		logrus.Errorf("[usecase][AddUser] error when publish message ", err)
		return err
	}

	return nil
}

func (uc userUsecase) Get(userID int64) (user models.User, err error) {
	ut := models.UserTable{}
	ut, err = uc.userRepo.Get(userID)

	if err != nil {
		logrus.Errorf("[usecase][GetUser] Error when calling repository to get user")
		return user, err
	}

	user.Role = ut.Role
	user.Username = ut.Username
	user.Email = ut.Email
	user.Name = ut.Name
	user.ID = ut.ID
	user.CreatedAt = ut.CreatedAt
	user.UpdatedAt = ut.UpdatedAt

	var profile models.Profile = models.Profile{}
	profile, err = uc.profileRepo.Get(userID)

	if err != nil {
		logrus.Errorf("[usecase][GetUser] Error when calling profile repository to get profile")
	}

	user.Profile = profile

	return user, nil
}

func (uc userUsecase) Update(us models.User) (user models.User, err error) {

	userTable, err := uc.userRepo.Get(int64(us.ID))

	user.Role = us.Role
	user.Email = us.Email
	user.Username = us.Username
	user.Name = us.Name
	user.UpdatedAt = us.UpdatedAt

	err = uc.userRepo.Update(&userTable)

	if err != nil {
		logrus.Errorf("[usecase][UpdateUser] Error when calling repository to update user")
		return user, err
	}

	profile := models.Profile{
		ID:          0,
		UserID:      userTable.ID,
		Address:     us.Profile.Address,
		WorkAt:      us.Profile.WorkAt,
		PhoneNumber: us.Profile.PhoneNumber,
		Gender:      us.Profile.Gender,
	}

	jsonData, err := json.Marshal(profile)

	if err != nil {
		logrus.Errorf("[usecase][AddUser] error when marshall data ", err)
		return user, err
	}

	err = uc.kafkaRepo.PublishMessage("update-finish", jsonData)

	if err != nil {
		logrus.Errorf("[usecase][AddUser] error when publish message ", err)
		return user, err
	}

	return user, nil
}

func (uc userUsecase) Delete(userID int64) (deleted bool, err error) {
	deleted, err = uc.userRepo.Delete(userID)

	if err != nil {
		logrus.Errorf("[usecase][Delete] Error when calling repository to delete user")
		return false, err
	}

	return deleted, nil
}

func (uc *userUsecase) fillProfileDetails(users []models.User) ([]models.User, error) {
	for i, _ := range users {
		profile, err := uc.profileRepo.Get(int64(users[i].ID))

		if err != nil {
			logrus.Errorf("[usecase][GetAllUser] Error when calling profile repository to get profile")
			return users, err
		}

		users[i].Profile = profile
	}
	return users, nil
}
