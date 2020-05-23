package main

import (
	"github.com/sirupsen/logrus"
)

func init() {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logrus.SetFormatter(formatter)
	loadConfig()
}

func main() {
	startService()
}
