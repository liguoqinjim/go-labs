package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path"
	"time"
)

var zaplog *zap.SugaredLogger

func main() {
	InitLogger()
	defer zaplog.Sync()

	zaplog.Infof("这是一条Info日志")
	zaplog.Errorf("这是一条Error日志")
}

func InitLogger() {
	encoder := getEncoder()

	infoWrite := getLogWriter("./", "info", 7)
	errorWrite := getLogWriter("./", "error", 7)

	//过滤写入条件
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWrite), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWrite), warnLevel),
	)

	//添加打印位置，方便调试
	logger := zap.New(core, zap.AddCaller())
	zaplog = logger.Sugar()
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05]"))
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	//自定义时间格式
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logPath, level string, save uint) io.Writer {
	logFullPath := path.Join(logPath, level)

	//日志切割
	hook, err := rotatelogs.New(
		logFullPath+".%Y%m%d%H",
		rotatelogs.WithLinkName(logFullPath),
		rotatelogs.WithRotationCount(save),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}

	return hook
}
