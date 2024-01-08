package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestFormatterJson(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("this is infor")
	logger.Warn("this is warning")
	logger.Error("this is error")
}
