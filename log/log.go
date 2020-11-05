package log

import (
	"github.com/Guiyunweb/shiki/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var logConf *conf.LogConf

type Level int

var log *zap.SugaredLogger

type SugaredLogger struct {
	base *zap.Logger
}

func Init() {
	logConf = conf.LoadLog()
	fileWriter := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, zapcore.AddSync(fileWriter), logConf.LogLevel)

	logger := zap.New(core)
	log = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() io.Writer {
	file, err := os.OpenFile(logConf.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error(err)
	}
	return file
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05]"))
}

func (s *SugaredLogger) Debug(args ...interface{}) {
	log.Debug(args)
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Warn(args ...interface{}) {
	log.Warn(args)
}

func Error(args ...interface{}) {
	log.Error(args)
}

func DPanic(args ...interface{}) {
	log.DPanic(args)
}
func Panic(args ...interface{}) {
	log.Panic(args)
}
func Fatal(args ...interface{}) {
	log.Fatal(args)
}
