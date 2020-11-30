package log

import (
	"github.com/sirupsen/logrus"
	"strings"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
}

func SetFormatterStr(format string) {
	switch strings.ToLower(format) {
	case "json":
		Logger.SetFormatter(&logrus.JSONFormatter{})
	case "text":
	default:
		Logger.SetFormatter(&logrus.TextFormatter{})
	}
}

func SetLevelStr(level string) {
	switch strings.ToLower(level) {
	case "panic":
		Logger.SetLevel(logrus.PanicLevel)
	case "fatal":
		Logger.SetLevel(logrus.FatalLevel)
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
	case "warning":
		Logger.SetLevel(logrus.WarnLevel)
	case "info":
		Logger.SetLevel(logrus.InfoLevel)
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
	case "trace":
		Logger.SetLevel(logrus.TraceLevel)
	}
}
