package kafka

import (
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/profile-service/helper"
	"github.com/sjuliper7/silhouette/services/profile-service/models"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func (kafkaService kafkaDelivery) createProfile(message *kafka.Message) (err error) {
	logrus.Infof("request : %v", string(message.Value))
	outputProfile := models.OutputKafkaProfile{}
	err = json.Unmarshal(message.Value, &outputProfile)
	helper.CheckError(err)

	profile := models.ProfileTable{}
	profile.IsActive = true
	profile.Gender = outputProfile.Gender
	profile.PhoneNumber = outputProfile.PhoneNumber
	profile.WorkAt = outputProfile.WorkAt
	profile.Address = outputProfile.Address
	profile.UserId = outputProfile.UserId

	err = kafkaService.profileUsecase.Add(profile)
	if err != nil {
		logrus.Errorf("[kafka-handler][createProfile] error when creating profile: %v", err)
		return err
	}

	return nil
}

func (kafkaService kafkaDelivery) updateProfile(message *kafka.Message) (err error) {
	logrus.Infof("request : %v", string(message.Value))

	temp := models.OutputKafkaProfile{}

	err = json.Unmarshal(message.Value, &temp)
	helper.CheckError(err)

	profile := models.ProfileTable{
		UserId:      temp.UserId,
		Address:     temp.Address,
		WorkAt:      temp.WorkAt,
		PhoneNumber: temp.PhoneNumber,
		Gender:      temp.Gender,
		IsActive:    true,
		UpdatedAt:   time.Now(),
	}

	err = kafkaService.profileUsecase.Update(profile)

	if err != nil {
		logrus.Println("[kafka-handler][updateProfile] error when updating profile: %v", err)
		return err
	}

	return nil
}

func (kafkaService kafkaDelivery) deleteProfile(message *kafka.Message) (err error) {
	logrus.Infof("request : %v", string(message.Value))
	profile := models.ProfileTable{}
	err = json.Unmarshal(message.Value, &profile)
	helper.CheckError(err)

	err = kafkaService.profileUsecase.Delete(int64(profile.UserId))

	if err != nil {
		logrus.Errorf("[kafka-handler][updateProfile] error when deleting profile: %v", err)
		return err
	}

	return nil
}
