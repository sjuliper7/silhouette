package rest

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"net/http"
	"strconv"
	"time"
)

func (usr UserServerRest) fetchUser(w http.ResponseWriter, r *http.Request) {
	users, err := usr.usecase.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, users)
}

func (usr UserServerRest) postUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Name:     r.FormValue("name"),
		Role:     r.FormValue("role"),
	}

	var profile = models.Profile{
		Address:     r.FormValue("address"),
		WorkAt:      r.FormValue("work_at"),
		PhoneNumber: r.FormValue("phone_number"),
		Gender:      r.FormValue("gender"),
	}

	user.Profile = profile

	err := usr.usecase.Add(&user)
	if err != nil {
		logrus.Errorf("error ,",err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (usr UserServerRest) getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		logrus.Errorf("error when casting params to int", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	user, err := usr.usecase.Get(int64(id))

	if err != nil {
		logrus.Errorf("error, ",err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (usr UserServerRest) updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		logrus.Errorf("Error when casting getting user", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	user := models.User{}

	user.ID = int64(id)
	user.Username = r.FormValue("username")
	user.Email = r.FormValue("email")
	user.Name = r.FormValue("name")
	user.Role = r.FormValue("role")
	user.UpdatedAt = time.Now()

	profile := models.Profile{
		UserID:      user.ID,
		Address:     r.FormValue("address"),
		WorkAt:      r.FormValue("work_at"),
		PhoneNumber: r.FormValue("phone_number"),
		Gender:      r.FormValue("gender"),
	}

	user.Profile = profile

	user, err = usr.usecase.Update(user)

	if err != nil {
		logrus.Errorf("error, ",err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (usr UserServerRest) deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		logrus.Errorf("Error when casting getting user",err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	deleted, err := usr.usecase.Delete(int64(id))

	if err != nil {
		logrus.Errorf("error, ",err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, deleted)
}
