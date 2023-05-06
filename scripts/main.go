package main

import (
	zrl "github.com/pnkj-kmr/zap-rotate-logger"
)

func main() {
	logger := zrl.New(
		zrl.WithFileName("app"),
		zrl.WithMaxSize(1),
		zrl.WithMaxBackups(30),
		zrl.WithMaxAge(60),
		zrl.WithCompress(true),
		zrl.WithDebug(true),
	)

	for i := 0; i < 100; i++ {
		logger.Info("testing...")
	}
}
