package main

import "go.uber.org/zap"

var zaplog *zap.SugaredLogger

func InitLogger() {
	logger, _ := zap.NewProduction()
	zaplog = logger.Sugar()
}

func main() {
	InitLogger()
	defer zaplog.Sync()

	zaplog.Infof("这是一条Info日志")
	zaplog.Errorf("这是一条Error日志")
}
