package kafka

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/constans"
	"github.com/sjuliper7/silhouette/services/notification-service/usecase"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type kafkaDelivery struct {
	kafkaConsumer    *kafka.Consumer
	notificationCase usecase.NotificationUsecase
}

var kafkaTopics = []string{string(constans.TopicUserRegistration),
	string(constans.TopicUserUpdated),
	string(constans.TopicUserDeleted),
}

//Consume ...
func Consume(kafkaConsumer *kafka.Consumer, notificationCase usecase.NotificationUsecase) error {
	kafkaSvc := kafkaDelivery{
		kafkaConsumer:    kafkaConsumer,
		notificationCase: notificationCase,
	}

	kafkaSvc.kafkaConsumer.SubscribeTopics(kafkaTopics, nil)

	go func() {
		for {
			msg, err := kafkaSvc.kafkaConsumer.ReadMessage(-1)
			if err != nil {
				logrus.Errorf("[delivery][kafka][Consume] notifcation-service error consumer %s, detail: %v (%v)", "", err, msg)
				continue
			}
			kafkaSvc.processingMessage(msg)
		}
	}()

	return nil
}
