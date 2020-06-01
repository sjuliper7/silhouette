package usecase

import (
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/constans"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/repository"
)

type userUsecase struct {
	userRepo    repository.UserRepository
	profileRepo repository.ProfileRepository
	kafkaRepo   repository.KafkaRepository
}

//NewUserUsecase ...
func NewUserUsecase(userRepo repository.UserRepository, profileRepo repository.ProfileRepository, kafkaRepo repository.KafkaRepository) UserUsecase {
	return userUsecase{userRepo, profileRepo, kafkaRepo}
}

func (uc userUsecase) GetAll() (users []models.User, err error) {
	usersTable, err := uc.userRepo.GetAll()

	if err != nil {
		logrus.Errorf("[usecase][GetAll] failed when call [repository][GetAll] %v", err)
		return nil, err
	}

	for _, u := range usersTable {
		user := models.User{}
		user.ID = u.ID
		user.Username = u.Username
		user.Email = u.Email
		// user.Name = u.Name
		user.Role = u.Role
		user.CreatedAt = u.CreatedAt.String()
		user.UpdatedAt = u.UpdatedAt.String()

		users = append(users, user)
	}

	users, err = uc.fillProfileDetails(users)
	if err != nil {
		logrus.Errorf("[usecase][GetAll] failed when call [usecase][fillProfileDetails] %v", err)
		return nil, err
	}

	return users, err
}

func (uc userUsecase) Add(user *models.User) (err error) {
	userTable := models.UserTable{
		Username: user.Username,
		Email:    user.Email,
		// Name:      user.Name,
		Role:      user.Role,
		IsActive:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = uc.userRepo.Add(&userTable)
	if err != nil {
		logrus.Errorf("[usecase][Add] failed when call [repository][Add]: %v", err)
		return err
	}

	user.ID = userTable.ID

	message := map[string]interface{}{
		"user_id":       userTable.ID,
		"address":       user.Profile.Address,
		"work_at":       user.Profile.WorkAt,
		"phone_number":  user.Profile.PhoneNumber,
		"gender":        user.Profile.Gender,
		"email":         user.Email,
		"type":          "email",
		"name":          user.Profile.Name,
		"date_of_birth": user.Profile.DateOfBirth,
	}

	jsonData, err := json.Marshal(message)

	if err != nil {
		logrus.Errorf("[usecase][Add] failed when marshall: %v", err)
		return err
	}

	err = uc.kafkaRepo.PublishMessage(string(constans.TopicUserRegistration), jsonData)

	if err != nil {
		logrus.Errorf("[usecase][AddUser] error when publish message: %v", err)
		return err
	}

	return nil
}

func (uc userUsecase) Get(userID int64) (user models.User, err error) {
	ut := models.UserTable{}
	ut, err = uc.userRepo.Get(userID)

	if err != nil {
		logrus.Errorf("[usecase][Get] Error when calling [repository][Get]: %v", err)
		return user, err
	}

	user.Role = ut.Role
	user.Username = ut.Username
	user.Email = ut.Email
	// user.Name = ut.Name
	user.ID = ut.ID
	user.CreatedAt = ut.CreatedAt.String()
	user.UpdatedAt = ut.UpdatedAt.String()

	var profile models.Profile = models.Profile{}
	profile, err = uc.profileRepo.Get(userID)

	if err != nil {
		logrus.Errorf("[usecase][Get] Error when calling [repository][Get]: %v", err)
		return user, err
	}

	user.Profile = profile

	return user, nil
}

func (uc userUsecase) Update(us models.User) (user models.User, err error) {

	logrus.Infof("params: %v", us)

	userTable, err := uc.userRepo.Get(int64(us.ID))

	if err != nil {
		logrus.Errorf("[usecase][Update] Error when calling [repository][Get]: %v", err)
		return user, err
	}

	userTable.Role = us.Role
	userTable.Email = us.Email
	userTable.Username = us.Username
	userTable.UpdatedAt = time.Now()

	err = uc.userRepo.Update(&userTable)

	if err != nil {
		logrus.Errorf("[usecase][Update] Error when calling [repository][Update]: %v", err)
		return us, err
	}

	message := map[string]interface{}{
		"user_id":       userTable.ID,
		"address":       us.Profile.Address,
		"work_at":       us.Profile.WorkAt,
		"phone_number":  us.Profile.PhoneNumber,
		"gender":        us.Profile.Gender,
		"email":         us.Email,
		"type":          "email",
		"name":          us.Profile.Name,
		"date_of_birth": us.Profile.DateOfBirth,
	}

	jsonData, err := json.Marshal(message)

	if err != nil {
		logrus.Errorf("[usecase][Update] error when marshall data: %v , %v", us, err)
		return us, err
	}

	err = uc.kafkaRepo.PublishMessage(string(constans.TopicUserUpdated), jsonData)

	if err != nil {
		logrus.Errorf("[usecase][Update] error when publish message: %v ,%v", us, err)
		return user, err
	}

	user, err = uc.Get(userTable.ID)
	if err != nil {
		logrus.Errorf("[usecase][Update] Error when calling [repository][Get]: %v", err)
		return user, err
	}

	return user, nil
}

func (uc userUsecase) Delete(userID int64) (deleted bool, err error) {
	deleted, err = uc.userRepo.Delete(userID)

	if err != nil {
		logrus.Errorf("[usecase][Delete] Error when calling [repository][delete]: %v", err)
		return false, err
	}

	message := map[string]interface{}{
		"user_id": userID,
	}

	jsonData, err := json.Marshal(message)

	if err != nil {
		logrus.Errorf("[usecase][Delete] error when marshall data: %v , %v", userID, err)
		return false, err
	}

	err = uc.kafkaRepo.PublishMessage(string(constans.TopicUserDeleted), jsonData)

	if err != nil {
		logrus.Errorf("[usecase][Delete] error when publish message: %v ,%v", userID, err)
		return false, err
	}

	return deleted, nil
}

func (uc *userUsecase) fillProfileDetails(users []models.User) ([]models.User, error) {
	for i, _ := range users {
		profile, err := uc.profileRepo.Get(int64(users[i].ID))

		if err != nil {
			logrus.Errorf("[usecase][fillProfileDetails] Error when calling [repository-profile][Get]: ")
			return users, err
		}

		users[i].Profile = profile
	}
	return users, nil
}
