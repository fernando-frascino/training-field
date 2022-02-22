package loggers

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupZerolog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func ZerologError(message string, meta map[string]interface{}) {
	zero := log.Error()
	for key, value := range meta {
		zero.Interface(fmt.Sprint(key), value)
	}
	zero.Msg(message)
}
