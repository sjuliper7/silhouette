package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/usecase"
)

// UserService  represent the http handler for users
type UserService struct {
	usecase usecase.UserUsecase
}

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

//NewUserServerRest ...
func NewUserServerRest(uc usecase.UserUsecase) UserService {
	return UserService{usecase: uc}
}

//Resource ...
func (userServerRest UserService) Resource(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("req: %s%s\n", r.Host, r.URL.Path)

	switch m := r.Method; m {
	case http.MethodGet:
		params := mux.Vars(r)
		var result interface{}
		var err error
		if len(params) == 0 {
			result, err = userServerRest.fetchUser(w, r)
		} else {
			result, err = userServerRest.getUser(w, r)
		}
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			break
		}
		respondWithJSON(w, http.StatusOK, result)

	case http.MethodPost:
		result, err := userServerRest.postUser(w, r)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			break
		}
		respondWithJSON(w, http.StatusOK, result)
	case http.MethodPut:
		result, err := userServerRest.updateUser(w, r)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			break
		}
		respondWithJSON(w, http.StatusOK, result)

	case http.MethodDelete:
		result, err := userServerRest.deleteUser(w, r)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			break
		}
		respondWithJSON(w, http.StatusOK, result)

	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}

}
