package config

import (
	"github.com/sirupsen/logrus"
	kf "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func (cfg *Config) initKafka() (err error) {

	cfg.KafkaProducer, err = kf.NewProducer(&kf.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "silhouette-registration",
	})

	if err != nil {
		logrus.Println("[config][initKafka] while create producer", err)
	}

	return nil
}
