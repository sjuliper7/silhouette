package usecase

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/repositories"
)

type userUsecase struct {
	userRepo    repositories.UserRepository
	profileRepo repositories.ProfileRepository
}

func NewUserUsecase(userRepo repositories.UserRepository, profileRepo repositories.ProfileRepository) UserUsecase {
	return userUsecase{userRepo, profileRepo}
}

func (uc userUsecase) GetAllUser() (users []models.User, err error) {
	usersTable, err := uc.userRepo.GetAllUser()

	if err != nil {
		logrus.Println("Failed when call [repositories][GetAlluser] ", err)
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
		logrus.Println("Failed when call [usecase][fillProfileDetails] ", err)
		return nil, err
	}

	return users, err
}

func (uc userUsecase) AddUser(user *models.UserTable) (err error) {

	err = uc.userRepo.AddUser(user)
	if err != nil {
		logrus.Println("[usecase][AddUser] Error when calling repository to save")
	}

	return nil
}

func (uc userUsecase) GetUser(userID int64) (user models.User, err error) {
	ut := models.UserTable{}
	ut, err = uc.userRepo.GetUser(userID)

	if err != nil {
		logrus.Println("[usecase][GetUser] Error when calling repository to get user")
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
	profile, err = uc.profileRepo.GetProfile(userID)

	if err != nil {
		logrus.Println("[usecase][GetUser] Error when calling profile repository to get profile")
	}

	user.Profile = profile

	return user, nil
}

func (uc userUsecase) UpdateUser(us models.UserTable) (user models.UserTable, err error) {

	user, err = uc.userRepo.GetUser(int64(us.ID))

	user.Role = us.Role
	user.Email = us.Email
	user.Username = us.Username
	user.Name = us.Name
	user.UpdatedAt = us.UpdatedAt


	err = uc.userRepo.UpdateUser(&user)

	if err != nil {
		logrus.Println("[usecase][UpdateUser] Error when calling repository to update user")
		return user,err
	}

	return user,nil
}

func (uc userUsecase) DeleteUser(userID int64) (deleted bool, err error) {
	deleted, err = uc.userRepo.DeleteUser(userID)

	if err != nil {
		logrus.Println("[usecase][Delete] Error when calling repository to delete user")
		return false, err
	}

	return deleted, nil
}

func (uc *userUsecase) fillProfileDetails(users []models.User) ([]models.User, error) {
	for i, _ := range users{
		profile, err := uc.profileRepo.GetProfile(int64(users[i].ID))

		if err != nil {
			logrus.Println("[usecase][GetAllUser] Error when calling profile repository to get profile")
			return users, err
		}

		users[i].Profile = profile
	}
	return users, nil
}