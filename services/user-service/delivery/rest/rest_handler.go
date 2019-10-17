package rest

import (
	"github.com/gorilla/mux"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"log"
	"net/http"
	"strconv"
)

func (usr UserServerRest) fetchUser(w http.ResponseWriter, r *http.Request) {
	users, err := usr.usecase.GetAlluser()
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

	err := usr.usecase.AddUser(&user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (usr UserServerRest) getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Println("Error when casting params to int")
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	user, err := usr.usecase.GetUser(int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (usr UserServerRest) updateUser(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Println("Error when casting getting user")
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	user := models.User{}

	user.ID = uint64(id)
	user.Username = r.FormValue("username")
	user.Email = r.FormValue("email")
	user.Name = r.FormValue("name")
	user.Role = r.FormValue("role")

	err = usr.usecase.UpdateUser(&user)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (usr UserServerRest) deleteUser(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Println("Error when casting getting user")
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	deleted, err := usr.usecase.DeleteUser(int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, deleted)
}