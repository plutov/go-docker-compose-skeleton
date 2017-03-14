package shared

import (
	"github.com/sirupsen/logrus"
)

// LogErr func
func LogErr(err error) {
	if err != nil {
		logrus.Errorln(err.Error())
	}
}

// Infof func
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// Errorf func
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}
