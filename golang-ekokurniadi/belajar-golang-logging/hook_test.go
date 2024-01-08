package belajar_golang_logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	logger.Trace("this is trace")
	logger.Debug("this is debug")
	logger.Info("this is infor")
	logger.Warn("this is warning")
	logger.Error("this is error")
}
