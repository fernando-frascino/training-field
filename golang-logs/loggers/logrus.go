package loggers

import (
	log "github.com/sirupsen/logrus"
)

func LogrusLogger() {

	log.SetFormatter(&log.JSONFormatter{})
	log.Info("My first event from golang to stdout")
}
