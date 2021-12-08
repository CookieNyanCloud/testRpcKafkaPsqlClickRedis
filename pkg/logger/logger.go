package logger

import (
	"github.com/cookienyancloud/testrpckafkapsqlclick/pkg/logger/logger"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func Check(msg string, err error) {
	if err != nil {
		logger.Errorf(msg, err)
	}
}
