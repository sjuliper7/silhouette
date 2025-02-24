package helper

import (
	"time"

	"github.com/sirupsen/logrus"
)

const (
	RFC3339 = "2006-01-02 15:04:05"
)

func CheckError(err error) {
	if err != nil {
		logrus.Errorf("[helper] error, %v", err)
	}
}

func ParseStringToTime(date string) time.Time {

	dateTime, err := time.Parse(time.RFC3339, date)

	if err != nil {
		logrus.Infof("[helper] error when parse string to time , %v", err)
		dateTime = time.Now()
	}

	return dateTime
}
