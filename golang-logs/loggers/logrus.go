package loggers

import (
	log "github.com/sirupsen/logrus"
)

func LogrusLogger() {
	log.SetFormatter(&log.JSONFormatter{})

	// log.SetReportCaller(true)

	//simple logging
	// log.WithFields(log.Fields{
	// 	"document": "2398472938423",
	// 	"error":    "Error calling API",
	// }).Error()

}

func LogrusError(message string) {
	if len(message) == 0 {
		message = "Empty message, please log something"
	}
	log.Error(message)
}

// func LogrusErrorWithMetadata(message string, meta map[string]) {

// 	if len(meta) == 0 {
// 		log.Error(message)
// 	} else {meta
// 		log.WithFields(lof.Field(meta)).Error(message)
// 	}
// }
