package kafka

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/common/constans"
	"github.com/sjuliper7/silhouette/services/user-service/models"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type kafkaProducerRepository struct {
	kafkaRepo *kafka.Producer
}

func (kRepo kafkaProducerRepository) RegisterDonePublishMessage(profile models.Profile) (err error) {
	defer kRepo.kafkaRepo.Close()

	topic := string(constans.TopicUserRegistration)

	message, err := json.Marshal(profile)
	if err != nil {
		logrus.Println("Failed when parse object to json")
		return err
	}

	//go routine to handle for produce message
	go func() {
		err := kRepo.kafkaRepo.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, nil)

		if err != nil {
			logrus.Println("Error when Publish message")
		}
	}()

	return nil
}
