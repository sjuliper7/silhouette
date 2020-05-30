package main

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	commonsConfig "github.com/sjuliper7/silhouette/commons/config"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

//Config is struct to access all config
type UserService struct {
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
		os.Getenv("DATABASE_NAME") + "?parseTime=true"

	logrus.Info(stringConnection)

	return stringConnection
}

func (cfg *UserService) initDatabase() {
	db, err := sqlx.Connect("mysql", populateStringConnection())

	if err != nil {
		logrus.Fatal(err)
	} else {
		logrus.Infof("Database connected successfully..")
	}

	cfg.DB = db
}

func (cfg *UserService) initKafka() (err error) {

	cfg.KafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1",
	})

	if err != nil {
		logrus.Errorf("[config][initKafka] Error while create producer: %v", err)
	} else {
		logrus.Infof("[config][initKafka] Success while create producer")
	}

	return nil
}

func startService() {

	var userService UserService
	userService.initDatabase()
	userService.initKafka()
	logrus.Infof("Starting Rest API at %v", commonsConfig.SERVICE_USER_PORT)

	http.ListenAndServe(commonsConfig.SERVICE_USER_PORT, initRouter(userService.DB, userService.KafkaProducer))
}
