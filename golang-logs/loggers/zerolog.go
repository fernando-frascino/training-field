package loggers

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupZerolog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Print("hello world")

}

func ZerologError(message string, meta ...interface{}) {
	l := log.Error()
	for key, value := range meta {
		l.Interface(fmt.Sprint(key), value)
	}

	l.Msg(message)
}
