package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sjuliper7/silhouette/services/profile-service/models"
)

//LoadRouter ...
func LoadRouter() (r *mux.Router) {

	r = mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/", hello).Methods("GET")

	return
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := models.Response{}
	response.Code = 200
	response.Status = "success"
	response.Message = "welcome to profile service"

	json.NewEncoder(w).Encode(response)
}
