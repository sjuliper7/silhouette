package main

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/user-service/config"
)

func init() {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logrus.SetFormatter(formatter)
	config.LoadConfig()
}

func main() {
	config.InitConfig()
}
