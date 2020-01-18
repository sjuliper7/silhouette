package config

import (
	"github.com/sirupsen/logrus"
	kf "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func (cfg *Config) initKafka() (err error)  {


	cfg.KafkaConsumer, err = kf.NewConsumer(&kf.ConfigMap{
		"bootstrap.servers": "127.0.0.1",
		"group.id":          "silhouette",
	})

	if err != nil {
		logrus.Errorf("[config][initKafka] Error while create consumer", err)
	}else{
		logrus.Infof("[config][initKafka] Success while create consumer")
	}
	return nil
}
