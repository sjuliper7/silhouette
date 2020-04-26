package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/config"
	"github.com/sjuliper7/silhouette/services/user-service/delivery/http/rest"
	kafkaProducer "github.com/sjuliper7/silhouette/services/user-service/repository/kafka"
	"github.com/sjuliper7/silhouette/services/user-service/repository/mysql"
	"github.com/sjuliper7/silhouette/services/user-service/repository/services"
	"github.com/sjuliper7/silhouette/services/user-service/usecase"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

//Config is struct to access all config
type Config struct {
	DB            *sqlx.DB
	KafkaProducer *kafka.Producer
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("error loading .env file", err)
	} else {
		logrus.Infof("env has been loaded successfully..")
	}
}

func populateStringConnection() string {
	stringConnection := ""

	stringConnection += os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") +
		"@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" +
		os.Getenv("DATABASE_NAME")
	stringConnection += fmt.Sprintf("?%s", "parseTime=true")

	return stringConnection
}

func initRestService(cg *Config) {

	kafkaRepo := kafkaProducer.NewKafkaRepository(cg.KafkaProducer)
	userRepo := mysql.NewUserMysqlRepository(cg.DB)
	profileRepo, err := services.NewProfileRepository()
	if err != nil {
		logrus.Infof("Error when to connect grpc to profile service %v", err)
	}

	profileUsecase := usecase.NewUserUsecase(userRepo, profileRepo, kafkaRepo)

	router := mux.NewRouter()
	userRest := rest.NewUserServerRest(profileUsecase)
	router.HandleFunc("/users", userRest.Resource).Methods("GET", "POST")
	router.HandleFunc("/users/{id}", userRest.Resource).Methods("GET", "PUT", "DELETE")

	logrus.Infof("Starting Rest API at %v", config.REST_USER_PORT)

	http.ListenAndServe(config.REST_USER_PORT, router)
}

func (cfg *Config) initDatabase() {
	db, err := sqlx.Connect("mysql", populateStringConnection())

	if err != nil {
		logrus.Fatal(err)
	} else {
		logrus.Infof("Database connected successfully..")
	}

	cfg.DB = db
}

func (cfg *Config) initKafka() (err error) {

	cfg.KafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1",
	})

	if err != nil {
		logrus.Errorf("[config][initKafka] Error while create producer", err)
	} else {
		logrus.Infof("[config][initKafka] Success while create producer")
	}

	return nil
}

func (cfg *Config) initService() {
	initRestService(cfg)
}

func initConfig() {

	var cf Config
	cf.initDatabase()
	cf.initKafka()
	cf.initService()
}
