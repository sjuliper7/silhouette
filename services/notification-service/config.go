package main

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"

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

//Config is struct to access all config
type Config struct {
	Dialler       *gomail.Dialer
	KafkaConsumer *kafka.Consumer
}

func (cfg *Config) startKafka() (err error) {

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

func (cfg Config) startDialler() error {

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

	cfg.Dialler = dialer

	return nil
}

func (cfg *Config) startNotificationService() {
	emailRepository := email.NewEmailRepository(cfg.Dialler)
	notificationUsecase := usecase.NewNotificatonUsecase(emailRepository)

	err := kafkaC.Consume(cfg.KafkaConsumer, notificationUsecase)
	if err != nil {
		logrus.Println("[config][service] failed to start consuming from kafka")
	}

}
