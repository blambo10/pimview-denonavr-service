package log

import (
	"github.com/sirupsen/logrus"
	"pimview.thelabshack.com/pkg/config"
)

var (
	cfg = config.GetLogger()
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()

	logger.Level = logrus.Level(cfg.Level)

	return logger
}
