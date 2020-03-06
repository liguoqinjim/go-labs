package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var zaplog *zap.SugaredLogger

func main() {
	InitLogger()
	defer zaplog.Sync()

	zaplog.Infof("这是一条Info日志")
	zaplog.Errorf("这是一条Error日志")
}

func InitLogger() {
	fileWriter := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, zapcore.AddSync(fileWriter), zapcore.DebugLevel)

	logger := zap.New(core)
	zaplog = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() io.Writer {
	file, _ := os.Create("./test.log")
	return file
}
