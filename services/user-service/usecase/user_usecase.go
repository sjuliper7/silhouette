package usecase

import (
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/repositories"
	"log"
)

type userUsecase struct {
	userRepo    repositories.UserRepository
	profileRepo repositories.ProfileRepository
}

func NewUserUsecase(userRepo repositories.UserRepository, profileRepo repositories.ProfileRepository) UserUsecase {
	return userUsecase{userRepo, profileRepo}
}

func (uc userUsecase) GetAlluser() (users []models.User, err error) {
	users, err = uc.userRepo.GetAlluser()

	if err != nil {
		log.Println("Failed when call [repositories][GetAlluser] ", err)
		return nil, err
	}

	return users, err
}
func (uc userUsecase) AddUser(user *models.UserTable) (err error) {

	err = uc.userRepo.AddUser(user)
	if err != nil {
		log.Println("[usecase][AddUser] Error when calling repository to save")
	}

	return nil
}

func (uc userUsecase) GetUser(userID int64) (user models.User, err error) {
	ut := models.UserTable{}
	ut, err = uc.userRepo.GetUser(userID)

	if err != nil {
		log.Println("[usecase][GetUser] Error when calling repository to get user")
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
		log.Println("[usecase][GetUser] Error when calling profile repository to get profile")
	}

	user.Profile = profile

	return user, nil
}

func (uc userUsecase) UpdateUser(user *models.UserTable) (err error) {

	err = uc.userRepo.UpdateUser(user)
	if err != nil {
		log.Println("[usecase][UpdateUser] Error when calling repository to update user")
		return err
	}

	return nil
}

func (uc userUsecase) DeleteUser(userID int64) (deleted bool, err error) {
	deleted, err = uc.userRepo.DeleteUser(userID)

	if err != nil {
		log.Println("[usecase][Delete] Error when calling repository to delete user")
		return false, err
	}

	return deleted, nil
}
