package loggers

import (
	"github.com/sirupsen/logrus"
)

func SetupLogrus() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
}

func LogrusError(meta interface{}) {
	// if meta == "" {
	// 	logrus.Error(message)
	// }
	// logrus.Error(message, meta)
	logrus.Error(meta)
}
