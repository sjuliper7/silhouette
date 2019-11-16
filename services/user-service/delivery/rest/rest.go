package rest

import (
	"github.com/gorilla/mux"
	"github.com/sjuliper7/silhouette/services/user-service/usecase"
	"net/http"
)

// UserServerRest  represent the http handler for users
type UserServerRest struct {
	usecase usecase.UserUsecase
}

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

func NewUserServerRest(uc usecase.UserUsecase) UserServerRest {
	return UserServerRest{usecase: uc}
}

func (usRest UserServerRest) Resource(w http.ResponseWriter, r *http.Request) {
	switch m := r.Method; m {
	case http.MethodGet:
		params := mux.Vars(r)
		if len(params) == 0 {
			usRest.fetchUser(w, r)
		} else {
			usRest.getUser(w, r)
		}
	case http.MethodPost:
		usRest.postUser(w, r)
	case http.MethodPut:
		usRest.updateUser(w, r)
	case http.MethodDelete:
		usRest.deleteUser(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
