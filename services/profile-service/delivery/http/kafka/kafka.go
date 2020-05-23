package kafka

import (
	"github.com/koinworks/asgard-heimdal/libs/logger"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/constans"
	"github.com/sjuliper7/silhouette/services/profile-service/usecase"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type kafkaDelivery struct {
	kafkaConsumer  *kafka.Consumer
	profileUsecase usecase.ProfileUseCase
}

var kafkaTopics = []string{string(constans.TopicUserRegistration),
	string(constans.TopicUserUpdated),
	string(constans.TopicUserDeleted),
}

//Consume ...
func Consume(kafkaConsumer *kafka.Consumer, profileUsecase usecase.ProfileUseCase) error {
	kafkaService := kafkaDelivery{
		kafkaConsumer:  kafkaConsumer,
		profileUsecase: profileUsecase,
	}

	kafkaService.kafkaConsumer.SubscribeTopics(kafkaTopics, nil)

	go func() {
		for {
			msg, err := kafkaService.kafkaConsumer.ReadMessage(-1)
			if err != nil {
				logger.Errf("[delivery][kafka][NewKafkaHandler] portfolio error occured on consumer %s, detail: %v (%v)", "", err, msg)
				continue
			}
			kafkaService.messageHandler(msg)
		}
	}()

	return nil
}

func (kafkaService kafkaDelivery) messageHandler(message *kafka.Message) {
	topic := ""
	var err error

	if message.TopicPartition.Topic != nil {
		topic = *message.TopicPartition.Topic
	}

	logrus.Infof("Receive message from kafka topic %v", topic)

	switch {
	case topic == string(constans.TopicUserRegistration):
		err = kafkaService.createProfile(message)
	case topic == string(constans.TopicUserUpdated):
		err = kafkaService.updateProfile(message)
	case topic == string(constans.TopicUserDeleted):
		err = kafkaService.deleteProfile(message)
	default:

	}

	if err != nil {
		logrus.Error(err)
	}
}
