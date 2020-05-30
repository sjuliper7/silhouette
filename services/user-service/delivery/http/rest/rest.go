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
		if len(params) == 0 {
			userServerRest.fetchUser(w, r)
		} else {
			userServerRest.getUser(w, r)
		}
	case http.MethodPost:
		userServerRest.postUser(w, r)
	case http.MethodPut:
		userServerRest.updateUser(w, r)
	case http.MethodDelete:
		userServerRest.deleteUser(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
