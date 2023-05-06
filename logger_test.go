package zaprotatelogger_test

import (
	"testing"

	zrl "github.com/pnkj-kmr/zap-rotate-logger"
)

func Test_logging(t *testing.T) {
	logger := zrl.New(
		zrl.WithFileName("test"),
		zrl.WithMaxSize(1),
		zrl.WithMaxBackups(30),
		zrl.WithMaxAge(60),
		zrl.WithCompress(true),
		zrl.WithDebug(true),
	)

	logger.Info("INFO log level message")
	logger.Warn("Warn log level message")
	logger.Error("Error log level message")
}

func Test_logging2(t *testing.T) {
	logger := zrl.New(
		zrl.WithFileName("test"),
		zrl.WithMaxSize(1),
		zrl.WithMaxBackups(30),
		zrl.WithMaxAge(60),
		zrl.WithCompress(true),
		zrl.WithDebug(false),
	)

	logger.Info("INFO log level message")
	logger.Warn("Warn log level message")
	logger.Error("Error log level message")
}
