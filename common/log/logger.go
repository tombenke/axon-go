// Package log provides the global logger for the application
package log

import (
	"github.com/sirupsen/logrus"
	"strings"
)

// Logger is the global logger
var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
}

// SetFormatterStr sets the log format to either `json` or `text`
func SetFormatterStr(format string) {
	switch strings.ToLower(format) {
	case "json":
		Logger.SetFormatter(&logrus.JSONFormatter{})
	case "text":
	default:
		Logger.SetFormatter(&logrus.TextFormatter{})
	}
}

// SetLevelStr sets the log level according to the `level` string parameter
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
