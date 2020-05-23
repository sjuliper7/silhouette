package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Response ...
type Response struct {
	Code    int
	Status  string
	Message string
}

//LoadRouter ...
func LoadRouter() (r *mux.Router) {

	r = mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/", hello).Methods("GET")

	return
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{}
	response.Code = 200
	response.Status = "success"
	response.Message = "welcome to notification service"

	json.NewEncoder(w).Encode(response)
}
