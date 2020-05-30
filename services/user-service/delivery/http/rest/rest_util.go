package rest

import (
	"encoding/json"
	"net/http"

	"github.com/sjuliper7/silhouette/services/user-service/models"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := models.Response{}

	response.Code = 200
	response.Status = "error"
	response.Message = message

	res, _ := json.Marshal(response)

	w.Write(res)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := models.Response{}

	response.Code = 200
	response.Status = "success"
	response.Data = payload

	res, _ := json.Marshal(response)

	w.Write(res)
}
