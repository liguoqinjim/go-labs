package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	//SugarLogger
	{
		logger, _ := zap.NewProduction()
		defer logger.Sync() // flushes buffer, if any

		sugar := logger.Sugar()
		sugar.Infow("failed to fetch URL",
			// Structured context as loosely typed key-value pairs.
			"url", "example.com",
			"attempt", 3,
			"backoff", time.Second,
		)

		sugar.Infof("Failed to fetch URL: %s", "example")
	}

	//Logger
	{
		logger2, _ := zap.NewProduction()
		defer logger2.Sync()
		logger2.Info("failed to fetch URL",
			// Structured context as strongly typed Field values.
			zap.String("url", "example.com"),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
	}
}
