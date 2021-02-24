package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path"
	"time"
)

var zapLog *zap.SugaredLogger
var Log *Logger

func initZap(infoFilePath, errorFilePath string) {
	encoder := getEncoder()

	infoWrite := getLogWriter(infoFilePath, "info", 7)
	errorWrite := getLogWriter(errorFilePath, "error", 7)

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
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zapLog = logger.Sugar()

	Log = &Logger{zaplog: zapLog}
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05]"))
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	//自定义时间格式
	encoderConfig.EncodeTime = customTimeEncoder
	//定义日志中的 日志级别显示和颜色
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

func Infof(template string, args ...interface{}) {
	zapLog.Infof(template, args...)
}

func Infow(template string, args ...interface{}) {
	zapLog.Infow(template, args...)
}

func Errorf(template string, args ...interface{}) {
	zapLog.Errorf(template, args...)
}

func Errorw(template string, args ...interface{}) {
	zapLog.Errorw(template, args...)
}

func Debugf(template string, args ...interface{}) {
	zapLog.Debugf(template, args...)
}

func Debugw(template string, args ...interface{}) {
	zapLog.Debugw(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	zapLog.Fatalf(template, args...)
}

func Fatalw(template string, args ...interface{}) {
	zapLog.Fatalw(template, args...)
}

type Logger struct {
	zaplog *zap.SugaredLogger
}

func (l *Logger) Print(msg ...interface{}) {
	zapLog.Info(msg...)
}

func (l *Logger) Println(msg ...interface{}) {
	zapLog.Info(msg...)
}
func (l *Logger) Error(msg ...interface{}) {
	zapLog.Error(msg...)
}
func (l *Logger) Warn(msg ...interface{}) {
	zapLog.Warn(msg...)
}
func (l *Logger) Info(msg ...interface{}) {
	zapLog.Info(msg...)
}
func (l *Logger) Debug(msg ...interface{}) {
	zapLog.Debug(msg...)
}

func (l *Logger) Get() *zap.SugaredLogger {
	return l.zaplog
}

func init() {
	initZap("logs/info.log", "logs/error.log")
}
