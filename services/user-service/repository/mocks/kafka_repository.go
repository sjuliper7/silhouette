package mocks

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
)

//KafkaRepositoryMock ...
type KafkaRepositoryMock struct {
	mock.Mock
}

// PublishMessage ...
func (m *KafkaRepositoryMock) PublishMessage(topic string, message []uint8) (err error) {
	logrus.Infof("publish message...")
	args := m.Called(topic, message)
	return args.Error(0)
}
