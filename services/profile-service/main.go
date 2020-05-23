package main

import (
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/config"
	"github.com/sjuliper7/silhouette/commons/models"
	GrpcDelivery "github.com/sjuliper7/silhouette/services/profile-service/delivery/http/grpc"
	KafkaDeivery "github.com/sjuliper7/silhouette/services/profile-service/delivery/http/kafka"
	"github.com/sjuliper7/silhouette/services/profile-service/repository/mysql"
	"github.com/sjuliper7/silhouette/services/profile-service/usecase"
	"google.golang.org/grpc"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

//ProfileService is struct to access all config
type ProfileService struct {
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

func init() {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logrus.SetFormatter(formatter)
	loadConfig()
}

func main() {
	profileService := ProfileService{}
	profileService.startDatabaseDatabase()
	profileService.startKafka()
	profileService.startService()
}

func populateStringConnection() string {
	stringConnection := ""

	stringConnection += os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") +
		"@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" +
		os.Getenv("DATABASE_NAME")

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
		"group.id":          "silhouette",
	})

	if err != nil {
		logrus.Errorf("[config][initKafka] Error while create consumer", err)
	} else {
		logrus.Infof("[config][initKafka] Success while create consumer")
	}
	return nil
}

func (profileService *ProfileService) startService() {
	repo := mysql.NewMysqlProfileRepository(profileService.DB)
	profileUc := usecase.NewProfileUseCase(repo)

	svr := grpc.NewServer()
	profileServer := GrpcDelivery.NewProfileServer(profileUc)

	models.RegisterProfilesServer(svr, profileServer)
	logrus.Infof("starting RPC server at %v", config.SERVICE_PROFILE_PORT)

	err := KafkaDeivery.Consume(profileService.KafkaConsumer, profileUc)
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
