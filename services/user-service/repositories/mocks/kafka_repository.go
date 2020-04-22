package mocks

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
)

type KafkaRepositoryMock struct {
	mock.Mock
}

func (m *KafkaRepositoryMock) PublishMessage(topic string, message []byte) (err error) {
	logrus.Infof("publish message to topic:  %v message: %v", topic, string(message))
	args := m.Called(topic, message)
	return args.Error(1)
}
