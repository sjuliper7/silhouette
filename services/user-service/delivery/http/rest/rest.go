package rest

import (
	"net/http"
	"strings"

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

//NewUserDelivery ...
func NewUserDelivery(uc usecase.UserUsecase) UserService {
	return UserService{usecase: uc}
}

//Resource ...
func (handler *UserService) Resource(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("req: %s%s\n", r.Host, r.URL.Path)

	switch m := r.Method; m {
	case http.MethodGet:
		params := mux.Vars(r)
		logrus.Infof("params: %v", params)
		var result interface{}
		var err error
		if len(params) == 0 {
			result, err = handler.fetchUser(w, r)
		} else {
			result, err = handler.getUser(w, r)
		}
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			break
		}
		respondWithJSON(w, http.StatusOK, result)

	case http.MethodPost:
		var result interface{}
		var err error

		if strings.Split(r.URL.Path, "/")[3] == "login" {
			result, err = handler.login(w, r)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				break
			}
		} else {
			result, err = handler.postUser(w, r)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				break
			}
		}

		respondWithJSON(w, http.StatusOK, result)
	case http.MethodPut:
		result, err := handler.updateUser(w, r)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			break
		}
		respondWithJSON(w, http.StatusOK, result)

	case http.MethodDelete:
		result, err := handler.deleteUser(w, r)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			break
		}
		respondWithJSON(w, http.StatusOK, result)

	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}

}
