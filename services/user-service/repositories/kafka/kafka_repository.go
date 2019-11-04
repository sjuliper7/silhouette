package kafka

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type kafkaProducer struct {
	kafkaRepo *kafka.Producer
}
