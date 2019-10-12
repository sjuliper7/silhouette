package delivery

import (
	"net/http"
)

func (usr UserServerRest) fetchUser(w http.ResponseWriter, r *http.Request) {
	users, err := usr.usecase.GetAlluser()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, users)
}
