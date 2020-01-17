package config

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

//Config is struct to access all config
type Config struct {
	DB            *sqlx.DB
	KafkaProducer *kafka.Producer
}

//LoadConfig is function to load config
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("error loading .env file", err)
	} else {
		logrus.Infof("env has been loaded successfully..")
	}

}

//InitConfig is function to init all config
func InitConfig() {

	var cf Config
	cf.initDatabase()
	cf.initKafka()
	cf.initService()
}
