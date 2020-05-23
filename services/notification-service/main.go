package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("error loading .env file")
	} else {
		logrus.Info("env has been loaded successfully..")
	}

}

//
func init() {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logrus.SetFormatter(formatter)
	loadConfig()
}

func main() {
	//start service
	notificationService := NotificationService{}
	notificationService.startService()

}
