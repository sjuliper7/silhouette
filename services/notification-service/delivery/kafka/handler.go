package kafka

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/constans"
	"github.com/sjuliper7/silhouette/services/notification-service/model"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func (kafkaSvc kafkaDelivery) processingMessage(message *kafka.Message) error {

	var topic string
	var err error

	if message.TopicPartition.Topic != nil {
		topic = *message.TopicPartition.Topic
		logrus.Infof("Receive message from kafka topic: %v", topic)
	}

	var notification model.Notification

	err = json.Unmarshal(message.Value, &notification)
	if err != nil {
		logrus.Errorf("[delivery][kafka] failed to unmarshall value of message kafka: %v", err)
		return err
	}

	logrus.Infof("payload: %v", notification)

	if topic == string(constans.TopicUserRegistration) {
		err = kafkaSvc.notificationCase.AccountRegisterNotification(notification)
	} else if topic == string(constans.TopicUserUpdated) {
		err = kafkaSvc.notificationCase.AccountUpdateNotifcation(notification)
	} else if topic == string(constans.TopicUserDeleted) {
		err = kafkaSvc.notificationCase.AccountDeleteNotification(notification)
	}

	if err != nil {
		logrus.Errorf("[delivery][kafka] failed to process: %v", err)
		return err
	}

	return nil
}
