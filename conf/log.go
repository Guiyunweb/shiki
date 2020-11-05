package conf

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

// 日志配置
type LogConf struct {
	LogLevel zapcore.Level
	LogPath  string
}

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel zapcore.Level = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
	_minLevel = DebugLevel
	_maxLevel = FatalLevel
)

func LoadLog() *LogConf {
	LogLevel := Conf.GetString("LOG_LEVEL")

	// 查询判断LOG_LEVEL等级,默认debug
	var LogLevelNum = DebugLevel
	switch strings.ToLower(LogLevel) {
	case "debug":
		LogLevelNum = DebugLevel
		break
	case "info":
		LogLevelNum = InfoLevel
		break
	case "warning":
		LogLevelNum = WarnLevel
		break
	case "error":
		LogLevelNum = ErrorLevel
		break
	case "fatal":
		LogLevelNum = FatalLevel
		break
	}

	LogPath := Conf.GetString("LOG_PATH")
	if LogLevel == "" {
		LogPath = "shike.log"
	}
	return &LogConf{
		LogLevel: LogLevelNum,
		LogPath:  LogPath,
	}
}
