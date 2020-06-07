package main

import (
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/config"
	KafkaDelivery "github.com/sjuliper7/silhouette/services/profile-service/delivery/message_broker/kafka"
	repository "github.com/sjuliper7/silhouette/services/profile-service/repository/mysql"
	"github.com/sjuliper7/silhouette/services/profile-service/usecase"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

//ProfileService is struct to access all config
type ProfileService struct {
	DB            *sqlx.DB
	KafkaConsumer *kafka.Consumer
}

func populateStringConnection() string {
	stringConnection := ""

	stringConnection += os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") +
		"@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" +
		os.Getenv("DATABASE_NAME") + "?parseTime=true"

	return stringConnection
}

func (profileService *ProfileService) startDatabaseDatabase() {
	db, err := sqlx.Connect("mysql", populateStringConnection())

	if err != nil {
		logrus.Fatal(err)
	} else {
		logrus.Println("database connected successfully..")
	}

	profileService.DB = db

}

func (profileService *ProfileService) startKafka() (err error) {

	profileService.KafkaConsumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1",
		"group.id":          "silhouette#1",
	})

	if err != nil {
		logrus.Errorf("[config][initKafka] Error while create consumer: %v", err)
	} else {
		logrus.Infof("[config][initKafka] Success while create consumer")
	}
	return nil
}

func (profileService *ProfileService) startService() {

	profileService.startDatabaseDatabase()
	profileService.startKafka()

	repo := repository.NewMysqlProfileRepository(profileService.DB)
	profileUc := usecase.NewProfileUseCase(repo)

	svr := initRemoteProcedurCall(profileService.DB, profileService.KafkaConsumer, profileUc)

	err := KafkaDelivery.Consume(profileService.KafkaConsumer, profileUc)
	if err != nil {
		logrus.Println("[config][service] failed to start consuming from kafka: %v", err)
	}

	//next running the to http
	net, err := net.Listen("tcp", config.SERVICE_PROFILE_PORT)
	if err != nil {
		logrus.Fatalln("could not listen to %s: %v", config.SERVICE_PROFILE_PORT, err)
	}

	logrus.Fatalln(svr.Serve(net))
}
