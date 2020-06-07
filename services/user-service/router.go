package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/delivery/http/rest"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	kafkaProducer "github.com/sjuliper7/silhouette/services/user-service/repository/kafka"
	repository "github.com/sjuliper7/silhouette/services/user-service/repository/mysql"
	"github.com/sjuliper7/silhouette/services/user-service/repository/services"
	"github.com/sjuliper7/silhouette/services/user-service/usecase"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func initRouter(db *sqlx.DB, kafka *kafka.Producer) *mux.Router {
	kafkaRepo := kafkaProducer.NewKafkaRepository(kafka)
	userRepo := repository.NewUserMysqlRepository(db)
	profileRepo, err := services.NewProfileRepository()
	if err != nil {
		logrus.Infof("Error when to connect grpc to profile service %v", err)
	}

	profileUsecase := usecase.NewUserUsecase(userRepo, profileRepo, kafkaRepo)
	userRest := rest.NewUserDelivery(profileUsecase)

	router := mux.NewRouter()
	v1 := router.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/", hello).Methods("GET")

	v1.HandleFunc("/users", userRest.Resource).Methods("GET", "POST")
	v1.HandleFunc("/users/{id}", userRest.Resource).Methods("GET", "PUT", "DELETE")

	return router
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := models.Response{}
	response.Code = 200
	response.Status = "success"
	response.Message = "welcome to user service"

	json.NewEncoder(w).Encode(response)
}
