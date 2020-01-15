package kafka

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/profile-service/helper"
	"github.com/sjuliper7/silhouette/services/profile-service/models"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func (kafkaService kafkaSvc) createProfile(message *kafka.Message) (err error) {
	profile := models.ProfileTable{}
	err = json.Unmarshal(message.Value, &profile)
	helper.CheckError(err)

	err = kafkaService.ProfileUC.AddProfile(profile)
	if err != nil {
		logrus.Errorf("[kafka-handler][createProfile] error when creating profile %v", err)
		return err
	}

	return nil
}

func (kafkaService kafkaSvc) updateProfile(message *kafka.Message) (err error) {
	profile := models.ProfileTable{}
	err = json.Unmarshal(message.Value, &profile)
	helper.CheckError(err)

	err = kafkaService.ProfileUC.UpdateProfile(profile)

	if err != nil {
		logrus.Println("[kafka-handler][updateProfile] error when updating profile", err)
		return err
	}

	return nil
}

func (kafkaService kafkaSvc) deleteProfile(message *kafka.Message) ( err error) {

	profile := models.ProfileTable{}
	err = json.Unmarshal(message.Value, &profile)
	helper.CheckError(err)

	err = kafkaService.ProfileUC.DeleteProfile(int64(profile.ID))

	if err != nil {
		logrus.Errorf("[kafka-handler][updateProfile] error when deleting profile %v", err)
		return  err
	}

	return  nil
}
