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

var (
	zapLog *zap.SugaredLogger
	Log    *Logger //外部调用
)

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

	//配置多种输出
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWrite), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWrite), warnLevel),
	)

	//添加打印位置，方便调试
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	zapLog = logger.Sugar()
	Log = &Logger{zaplog: zapLog}

	//替换global logger，但是不推荐使用global logger的方式
	//调用方式：`zap.S().Infof("this is a global info")`
	zap.ReplaceGlobals(logger)
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

	//没有EncodeTime的时候起作用
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	//caller是全路径还是短路径，FullCallerEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

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

func Debug(args ...interface{}) {
	zapLog.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	zapLog.Debugf(template, args...)
}

func Debugw(template string, keysAndValues ...interface{}) {
	zapLog.Debugw(template, keysAndValues...)
}

func Info(args ...interface{}) {
	zapLog.Info(args...)
}

func Infof(template string, args ...interface{}) {
	zapLog.Infof(template, args...)
}

func Infow(template string, keysAndValues ...interface{}) {
	zapLog.Infow(template, keysAndValues...)
}

func Warn(args ...interface{}) {
	zapLog.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	zapLog.Warnf(template, args...)
}

func WarnW(template string, keysAndValues ...interface{}) {
	zapLog.Warnw(template, keysAndValues...)
}

func Error(args ...interface{}) {
	zapLog.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	zapLog.Errorf(template, args...)
}

func Errorw(template string, keysAndValues ...interface{}) {
	zapLog.Errorw(template, keysAndValues...)
}

func DPanic(args ...interface{}) {
	zapLog.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	zapLog.DPanicf(template, args...)
}

func DPanicw(template string, keysAndValues ...interface{}) {
	zapLog.DPanicw(template, keysAndValues...)
}

func Panic(args ...interface{}) {
	zapLog.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	zapLog.Panicf(template, args...)
}

func Panicw(template string, keysAndValues ...interface{}) {
	zapLog.Panicw(template, keysAndValues...)
}

func Fatal(args ...interface{}) {
	zapLog.Fatal(args...)
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

//输出给外部调用，现在只是给gorm使用
func (l *Logger) Get() *zap.SugaredLogger {
	return l.zaplog
}

func init() {
	initZap("logs/info.log", "logs/error.log")
}
