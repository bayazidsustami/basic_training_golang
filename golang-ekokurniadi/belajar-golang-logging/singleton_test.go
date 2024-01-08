package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSingleton(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Trace("this is trace")
	logrus.Debug("this is debug")
	logrus.Info("this is infor")
	logrus.Warn("this is warning")
	logrus.Error("this is error")
}
