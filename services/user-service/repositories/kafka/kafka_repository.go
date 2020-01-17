package kafka

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/repositories"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type kafkaRepository struct {
	kafkaProducer *kafka.Producer
}

func NewKafkaRepository(kafkaProducer *kafka.Producer) repositories.KafkaRepository {
	return &kafkaRepository{kafkaProducer: kafkaProducer}
}

func (kafkaRepo *kafkaRepository) PublishMessage(topic string, message []byte) (err error) {

	deliverChan := make(chan kafka.Event)

	go func() {
		err := kafkaRepo.kafkaProducer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: message,
		}, deliverChan)

		if err != nil {
			logrus.Errorf("[kafka-repository][PublishMessage] error while producing, %v",err)
			deliverChan <- nil
		}
	}()

	kafkaEvent := <-deliverChan

	if err != nil {
		return nil
	}

	msg := kafkaEvent.(*kafka.Message)

	if msg.TopicPartition.Error != nil {
		err = errors.New("error while publish kafka message")
		logrus.Errorf("[kafka-repository][PublishMessage] %v, ",err)
		return err
	}

	return nil
}
