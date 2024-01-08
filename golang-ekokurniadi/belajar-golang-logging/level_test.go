package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("this is trace")
	logrus.Debug("this is debug")
	logrus.Info("this is infor")
	logrus.Warn("this is warning")
	logrus.Error("this is error")
	//logrus.Fatal("this is fatal")
	//logrus.Panic("this is panic")
}

func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel) //default info level above

	logger.Trace("this is trace")
	logrus.Debug("this is debug")
	logrus.Info("this is infor")
	logrus.Warn("this is warning")
	logrus.Error("this is error")
	//logrus.Fatal("this is fatal")
	//logrus.Panic("this is panic")
}
