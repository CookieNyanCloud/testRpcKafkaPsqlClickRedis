package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init()  {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{})
}