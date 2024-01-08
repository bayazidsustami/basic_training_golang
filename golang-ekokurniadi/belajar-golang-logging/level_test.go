package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("this is trace")
	logger.Debug("this is debug")
	logger.Info("this is infor")
	logger.Warn("this is warning")
	logger.Error("this is error")
	//logrus.Fatal("this is fatal")
	//logrus.Panic("this is panic")
}

func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel) //default info level above

	logger.Trace("this is trace")
	logger.Debug("this is debug")
	logger.Info("this is infor")
	logger.Warn("this is warning")
	logger.Error("this is error")
	//logrus.Fatal("this is fatal")
	//logrus.Panic("this is panic")
}
