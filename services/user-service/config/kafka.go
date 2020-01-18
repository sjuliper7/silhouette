package config

import (
	"github.com/sirupsen/logrus"
	kf "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func (cfg *Config) initKafka() (err error) {

	cfg.KafkaProducer, err = kf.NewProducer(&kf.ConfigMap{
		"bootstrap.servers": "127.0.0.1",
	})

	if err != nil {
		logrus.Errorf("[config][initKafka] Error while create producer", err)
	}else{
		logrus.Infof("[config][initKafka] Success while create producer")
	}

	return nil
}
