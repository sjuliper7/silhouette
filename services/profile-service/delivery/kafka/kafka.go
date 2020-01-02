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
	ProfileUC usecase.ProfileUsecase
}

func (kafkaService kafkaSvc) topics() []string {
	return []string{
		string(constans.TopicUserRegistration),
	}
}

func ConsumeHandler(kafkaConsumer *kafka.Consumer, profileUsecase usecase.ProfileUsecase) error {
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
			kafkaService.messageHandling(msg)
		}
	}()

	return nil
}

func (kafkaService kafkaSvc) messageHandling(message *kafka.Message)  {
	topic := ""

	if message.TopicPartition.Topic != nil{
		topic = *message.TopicPartition.Topic
	}

	logrus.Infof("Receive message from kafka topic ", topic)

	var err error

	switch {
	case topic == "kafka.registration.finish":
		err = kafkaService.createPortfolio(message)
	default:

	}


	if err != nil {
		logrus.Print(err)
	}
}
