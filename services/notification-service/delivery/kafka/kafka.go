package kafka

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/constans"
	"github.com/sjuliper7/silhouette/services/notification-service/model"
	"github.com/sjuliper7/silhouette/services/notification-service/usecase"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type kafkaService struct {
	kafkaConsumer    *kafka.Consumer
	notificationCase usecase.NotificationUsecase
}

var kafkaTopics = []string{string(constans.TopicUserRegistration),
	string(constans.TopicUserUpdated),
	string(constans.TopicUserDeleted),
}

//Consume ...
func Consume(kafkaConsumer *kafka.Consumer, notificationCase usecase.NotificationUsecase) error {
	kafkaSvc := kafkaService{
		kafkaConsumer:    kafkaConsumer,
		notificationCase: notificationCase,
	}

	kafkaSvc.kafkaConsumer.SubscribeTopics(kafkaTopics, nil)

	go func() {
		for {
			msg, err := kafkaSvc.kafkaConsumer.ReadMessage(-1)
			if err != nil {
				logrus.Errorf("[delivery][kafka][Consume] notifcation-service error occured on consumer %s, detail: %v (%v)", "", err, msg)
				continue
			}
			kafkaSvc.processingMessage(msg)
		}
	}()

	return nil
}

func (kafkaSvc kafkaService) processingMessage(message *kafka.Message) error {

	var topic string
	var err error

	if message.TopicPartition.Topic != nil {
		topic = *message.TopicPartition.Topic
		logrus.Infof("Receive message from kafka topic %v", topic)
	}

	var notification model.Notification

	err = json.Unmarshal(message.Value, &notification)
	if err != nil {
		logrus.Errorf("[delivery][kafka] failed to unmarshall value of message kafka %v", err)
		return err
	}

	if topic == string(constans.TopicUserRegistration) {
		err = kafkaSvc.notificationCase.AccountRegisterNotification(notification)
	} else if topic == string(constans.TopicUserUpdated) {
		err = kafkaSvc.notificationCase.AccountUpdateNotifcation(notification)
	} else if topic == string(constans.TopicUserDeleted) {
		err = kafkaSvc.notificationCase.AccountDeleteNotification(notification)
	}

	return nil
}
