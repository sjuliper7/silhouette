package main

import (
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/config"
	"github.com/sjuliper7/silhouette/commons/models"
	grpc2 "github.com/sjuliper7/silhouette/services/profile-service/delivery/grpc"
	kafkaC "github.com/sjuliper7/silhouette/services/profile-service/delivery/kafka"
	"github.com/sjuliper7/silhouette/services/profile-service/repository/mysql"
	"github.com/sjuliper7/silhouette/services/profile-service/usecase"
	"google.golang.org/grpc"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

//Config is struct to access all config
type Config struct {
	DB            *sqlx.DB
	KafkaConsumer *kafka.Consumer
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("error loading .env file")
	} else {
		logrus.Info("env has been loaded successfully..")
	}

}

func populateStringConnection() string {
	stringConnection := ""

	stringConnection += os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") +
		"@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" +
		os.Getenv("DATABASE_NAME")

	return stringConnection
}

func (cfg *Config) initDatabase() {
	db, err := sqlx.Connect("mysql", populateStringConnection())

	if err != nil {
		logrus.Fatal(err)
	} else {
		logrus.Println("database connected successfully..")
	}

	cfg.DB = db

}

func (cfg *Config) initKafka() (err error) {

	cfg.KafkaConsumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1",
		"group.id":          "silhouette",
	})

	if err != nil {
		logrus.Errorf("[config][initKafka] Error while create consumer", err)
	} else {
		logrus.Infof("[config][initKafka] Success while create consumer")
	}
	return nil
}

func (cfg *Config) initService() {
	repo := mysql.NewMysqlProfileRepository(cfg.DB)
	profileUc := usecase.NewProfileUseCase(repo)

	svr := grpc.NewServer()
	profileServer := grpc2.NewProfileServer(profileUc)
	//
	models.RegisterProfilesServer(svr, profileServer)
	logrus.Infof("starting RPC server at %v", config.SERVICE_PROFILE_PORT)

	err := kafkaC.ConsumeHandler(cfg.KafkaConsumer, profileUc)
	if err != nil {
		logrus.Println("[config][service] failed to start consuming from kafka")
	}

	//next running the to http
	net, err := net.Listen("tcp", config.SERVICE_PROFILE_PORT)
	if err != nil {
		logrus.Fatalln("could not listen to %s: %v", config.SERVICE_PROFILE_PORT, err)
	}

	logrus.Fatalln(svr.Serve(net))
}

func initConfig() {

	var cf Config
	cf.initDatabase()
	cf.initKafka()
	cf.initService()
}
