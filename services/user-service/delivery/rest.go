package delivery

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

func (usrest UserServerRest) Resource(w http.ResponseWriter, r *http.Request) {
	switch m := r.Method; m {
	case http.MethodGet:
		params := mux.Vars(r)
		if len(params) == 0 {
			usrest.fetchUser(w, r)
		} else {
			usrest.getUser(w, r)
		}
	case http.MethodPost:
		usrest.postUser(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
