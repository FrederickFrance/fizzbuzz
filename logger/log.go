package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is the global logger instance which can be used in all the packages.
var Logger logrus.Logger

func init() {
	level := logrus.InfoLevel

	formatter := new(logrus.JSONFormatter)
	formatter.DisableTimestamp = true

	Logger = logrus.Logger{
		Out:       os.Stdout,
		Formatter: formatter,
		Level:     level,
	}

}
