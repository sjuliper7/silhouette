package usecase

import (
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"github.com/sjuliper7/silhouette/services/user-service/repositories"
	"log"
)

type userUsecase struct {
	repo repositories.Repository
}

func NewUserUsecase(repo repositories.Repository) UserUsecase {
	return userUsecase{repo}
}

func (uc userUsecase) GetAlluser() (users []models.User, err error) {
	users, err = uc.repo.GetAlluser()

	if err != nil {
		log.Println("Failed when call [repositories][GetAlluser] ", err)
		return nil, err
	}

	return users, err
}
func (uc userUsecase) AddUser(user *models.User) (err error) {

	err = uc.repo.AddUser(user)
	if err != nil {
		log.Println("[usecase][AddUser] Error when calling repository to save")
	}

	return nil
}

func (uc userUsecase) GetUser(userID int64) (user models.User, err error) {
	user, err = uc.repo.GetUser(userID)

	if err != nil {
		log.Println("[usecase][GetUser] Error when calling repository to get user")
		return user, err
	}

	return user, nil
}

func (uc userUsecase) UpdateUser(user *models.User) (err error){

	err = uc.repo.UpdateUser(user)
	if err != nil {
		log.Println("[usecase][UpdateUser] Error when calling repository to update user")
		return err
	}

	return nil
}
