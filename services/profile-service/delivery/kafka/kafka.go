package kafka

import (
	"github.com/koinworks/asgard-heimdal/libs/logger"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/constans"
	"github.com/sjuliper7/silhouette/services/profile-service/usecase"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type kafkaSvc struct {
	kafkaConsumer *kafka.Consumer
	ProfileUC usecase.ProfileUseCase
}

func (kafkaService kafkaSvc) topics() []string {
	return []string{
		string(constans.TopicUserRegistration),
	}
}

func ConsumeHandler(kafkaConsumer *kafka.Consumer, profileUsecase usecase.ProfileUseCase) error {
	kafkaService := kafkaSvc{
		kafkaConsumer: kafkaConsumer,
		ProfileUC:     profileUsecase,
	}

	kafkaService.kafkaConsumer.SubscribeTopics(kafkaService.topics(), nil)

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

func (kafkaService kafkaSvc) messageHandler(message *kafka.Message)  {
	topic := ""

	if message.TopicPartition.Topic != nil{
		topic = *message.TopicPartition.Topic
	}

	logrus.Infof("Receive message from kafka topic %v", topic)

	var err error

	switch {
	case topic == "user.registration.finish":
		err = kafkaService.createProfile(message)
	case topic == "user.updated.finish":
		err = kafkaService.updateProfile(message)
	case topic == "user.deleted.finish":
		err = kafkaService.deleteProfile(message)
	default:

	}


	if err != nil {
		logrus.Print(err)
	}
}
