package rest

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"net/http"
	"strconv"
	"time"
)

func (userServerRest UserServerRest) fetchUser(w http.ResponseWriter, r *http.Request) {
	users, err := userServerRest.usecase.GetAll()
	if err != nil {
		logrus.Errorf("[delivery][fetchUser] error when call [user-usecase][GetAll], %v", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, users)
}

func (userServerRest UserServerRest) postUser(w http.ResponseWriter, r *http.Request) {
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

	err := userServerRest.usecase.Add(&user)
	if err != nil {
		logrus.Errorf("[delivery][postUser] error when call [user-usecase][Add], %v", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (userServerRest UserServerRest) getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		logrus.Errorf("error when casting params to int, %v", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	user, err := userServerRest.usecase.Get(int64(id))

	if err != nil {
		logrus.Errorf("[delivery][postUser] error when call [user-usecase][Get], %v", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (userServerRest UserServerRest) updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		logrus.Errorf("Error when casting getting user %v", err)
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

	user, err = userServerRest.usecase.Update(user)

	if err != nil {
		logrus.Errorf("[delivery][postUser] error when call [user-usecase][Update], %v", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (userServerRest UserServerRest) deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		logrus.Errorf("Error when casting getting user , %v",err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	deleted, err := userServerRest.usecase.Delete(int64(id))

	if err != nil {
		logrus.Errorf("[delivery][postUser] error when call [user-usecase][Delete], %v", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, deleted)
}
