package main

import (
	"go.uber.org/zap"
	"lab164/lab005/logger"
)

func main() {
	{
		logger.Debug("this is a debug", "debug message")
		logger.Infof("this is a info")
		logger.Warn("this is a warning")
		logger.Errorf("this is a error")
	}

	//global logger
	{
		zap.S().Infof("this is a global info")
		zap.S().Errorf("this is a error info")
	}
}
