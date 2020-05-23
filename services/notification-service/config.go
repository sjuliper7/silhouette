package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/config"
	kafkaC "github.com/sjuliper7/silhouette/services/notification-service/delivery/kafka"
	"github.com/sjuliper7/silhouette/services/notification-service/repository/email"
	"github.com/sjuliper7/silhouette/services/notification-service/usecase"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"gopkg.in/gomail.v2"
)

var (
	//ConfigEmail ...
	ConfigEmail string

	//ConfigPassword ...
	ConfigPassword string

	//ConfigSMTPHost ...
	ConfigSMTPHost string

	//ConfigSMTPPort ...
	ConfigSMTPPort int
)

//NotificationService is struct to access all config
type NotificationService struct {
	Dialler       *gomail.Dialer
	KafkaConsumer *kafka.Consumer
}

func (nfs *NotificationService) startKafka() (err error) {

	nfs.KafkaConsumer, err = kafka.NewConsumer(&kafka.ConfigMap{
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

func (nfs *NotificationService) startEmailDialler() error {

	ConfigEmail = os.Getenv("CONFIG_EMAI")
	ConfigPassword = os.Getenv("CONFIG_PASSWORD")
	ConfigSMTPHost = os.Getenv("CONFIG_SMTP_HOST")
	port := os.Getenv("CONFIG_SMTP_PORT")

	ConfigSMTPPort, _ = strconv.Atoi(port)

	dialer := gomail.NewDialer(
		ConfigSMTPHost,
		ConfigSMTPPort,
		ConfigEmail,
		ConfigPassword,
	)

	nfs.Dialler = dialer

	return nil
}

func (nfs *NotificationService) startService() {

	nfs.startKafka()
	nfs.startEmailDialler()

	emailRepository := email.NewEmailRepository(nfs.Dialler)
	notificationUsecase := usecase.NewNotificatonUsecase(emailRepository)

	err := kafkaC.Consume(nfs.KafkaConsumer, notificationUsecase)
	if err != nil {
		logrus.Println("[config][service] failed to start consuming from kafka")
	}

	http.ListenAndServe(config.SERVICE_NOTIFICATION_PORT, LoadRouter())

}
