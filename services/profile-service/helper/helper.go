package helper

import (
	"github.com/sirupsen/logrus"
)

func CheckError(err error) {
	if err != nil {
		logrus.Errorf("[helper] error, %v", err)
	}
}
