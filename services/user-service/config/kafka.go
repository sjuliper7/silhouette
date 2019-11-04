package config

import (
	kf "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

func (cfg *Config) initKafka() (err error) {

	cfg.KafkaProducer, err = kf.NewProducer(&kf.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "asgard-koinp2p",
	})

	if err != nil {
		log.Println("[config][initKafka] while create producer", err)
	}

	return nil
}
