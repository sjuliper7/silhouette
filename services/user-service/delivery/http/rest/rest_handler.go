package rest

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/models"
)

func (handler UserService) fetchUser(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	users, err := handler.usecase.GetAll()
	if err != nil {
		logrus.Errorf("[delivery][fetchUser] error when call [user-usecase][GetAll], %v", err)
		return nil, err
	}

	return users, nil
}

func (handler UserService) postUser(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	var user = models.User{
		Password: r.FormValue("password"),
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Role:     r.FormValue("role"),
	}

	var profile = models.Profile{
		Address:     r.FormValue("address"),
		WorkAt:      r.FormValue("work_at"),
		PhoneNumber: r.FormValue("phone_number"),
		Gender:      r.FormValue("gender"),
		Name:        r.FormValue("name"),
		DateOfBirth: r.FormValue("date_of_birth"),
	}

	user.Profile = profile

	err := handler.usecase.Add(&user)

	if err != nil {
		logrus.Errorf("[delivery][postUser] error when call [user-usecase][Add], %v", err)
		return nil, err
	}

	user, err = handler.usecase.Get(user.ID)
	if err != nil {
		logrus.Errorf("[delivery][postUser] error when call [user-usecase][Add], %v", err)
		return nil, err
	}

	return user, nil
}

func (handler UserService) getUser(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		logrus.Errorf("error when casting params to int, %v", err)
		return nil, err
	}

	user, err := handler.usecase.Get(int64(id))

	if err != nil {
		logrus.Errorf("[delivery][postUser] error when call [user-usecase][Get], %v", err)
		return nil, err
	}

	return user, nil
}

func (handler UserService) updateUser(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		logrus.Errorf("Error when casting getting user %v", err)
		return nil, err
	}

	user := models.User{}

	user.ID = int64(id)
	user.Username = r.FormValue("username")
	user.Email = r.FormValue("email")
	user.Role = r.FormValue("role")

	profile := models.Profile{
		UserID:      user.ID,
		Address:     r.FormValue("address"),
		WorkAt:      r.FormValue("work_at"),
		PhoneNumber: r.FormValue("phone_number"),
		Gender:      r.FormValue("gender"),
		Name:        r.FormValue("name"),
		DateOfBirth: r.FormValue("date_of_birth"),
	}

	user.Profile = profile

	user, err = handler.usecase.Update(user)

	if err != nil {
		logrus.Errorf("[delivery][postUser] error when call [user-usecase][Update], %v", err)
		return nil, err
	}

	return user, nil
}

func (handler UserService) deleteUser(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		logrus.Errorf("Error when casting getting user , %v", err)
		return nil, err
	}

	deleted, err := handler.usecase.Delete(int64(id))

	if err != nil {
		logrus.Errorf("[delivery][deleteUser] error when call [user-usecase][Delete], %v", err)
		return nil, err
	}

	return deleted, nil
}

func (handler UserService) login(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	user := models.User{}

	user.Password = r.FormValue("password")
	user.Email = r.FormValue("email")

	result, err := handler.usecase.Login(&user)
	if err != nil {
		logrus.Errorf("[delivery][login] error when call [user-usecase][Login], %v", err)
		return nil, err
	}
	return result, nil
}
